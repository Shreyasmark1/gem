package store

import (
	"context"
	"database/sql"
	"gem/internal/store/repository"
	"gem/pkg/response"
)

type Store struct {
	db      *sql.DB
	Queries *repository.Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: repository.New(db),
	}
}

func WithTransaction[T any](ctx context.Context, s *Store, fn func(*repository.Queries) (T, error)) (T, error) {
	var result T

	tx, err := s.db.BeginTx(ctx, nil)

	if err != nil {
		return result, err
	}

	q := repository.New(tx)
	result, err = fn(q)

	if err != nil {

		if rbErr := tx.Rollback(); rbErr != nil {
			return result, response.InternalServerError("Failed to commit transaction", err)
		}

		return result, err
	}

	if err := tx.Commit(); err != nil {
		return result, err
	}

	return result, nil
}
