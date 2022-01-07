package postgres

import (
	"database/sql"
	"time"

	"github.com/huandu/go-sqlbuilder"
	pb "github.com/muhriddinsalohiddin/online_store_catalog/genproto/catalog_service"
	"github.com/muhriddinsalohiddin/online_store_catalog/pkg/utils"
)

// CRUD for Books

func (c *catalogRepo) CreateBook(in pb.Book) (pb.Book, error) {
	err := c.db.QueryRow(`
		INSERT INTO books (id, name, author_id, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5) returning id`,
		in.Id,
		in.Name,
		in.AuthorId,
		time.Now().UTC(),
		time.Now().UTC(),
	).Scan(&in.Id)
	if err != nil {
		return pb.Book{}, err
	}
	for _, categoryId := range in.CategoryId {
		err = c.db.QueryRow(`
			INSERT INTO books_categories (book_id, category_id) 
			VALUES ($1, $2)`,
			in.Id,
			categoryId,
		).Err()
		if err != nil {
			return pb.Book{}, err
		}
	}

	book, err := c.GetBookById(pb.GetBookByIdReq{Id: in.Id})
	if err != nil {
		return pb.Book{}, err
	}
	return book, nil
}

func (c *catalogRepo) UpdateBook(in pb.Book) (pb.Book, error) {
	result, err := c.db.Exec(`
		UPDATE books
		SET name=$1,
			author_id = $2
			updated_at=$3
		WHERE id = $4`,
		in.Name,
		in.AuthorId,
		time.Now().UTC(),
		in.Id,
	)
	if err != nil {
		return pb.Book{}, err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Book{}, sql.ErrNoRows
	}
	book, err := c.GetBookById(pb.GetBookByIdReq{Id: in.Id})
	if err != nil {
		return pb.Book{}, err
	}
	return book, nil
}

func (c *catalogRepo) GetBookById(in pb.GetBookByIdReq) (pb.Book, error) {
	var book pb.Book
	err := c.db.QueryRow(`
		SELECT
			name,
			author_id,
			created_at,
			updated_at
		FROM books
		WHERE id = $1
		AND deleated_at IS NULL`,
		in.Id).Scan(
		&book.Name,
		&book.AuthorId,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if err != nil {
		return pb.Book{}, err
	}

	return book, nil
}

func (c *catalogRepo) DeletedBookById(in pb.GetBookByIdReq) (pb.EmptyResp, error) {
	result, err := c.db.Exec(`
		UPDATE books
		SET deleated_at=$1
		WHERE id = $2`,
		time.Now().UTC(),
		in.Id,
	)
	if err != nil {
		return pb.EmptyResp{}, err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return pb.EmptyResp{}, sql.ErrNoRows
	}
	return pb.EmptyResp{}, nil
}

func (c *catalogRepo) ListBooks(in pb.ListBookReq) (pb.ListBookResp, error) {
	offset := (in.Page - 1) * in.Limit

	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "name", "author_id", "created_at", "updated_at")
	sb.From("books b")
	if value, ok := in.Filters["category"]; ok && value != "" {
		args := utils.StringSliceToInterfaceSlice(utils.ParseFilter(value))
		sb.JoinWithOption("LEFT", "books_categories bc", "b.id=bc.book_id")
		sb.Where(sb.In("bc.category_id", args...))
	}
	if value, ok := in.Filters["author"]; ok && value != "" {
		sb.Where(sb.Equal("author", value))
	}
	sb.Limit(int(in.Limit))
	sb.Offset(int(offset))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := c.db.Queryx(query, args...)
	if err != nil {
		return pb.ListBookResp{}, err
	}
	if err = rows.Err(); err != nil {
		return pb.ListBookResp{}, err
	}
	defer rows.Close() // nolint:errcheck

	var (
		books []*pb.Book
		count int64
	)

	for rows.Next() {

		var book pb.Book
		err = rows.Scan(
			&book.Name,
			&book.AuthorId,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return pb.ListBookResp{}, err
		}
		books = append(books, &book)
	}

	sbc := sqlbuilder.NewSelectBuilder()
	sbc.Select("count(*)")
	sbc.From("books b")

	if value, ok := in.Filters["category"]; ok {
		args = utils.StringSliceToInterfaceSlice(utils.ParseFilter(value))
		sbc.JoinWithOption("LEFT", "books_categories bc", "b.id=bc.book_id")
		sbc.Where(sbc.In("bc.book_id", args...))
	}

	if value, ok := in.Filters["author"]; ok {
		sbc.Where(sbc.Equal("author", value))
	}
	query, args = sbc.BuildWithFlavor(sqlbuilder.PostgreSQL)

	err = c.db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return pb.ListBookResp{}, err
	}

	return pb.ListBookResp{Books: books, Count: count}, nil
}
