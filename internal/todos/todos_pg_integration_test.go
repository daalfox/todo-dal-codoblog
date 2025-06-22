package todos_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/daalfox/todo-dal-codoblog/internal/todos"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestPostgresRepo(t *testing.T) {
	t.Run("test crud", func(t *testing.T) {
		ctx := context.Background()
		conn := getTestPostgresConn(t, ctx)
		todoRepo := todos.NewPostgresRepo(ctx, conn)

		todo := todos.Todo{
			Title: "Learn to use testcontainers",
		}
		todoId, err := todoRepo.Create(todo)
		assert.NoError(t, err)
		todo.Id = todoId

		todoFromDb, err := todoRepo.Get(todoId)
		assert.NoError(t, err)

		assert.Equal(t, todo, todoFromDb)
	})
}

func getTestPostgresConn(t testing.TB, ctx context.Context) *pgx.Conn {
	t.Helper()
	postgresContainer, err := postgres.Run(
		ctx,
		"postgres:17.5-alpine",
		postgres.WithInitScripts(filepath.Join("..", "..", "migrations", "20250622081024_init.sql")),
		postgres.WithDatabase("test"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	connStr, err := postgresContainer.ConnectionString(ctx)
	if err != nil {
		t.Fatal(err)
	}
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		t.Fatal(err)
	}
	return conn
}
