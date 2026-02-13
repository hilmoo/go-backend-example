package todo

import "errors"

var ErrNotFound = errors.New("Todo not found")

const (
	ErrTodoNotFound       = "TODO_NOT_FOUND"
	ErrTodoCreationFailed = "TODO_CREATION_FAILED"
	ErrTodoUpdateFailed   = "TODO_UPDATE_FAILED"
	ErrTodoDeleteFailed   = "TODO_DELETE_FAILED"
)

type GetTodoRequest struct {
	ID string `param:"id" json:"id" validate:"required,uuid4" jsonschema:"The UUID of the todo item to be retrieved"`
}

type GetTodoResponse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
	Status  string `json:"status"`
}

type CreateTodoRequest struct {
	Title   string `json:"title" validate:"required,min=1,max=100" jsonschema:"A short summary of the task (1-100 characters)"`
	Details string `json:"details,omitempty" validate:"max=500" jsonschema:"A detailed description of the task (max 500 characters)"`
	Status  string `json:"status" validate:"required,oneof=pending in_progress completed" jsonschema:"The initial state. Allowed: pending, in_progress, completed"`
}

type CreateTodoResponse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
	Status  string `json:"status"`
}

type UpdateTodoRequest struct {
	ID      string `param:"id" json:"id" validate:"required,uuid4" jsonschema:"The UUID of the todo item to be updated"`
	Title   string `json:"title" validate:"required,min=1,max=100" jsonschema:"The updated short summary of the task (1-100 characters)"`
	Details string `json:"details,omitempty" validate:"max=500" jsonschema:"The updated detailed description of the task (max 500 characters)"`
	Status  string `json:"status" validate:"required,oneof=pending in_progress completed" jsonschema:"The updated state. Allowed: pending, in_progress, completed"`
}

type UpdateTodoResponse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
	Status  string `json:"status"`
}

type DeleteTodoRequest struct {
	ID string `param:"id" json:"id" validate:"required,uuid4" jsonschema:"The UUID of the todo item to be permanently deleted"`
}

type DeleteTodoResponse struct {
	Message string `json:"message" jsonschema:"A confirmation message of the operation results"`
}

type ListTodosRequest struct {
	Status string `json:"status,omitempty" validate:"omitempty,oneof=pending in_progress completed" jsonschema:"Filter todos by their status. Allowed: pending, in_progress, completed. If omitted, all todos are returned."`
}

type ListTodosResponse struct {
	Todos []GetTodoResponse `json:"todos" jsonschema:"An array containing the requested todo items"`
	Total int               `json:"total" jsonschema:"The total number of todos matching the current filter"`
}
