package postgres

import (
	"database/sql"
	"time"

	pb "github.com/muhriddinsalohiddin/online_store_catalog/genproto/catalog_service"
)

func (c *catalogRepo) CreateAuthor(in pb.Author) (pb.Author, error) {
	err := c.db.QueryRow(`
		INSERT INTO authors (id, name, created_at, updated_at)
		VALUES ($1,$2,$3,$4) RETURNING id`,
		in.Id,
		in.Name,
		time.Now().UTC(),
		time.Now().UTC()).Scan(&in.Id)
	if err != nil {
		return pb.Author{}, err
	}

	author, err := c.GetAuthorById(pb.GetAuthorByIdReq{Id: in.Id})
	if err != nil {
		return pb.Author{}, err
	}
	return author, nil
}

func (c *catalogRepo) UpdateAuthor(in pb.Author) (pb.Author, error) {
	result, err := c.db.Exec(`
		UPDATE authors
		SET name=$1,
			updated_at=$2
		WHERE id = $3`,
		in.Name,
		time.Now().UTC(),
		in.Id,
	)
	if err != nil {
		return pb.Author{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Author{}, sql.ErrNoRows
	}

	author, err := c.GetAuthorById(pb.GetAuthorByIdReq{Id: in.Id})
	if err != nil {
		return pb.Author{}, err
	}

	return author, nil
}

func (c *catalogRepo) GetAuthorById(in pb.GetAuthorByIdReq) (pb.Author, error) {
	var author pb.Author

	err := c.db.QueryRow(`
		SELECT 
			name,
			created_at,
			updated_at
		FROM authors
		WHERE id = $1
		AND deleated_at IS NULL`,
		in.Id).Scan(
		&author.Name,
		&author.CreatedAt,
		&author.UpdatedAt,
	)
	if err != nil {
		return pb.Author{}, err
	}

	return author, nil
}

func (c *catalogRepo) DeleteAuthorById(in pb.GetAuthorByIdReq) error {
	result, err := c.db.Exec(`
		UPDATE authors
		SET deleated_at=$1
		WHERE id = $2`,
		time.Now().UTC(),
		in.Id,
	)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (c *catalogRepo) ListAuthors(in pb.ListAuthorReq) (pb.ListAuthorResp, error) {
	offset := (in.Page - 1) * in.Limit
	rows, err := c.db.Query(`
		SELECT
			name,
			created_at,
			updated_at
		FROM authors
		WHERE deleated_at IS NULL
		LIMIT $1
		OFFSET $2`,
		in.Limit,
		offset,
	)
	if err != nil {
		return pb.ListAuthorResp{}, err
	}
	defer rows.Close()
	var authors pb.ListAuthorResp

	for rows.Next() {

		var author pb.Author
		err := rows.Scan(
			&author.Name,
			&author.CreatedAt,
			&author.UpdatedAt,
		)
		if err != nil {
			return pb.ListAuthorResp{}, err
		}
		authors.Authors = append(authors.Authors, &author)
	}
	err = c.db.QueryRow(`
		SELECT COUNT(*)
		FROM authors
		WHERE deleated_at IS NULL`,
	).Scan(&authors.Count)
	if err != nil {
		return pb.ListAuthorResp{},err
	}
	return authors, nil
}
