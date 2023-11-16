package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Category struct {
	ID   int
	Name string
}

type CategoryRepository struct {
	dbPool *pgxpool.Pool
}

func NewCategoryRepository(dbPool *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{
		dbPool: dbPool,
	}
}

func (r *CategoryRepository) GetAllCategories(ctx context.Context) ([]Category, error) {
	rows, err := r.dbPool.Query(ctx, "SELECT * FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}
