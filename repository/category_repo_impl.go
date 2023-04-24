package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
)

type CategoryRepoImpl struct {
}

func (repository *CategoryRepoImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into mst_category.public.category (name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func ((repository *CategoryRepoImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	//TODO implement me
	panic("implement me")
}

func ((repository *CategoryRepoImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	//TODO implement me
	panic("implement me")
}

func ((repository *CategoryRepoImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) domain.Category {
	//TODO implement me
	panic("implement me")
}

func ((repository *CategoryRepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	//TODO implement me
	panic("implement me")
}
