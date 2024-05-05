package todos

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type API struct {
	repository *Repository
}

func NewRouter(repository *Repository) *http.ServeMux {
	mux := http.NewServeMux()
	api := &API{
		repository: repository,
	}
	mux.HandleFunc("GET /v1/todos", api.getTodos)
	mux.HandleFunc("GET /v1/todos/{id}", api.getTodo)
	mux.HandleFunc("POST /v1/todos", api.createTodo)
	mux.HandleFunc("PUT /v1/todos/{id}", api.updateTodo)
	mux.HandleFunc("DELETE /v1/todos/{id}", api.deleteTodo)
	return mux
}

type Response struct {
	Data  map[string]any `json:"data,omitempty"`
	Error string         `json:"error,omitempty"`
}

func ResponseErrorBytes(httpCode int) []byte {
	response := Response{
		Error: http.StatusText(httpCode),
	}
	bytes, _ := json.Marshal(response) // ignores error for convenience
	return bytes
}

func (a API) getTodos(w http.ResponseWriter, r *http.Request) {
	todos := a.repository.getTodos()
	response := Response{
		Data: map[string]any{
			"todos": todos,
		},
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		code := http.StatusInternalServerError
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	w.Write(bytes)
}

func (a API) getTodo(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		slog.Error("could not parse uuid",
			"uuid", r.PathValue("id"),
			"error", err,
		)
		code := http.StatusBadRequest
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	todo, err := a.repository.getTodo(id)
	if errors.Is(err, ErrTodoNotFound) {
		slog.Error("todo with given uuid does not exist",
			"uuid", id.String(),
			"error", err,
		)
		code := http.StatusNotFound
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	if err != nil {
		slog.Error("could not retrieve todos from database",
			"error", err,
		)
		code := http.StatusInternalServerError
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	response := Response{
		Data: map[string]any{
			"todo": todo,
		},
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		code := http.StatusInternalServerError
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	w.Write(bytes)
}

type CreateTodoRequest struct {
	Description string `json:"description" binding:"required"`
}

func (a API) createTodo(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		code := http.StatusInternalServerError
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	defer body.Close()
	var request CreateTodoRequest
	err = json.Unmarshal(bodyBytes, &request)
	if err != nil {
		slog.Error("unbindable body received",
			"error", err,
		)
		code := http.StatusBadRequest
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	todo := Todo{
		Description: request.Description,
	}
	todo = a.repository.createTodo(todo)
	response := Response{
		Data: map[string]any{
			"todo": todo,
		},
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		code := http.StatusInternalServerError
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(bytes)
}

type UpdateTodoRequest struct {
	Description string     `json:"description" binding:"required"`
	CompletedAt *time.Time `json:"completed_at"`
}

func (a API) updateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		slog.Error("could not parse uuid",
			"uuid", r.PathValue("id"),
			"error", err,
		)
		code := http.StatusBadRequest
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	body := r.Body
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		code := http.StatusInternalServerError
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	defer body.Close()
	var request UpdateTodoRequest
	err = json.Unmarshal(bodyBytes, &request)
	if err != nil {
		slog.Error("unbindable body received",
			"error", err,
		)
		code := http.StatusBadRequest
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	todo := Todo{
		ID:          id,
		Description: request.Description,
		CompletedAt: request.CompletedAt,
	}
	todo = a.repository.saveTodo(todo)
	response := Response{
		Data: map[string]any{
			"todo": todo,
		},
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		code := http.StatusInternalServerError
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (a API) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		slog.Error("could not parse uuid",
			"uuid", r.PathValue("id"),
			"error", err,
		)
		code := http.StatusBadRequest
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	err = a.repository.deleteTodo(id)
	if errors.Is(err, ErrTodoNotFound) {
		code := http.StatusNotFound
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	if err != nil {
		code := http.StatusInternalServerError
		w.WriteHeader(code)
		w.Write(ResponseErrorBytes(code))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
