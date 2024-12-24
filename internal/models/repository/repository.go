package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Repository[T any] interface {
	Create(ctx context.Context, entity *T) error
	GetByID(ctx context.Context, id string) (*T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]T, error)
}

type BaseRepository[T any] struct {
	db        *sqlx.DB
	tableName string
	fields    []string
}

func NewBaseRepository[T any](db *sqlx.DB, tablename string, fields []string) *BaseRepository[T] {
	return &BaseRepository[T]{
		db:        db,
		tableName: tablename,
		fields:    fields,
	}
}

// buildInsertQuery constructs an INSERT query with named parameters
func (r *BaseRepository[T]) buildInsertQuery() string {
	placeholders := make([]string, len(r.fields))
	for i, field := range r.fields {
		placeholders[i] = fmt.Sprintf(":%s", field)
	}

	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING *",
		r.tableName,
		strings.Join(r.fields, ", "),
		strings.Join(placeholders, ", "),
	)
}

// buildUpdateQuery constructs an UPDATE query with named parameters
func (r *BaseRepository[T]) buildUpdateQuery() string {
	setClauses := make([]string, len(r.fields)-1) // Excluding ID
	for i, field := range r.fields[1:] {          // Skip ID field
		setClauses[i] = fmt.Sprintf("%s = :%s", field, field)
	}

	return fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = :id RETURNING *",
		r.tableName,
		strings.Join(setClauses, ", "),
	)
}

func (r *BaseRepository[T]) Create(ctx context.Context, entity *T) error {
	query := r.buildInsertQuery()
	return r.db.GetContext(ctx, entity, query, entity)
}

func (r *BaseRepository[T]) GetById(ctx context.Context, id string) (*T, error) {
	var entity T
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", r.tableName)

	err := r.db.GetContext(ctx, &entity, query, id)

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *BaseRepository[T]) Update(ctx context.Context, entity *T) error {
	query := r.buildUpdateQuery()
	return r.db.GetContext(ctx, entity, query, entity)
}

func (r *BaseRepository[T]) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", r.tableName)
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *BaseRepository[T]) List(ctx context.Context, limit, offset int) ([]T, error) {
	var entities []T
	query := fmt.Sprintf("SELECT * FROM %s LIMIT $1 OFFSET $2", r.tableName)
	err := r.db.SelectContext(ctx, &entities, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return entities, nil
}
