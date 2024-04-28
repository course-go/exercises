package todos

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewRepository(logger *zap.Logger, db *gorm.DB) Repository {
	db.AutoMigrate(&Todo{})
	return Repository{
		logger: logger,
		db:     db,
	}
}

func (r Repository) getTodos() (todos []Todo, err error) {
	result := r.db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func (r Repository) getTodo(id uuid.UUID) (todo Todo, err error) {
	result := r.db.First(&todo, id)
	if result.Error != nil {
		return Todo{}, result.Error
	}

	return todo, nil
}

func (r Repository) createTodo(todo Todo) (createdTodo Todo, err error) {
	result := r.db.Create(&todo)
	if result.Error != nil {
		return Todo{}, result.Error
	}

	return todo, nil
}

func (r Repository) saveTodo(todo Todo) (savedTodo Todo, err error) {
	result := r.db.First(&savedTodo, todo.ID)
	if result.Error != nil {
		return Todo{}, result.Error
	}

	savedTodo.Description = todo.Description
	savedTodo.CompletedAt = todo.CompletedAt
	result = r.db.Save(&savedTodo)
	if result.Error != nil {
		return Todo{}, result.Error
	}

	return savedTodo, nil
}

func (r Repository) deleteTodo(id uuid.UUID) (err error) {
	result := r.db.Delete(&Todo{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
