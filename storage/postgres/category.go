package postgres

import (
	"database/sql"
	"time"

	pb "github.com/muhriddinsalohiddin/online_store_catalog/genproto/catalog_service"
)

func (c *catalogRepo) CreateCategory(in pb.Category) (pb.Category, error) {
	err := c.db.QueryRow(`
	INSERT INTO categories (id, name, parent_id, created_at, updated_at)
	VALUES ($1,$2,$3,$4,$5)`,
		in.Id,
		in.Name,
		in.ParentId,
		time.Now().UTC(),
		time.Now().UTC()).Err()
	if err != nil {
		return pb.Category{}, err
	}
	category, err := c.GetCategoryById(pb.GetCategoryByIdReq{Id: in.Id})
	if err != nil {
		return pb.Category{}, err
	}

	return category, nil
}

func (c *catalogRepo) UpdateCategory(in pb.Category) (pb.Category, error) {
	result, err := c.db.Exec(`
		UPDATE categories
		SET name=$1,
			parent_id=$2,
			updated_at=$3
		WHERE id = $4`,
		in.Name,
		in.ParentId,
		time.Now().UTC(),
		in.Id,
	)
	if err != nil {
		return pb.Category{}, err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Category{}, sql.ErrNoRows
	}
	category, err := c.GetCategoryById(pb.GetCategoryByIdReq{Id: in.Id})
	if err != nil {
		return pb.Category{}, err
	}
	return category, nil
}

func (c *catalogRepo) GetCategoryById(in pb.GetCategoryByIdReq) (pb.Category, error) {
	var category pb.Category
	err := c.db.QueryRow(`
		SELECT
			name,
			parent_id,
			created_at,
			updated_at
		FROM categories
		WHERE id = $1
		AND deleated_at IS NULL`,
		in.Id).Scan(
		&category.Name,
		&category.ParentId,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		return pb.Category{}, err
	}
	return category, nil
}

func (c *catalogRepo) DeleteCategoryById(in pb.GetCategoryByIdReq) error {
	result, err := c.db.Exec(`
		UPDATE categories
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

func (c *catalogRepo) ListCategories(in pb.ListCategoryReq) (pb.ListCategoryResp, error) {
	offset := (in.Page - 1) * in.Limit
	rows, err := c.db.Query(`
		SELECT
			name,
			parent_id,
			created_at,
			updated_at
		FROM categories
		WHERE deleated_at IS NULL
		LIMIT $1
		OFFSET $2`,
		in.Limit,
		offset,
	)
	if err != nil {
		return pb.ListCategoryResp{}, err
	}
	var categories pb.ListCategoryResp

	for rows.Next() {

		var category pb.Category
		err := rows.Scan(
			&category.Name,
			&category.ParentId,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return pb.ListCategoryResp{}, err
		}
		categories.Categories = append(categories.Categories, &category)
	}
	err = c.db.QueryRow(`
		SELECT COUNT(*)
		FROM categories
		WHERE deleated_at IS NULL`,
	).Err()
	if err != nil {
		return pb.ListCategoryResp{}, err
	}
	return categories, nil
}
