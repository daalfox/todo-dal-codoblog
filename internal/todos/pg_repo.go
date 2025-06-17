package todos

type PostgresRepo struct {
}

func NewPostgresRepo() PostgresRepo {
	return PostgresRepo{}
}
func (r *PostgresRepo) Create(t Todo) (int, error) {
	// TODO
	return 0, nil
}
func (r *PostgresRepo) Get(id int) (Todo, error) {
	// TODO
	return Todo{}, nil
}

type Todo struct {
	Id    int
	Title string
	Done  bool
}
