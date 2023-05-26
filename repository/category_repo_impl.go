package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
)

type CategoryRepoImpl struct {
}

func NewCategoryRepo() CategoryRepository {
	return &CategoryRepoImpl{}
}

func (repository *CategoryRepoImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	var lastid int
	name := category.Name
	fmt.Print(name)
	SQL := `insert into mst_category.public.category(name) values ($1) returning id`
	err := tx.QueryRowContext(ctx, SQL, category.Name).Scan(&lastid)
	helper.PanicIfError(err)
	category.Id = lastid
	return category
}

func (repository *CategoryRepoImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepoImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepoImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
