package todos

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type API struct {
	logger     *zap.Logger
	repository Repository
}

func NewRouter(logger *zap.Logger, repository Repository) *gin.Engine {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, Response{
			Error: http.StatusText(http.StatusNotFound),
		})
	})

	api := &API{
		logger:     logger,
		repository: repository,
	}
	todos := router.Group("/v1/todos")
	todos.GET("", api.getTodos)
	todos.GET("/:id", api.getTodo)
	todos.POST("", api.createTodo)
	todos.PUT("/:id", api.updateTodo)
	todos.DELETE("/:id", api.deleteTodo)

	return router
}

type Response struct {
	Data  map[string]any `json:"data,omitempty"`
	Error string         `json:"error,omitempty"`
}

func (a API) getTodos(c *gin.Context) {
	todos, err := a.repository.getTodos()
	if err != nil {
		a.logger.Error("could not retrieve todos from database",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, Response{
			Error: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data: gin.H{
			"todos": todos,
		},
	})
}

func (a API) getTodo(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		a.logger.Error("could not parse uuid",
			zap.String("uuid", c.Param("id")),
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, Response{
			Error: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	todo, err := a.repository.getTodo(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		a.logger.Error("todo with given uuid does not exist",
			zap.String("uuid", id.String()),
			zap.Error(err),
		)
		c.JSON(http.StatusNotFound, Response{
			Error: http.StatusText(http.StatusNotFound),
		})
		return
	}

	if err != nil {
		a.logger.Error("could not retrieve todos from database",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, Response{
			Error: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data: gin.H{
			"todo": todo,
		},
	})
}

type CreateTodoRequest struct {
	Description string `json:"description" binding:"required"`
}

func (a API) createTodo(c *gin.Context) {
	var request CreateTodoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		a.logger.Error("unbindable body received",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, Response{
			Error: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	todo := Todo{
		Description: request.Description,
	}
	todo, err = a.repository.createTodo(todo)
	if err != nil {
		a.logger.Error("could not save todo to database",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, Response{
			Error: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data: gin.H{
			"todo": todo,
		},
	})
}

type UpdateTodoRequest struct {
	Description string     `json:"description" binding:"required"`
	CompletedAt *time.Time `json:"completed_at"`
}

func (a API) updateTodo(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		a.logger.Error("could not parse uuid",
			zap.String("uuid", c.Param("id")),
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, Response{
			Error: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	var request UpdateTodoRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		a.logger.Error("unbindable body received",
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, Response{
			Error: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	todo := Todo{
		Model: Model{
			ID: id,
		},
		Description: request.Description,
		CompletedAt: request.CompletedAt,
	}
	todo, err = a.repository.saveTodo(todo)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		a.logger.Error("todo with given uuid does not exist",
			zap.String("uuid", id.String()),
			zap.Error(err),
		)
		c.JSON(http.StatusNotFound, Response{
			Error: http.StatusText(http.StatusNotFound),
		})
		return
	}

	if err != nil {
		a.logger.Error("could not save todo to database",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, Response{
			Error: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data: gin.H{
			"todo": todo,
		},
	})
}

func (a API) deleteTodo(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		a.logger.Error("could not parse uuid",
			zap.String("uuid", c.Param("id")),
			zap.Error(err),
		)
		c.JSON(http.StatusBadRequest, Response{
			Error: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	err = a.repository.deleteTodo(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		a.logger.Debug("could not delete non-existend todo",
			zap.String("uuid", c.Param("id")),
			zap.Error(err),
		)
		c.JSON(http.StatusNotFound, Response{
			Error: http.StatusText(http.StatusNotFound),
		})
		return
	}

	if err != nil {
		a.logger.Error("could not delete todo from database",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, Response{
			Error: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
