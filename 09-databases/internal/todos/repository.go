package todos

import (
	"errors"
	"slices"
	"sync"
	"time"

	"github.com/google/uuid"
)

var ErrTodoNotFound = errors.New("todo with given UUID does not exist")

type Repository struct {
	mu    sync.Mutex
	todos []Todo
}

func NewRepository() *Repository {
	return &Repository{
		todos: make([]Todo, 0),
	}
}

func (r *Repository) getTodos() (todos []Todo) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.todos
}

func (r *Repository) getTodo(id uuid.UUID) (todo Todo, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, todo := range r.todos {
		if todo.ID == id {
			return todo, nil
		}
	}

	return Todo{}, ErrTodoNotFound
}

func (r *Repository) createTodo(todo Todo) (createdTodo Todo) {
	r.mu.Lock()
	defer r.mu.Unlock()
	todo.ID = uuid.New()
	todo.CreatedAt = time.Now()
	r.todos = append(r.todos, todo)
	return todo
}

func (r *Repository) saveTodo(todo Todo) (savedTodo Todo) {
	r.mu.Lock()
	defer r.mu.Unlock()
	index := slices.IndexFunc(r.todos, func(t Todo) bool {
		return t.ID == todo.ID
	})
	if index == -1 {
		todo.CreatedAt = time.Now()
		r.todos = append(r.todos, todo)
		return todo
	}

	todo.CreatedAt = r.todos[index].CreatedAt
	now := time.Now()
	todo.UpdatedAt = &now
	r.todos[index] = todo
	return todo
}

func (r *Repository) deleteTodo(id uuid.UUID) (err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	index := slices.IndexFunc(r.todos, func(todo Todo) bool {
		return id == todo.ID
	})
	if index == -1 {
		return ErrTodoNotFound
	}

	slice := slices.Delete(r.todos, index, index+1)
	r.todos = slice[:]
	return nil
}
