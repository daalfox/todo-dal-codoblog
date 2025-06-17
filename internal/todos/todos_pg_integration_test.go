package todos_test

import (
	"testing"

	"github.com/daalfox/todo-dal-codoblog/internal/todos"
	"github.com/stretchr/testify/assert"
)

func TestPostgresRepo(t *testing.T) {
	t.Run("test crud", func(t *testing.T) {
		todoRepo := todos.NewPostgresRepo()
		todo := todos.Todo{
			Title: "Learn to use testcontainers",
		}
		todoId, _ := todoRepo.Create(todo)
		todo.Id = todoId

		todoFromDb, _ := todoRepo.Get(todoId)

		assert.Equal(t, todo, todoFromDb)
	})
}
