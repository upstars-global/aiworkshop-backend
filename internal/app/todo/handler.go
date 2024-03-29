package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TaskHandler використовується для обробки HTTP запитів для задач.
type TaskHandler struct {
	service *TaskService
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// responseError представляє собою структуру для відправки помилок.
type responseError struct {
	Error string `json:"error"`
}

// NewTaskHandler створює новий екземпляр TaskHandler.
func NewTaskHandler(service *TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/tasks", h.CreateTask)
	router.GET("/tasks", h.GetAllTasks)
	router.GET("/tasks/:id", h.GetTask)
	router.PUT("/tasks/:id", h.UpdateTask)
	router.PUT("/tasks/:id/check", h.CheckTask)
	router.PUT("/tasks/:id/uncheck", h.UncheckTask)
	router.DELETE("/tasks/:id", h.DeleteTask)
}

// GetAllTasks обробляє запити на отримання усіх задач.
// @Summary Get all tasks
// @Description Get all tasks
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {array} Task
// @Router /tasks [get]
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// CreateTask обробляє створення нової задачі.
// @Summary Create a new task
// @Description Create a new task
// @Tags Task
// @Accept json
// @Produce json
// @Param task body TaskRequest true "Task object"
// @Success 201 {object} Task
// @Failure 400 {object} responseError
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task Task

	var taskRequest TaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.Title = taskRequest.Title
	task.Description = taskRequest.Description
	if err := h.service.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GetTask обробляє запити на отримання задачі за її ідентифікатором.
// @Summary Get a task
// @Description Get a task
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} Task
// @Failure 404 {object} responseError
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTask(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task ID is required"})
		return
	}
	task, err := h.service.GetTaskByID(id)
	if err != nil {
		if err == ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// UpdateTask обробляє оновлення задачі.
// @Summary Update a task
// @Description Update a task
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body TaskRequest true "Task object"
// @Success 200 {object} Task
// @Failure 400 {object} responseError
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task ID is required"})
		return
	}

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		if err == ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	var taskRequest TaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Title = taskRequest.Title
	task.Description = taskRequest.Description

	if err := h.service.UpdateTask(task); err != nil {
		if err == ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask обробляє видалення задачі.
// @Summary Delete a task
// @Description Delete a task
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 204
// @Failure 404 {object} responseError
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task ID is required"})
		return
	}
	if err := h.service.DeleteTask(id); err != nil {
		if err == ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// CheckTask обробляє відмітку задачі як виконаної.
// @Summary Mark a task as completed
// @Description Mark a task as completed
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 204
// @Failure 404 {object} responseError
// @Router /tasks/{id}/check [put]
func (h *TaskHandler) CheckTask(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task ID is required"})
		return
	}
	if err := h.service.CheckTask(id); err != nil {
		if err == ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UncheckTask обробляє відмітку задачі як не виконаної.
// @Summary Mark a task as uncompleted
// @Description Mark a task as uncompleted
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 204
// @Failure 404 {object} responseError
// @Router /tasks/{id}/uncheck [put]
func (h *TaskHandler) UncheckTask(c *gin.Context) {
	id, exist := c.Params.Get("id")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task ID is required"})
		return
	}
	if err := h.service.UncheckTask(id); err != nil {
		if err == ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
