package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/hilmoo/go-backend-example/internal/feature/todo"
	"github.com/hilmoo/go-backend-example/internal/model"
)

type todoRepository struct {
	db *sql.DB
}

var _ todo.DbRepository = (*todoRepository)(nil)

func NewTodoRepository(db *sql.DB) *todoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) Create(ctx context.Context, t *model.Todo) error {
	query := `INSERT INTO todos (id, title, details, status, created_at, updated_at) 
              VALUES (?, ?, ?, ?, ?, ?)`

	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}
	t.UpdatedAt = time.Now()

	_, err := r.db.ExecContext(ctx, query, t.ID, t.Title, t.Details, t.Status, t.CreatedAt, t.UpdatedAt)
	return err
}

func (r *todoRepository) GetByID(ctx context.Context, id string) (*model.Todo, error) {
	query := `SELECT id, title, details, status, created_at, updated_at FROM todos WHERE id = ?`

	var t model.Todo
	if err := sqlscan.Get(ctx, r.db, &t, query, id); err != nil {
		if sqlscan.NotFound(err) {
			return nil, todo.ErrNotFound
		}
		return nil, err
	}
	return &t, nil
}

func (r *todoRepository) Update(ctx context.Context, t *model.Todo) error {
	query := `UPDATE todos SET title=?, details=?, status=?, updated_at=? WHERE id=?`

	t.UpdatedAt = time.Now()

	res, err := r.db.ExecContext(ctx, query, t.Title, t.Details, t.Status, t.UpdatedAt, t.ID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return todo.ErrNotFound
	}

	return nil
}

func (r *todoRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM todos WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *todoRepository) List(ctx context.Context, status string) ([]*model.Todo, error) {
	var query string
	var args []any

	query = `SELECT id, title, details, status, created_at, updated_at FROM todos`
	if status != "" {
		query = query + ` WHERE status = ?`
		args = append(args, status)
	}

	var todos []*model.Todo
	if err := sqlscan.Select(ctx, r.db, &todos, query, args...); err != nil {
		return nil, err
	}

	return todos, nil
}
