package todos_test

import (
	"context"
	"testing"

	"github.com/daalfox/todo-dal-codoblog/internal/todos"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestPostgresRepo(t *testing.T) {
	t.Run("test crud", func(t *testing.T) {
		ctx := context.Background()
		connStr := ""
		conn, _ := pgx.Connect(ctx, connStr)
		todoRepo := todos.NewPostgresRepo(ctx, conn)

		todo := todos.Todo{
			Title: "Learn to use testcontainers",
		}
		todoId, _ := todoRepo.Create(todo)
		todo.Id = todoId

		todoFromDb, _ := todoRepo.Get(todoId)

		assert.Equal(t, todo, todoFromDb)
	})
}
