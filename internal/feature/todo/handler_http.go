package todo

import (
	"net/http"

	errort "github.com/hilmoo/go-backend-example/internal/transport/error"
	"github.com/hilmoo/go-backend-example/internal/transport/validation"

	"github.com/labstack/echo/v5"
)

type httpHandler struct {
	validate *validation.Vld
	db       DbRepository
}

func NewHTTPHandler(vld *validation.Vld, db DbRepository) *httpHandler {
	return &httpHandler{
		validate: vld,
		db:       db,
	}
}

func (h *httpHandler) Register(e *echo.Group) {
	todoGroup := e.Group("/todos")

	todoGroup.GET("", h.listTodos)
	todoGroup.POST("", h.createTodo)
	todoGroup.GET("/:id", h.getTodo)
	todoGroup.PUT("/:id", h.updateTodo)
	todoGroup.DELETE("/:id", h.deleteTodo)
}

// getTodo godoc
// @Summary      Get a todo item by ID
// @Description  Retrieves a todo item by its unique ID.
// @Tags         todos
// @Param        id   path      string  true  "Todo ID"
// @Success      200  {object}  GetTodoResponse
// @Failure      400  {object}  error.HttpErrorResponse
// @Failure      404  {object}  error.HttpErrorResponse
// @Router       /todos/{id} [get]
func (h *httpHandler) getTodo(c *echo.Context) error {
	payload, err := validation.BindValidatePayload[GetTodoRequest](c, h.validate)
	if err != nil {
		return errort.HttpError(c, err)
	}

	resp, err := getTodo(c.Request().Context(), payload, h.db)
	if err != nil {
		return errort.HttpError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// createTodo godoc
// @Summary      Create a new todo item
// @Description  Creates a new todo item with the provided details.
// @Tags         todos
// @Accept       json
// @Param        todo  body      CreateTodoRequest  true  "Todo to create"
// @Success      201   {object}  CreateTodoResponse
// @Failure      400   {object}  error.HttpErrorResponse
// @Router       /todos [post]
func (h *httpHandler) createTodo(c *echo.Context) error {
	payload, err := validation.BindValidatePayload[CreateTodoRequest](c, h.validate)
	if err != nil {
		return errort.HttpError(c, err)
	}

	resp, err := createTodo(c.Request().Context(), payload, h.db)
	if err != nil {
		return errort.HttpError(c, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

// updateTodo godoc
// @Summary      Update an existing todo item
// @Description  Updates the details of an existing todo item by its ID.
// @Tags         todos
// @Accept       json
// @Param        id    path      string             true  "Todo ID"
// @Param        todo  body      UpdateTodoRequest  true  "Updated todo details"
// @Success      200   {object}  UpdateTodoResponse
// @Failure      400   {object}  error.HttpErrorResponse
// @Failure      404   {object}  error.HttpErrorResponse
// @Router       /todos/{id} [put]
func (h *httpHandler) updateTodo(c *echo.Context) error {
	payload, err := validation.BindValidatePayload[UpdateTodoRequest](c, h.validate)
	if err != nil {
		return errort.HttpError(c, err)
	}

	resp, err := updateTodo(c.Request().Context(), payload, h.db)
	if err != nil {
		return errort.HttpError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// deleteTodo godoc
// @Summary      Delete a todo item
// @Description  Deletes a todo item by its unique ID.
// @Tags         todos
// @Param        id   path      string  true  "Todo ID"
// @Success      200  {object}  DeleteTodoResponse
// @Failure      400  {object}  error.HttpErrorResponse
// @Failure      404  {object}  error.HttpErrorResponse
// @Router       /todos/{id} [delete]
func (h *httpHandler) deleteTodo(c *echo.Context) error {
	payload, err := validation.BindValidatePayload[DeleteTodoRequest](c, h.validate)
	if err != nil {
		return errort.HttpError(c, err)
	}

	resp, err := deleteTodo(c.Request().Context(), payload, h.db)
	if err != nil {
		return errort.HttpError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// listTodos godoc
// @Summary      List todo items
// @Description  Retrieves a list of todo items, optionally filtered by status.
// @Tags         todos
// @Param        status  query     string  false  "Filter by status"  Enums(pending, in_progress, completed)
// @Success      200     {object}  ListTodosResponse
// @Failure      400     {object}  error.HttpErrorResponse
// @Router       /todos [get]
func (h *httpHandler) listTodos(c *echo.Context) error {
	payload, err := validation.BindValidatePayload[ListTodosRequest](c, h.validate)
	if err != nil {
		return errort.HttpError(c, err)
	}

	resp, err := listTodos(c.Request().Context(), payload, h.db)
	if err != nil {
		return errort.HttpError(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}
