package todos

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PostgresRepo struct {
	dbConn *pgx.Conn
	ctx    context.Context
}

func NewPostgresRepo(ctx context.Context, dbConn *pgx.Conn) PostgresRepo {
	return PostgresRepo{
		ctx:    ctx,
		dbConn: dbConn,
	}
}
func (r *PostgresRepo) Create(t Todo) (int, error) {
	stmt := "INSERT INTO todo (title) VALUES ($1) RETURNING id"
	var id int
	if err := r.dbConn.QueryRow(r.ctx, stmt, t.Title).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
func (r *PostgresRepo) Get(id int) (Todo, error) {
	stmt := "SELECT * FROM todo WHERE id = $1"
	var todo Todo
	err := r.dbConn.QueryRow(r.ctx, stmt, id).Scan(&todo.Id, &todo.Title, &todo.Done)
	if err != nil {
		return todo, err
	}
	return todo, nil
}
func (r *PostgresRepo) Update(id int, updatedTodo Todo) (Todo, error) {
	stmt := "UPDATE todo SET title = $1, done = $2 WHERE id = $3 RETURNING *"
	var todo Todo
	err := r.dbConn.QueryRow(r.ctx, stmt, updatedTodo.Title, updatedTodo.Done, id).Scan(&todo.Id, &todo.Title, &todo.Done)
	if err != nil {
		return todo, err
	}
	return todo, nil
}

type Todo struct {
	Id    int
	Title string
	Done  bool
}
