package repository

import (
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx sql.Tx)
	Update(ctx context.Context, tx sql.Tx)
	Delete(ctx context.Context, tx sql.Tx)
	FindById(ctx context.Context, tx sql.Tx)
	FindAll(ctx context.Context, tx sql.Tx)
}
