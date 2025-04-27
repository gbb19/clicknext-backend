package http

import (
	"clicknext-backend/internal/delivery/http/dto"
	"clicknext-backend/internal/usecase"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskUsecase *usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{taskUsecase: taskUsecase}
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a task in a column
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body dto.TaskCreateRequest true "Task"
// @Success 201 {object} dto.TaskResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	// Get userID from Locals
	userID := c.Locals("userID").(uint)
	// Parse request body
	var taskRequest dto.TaskCreateRequest
	if err := c.BodyParser(&taskRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Set the CreatedBy field to the current user ID
	taskRequest.CreatedBy = userID

	// Convert to domain model
	task := taskRequest.ToTaskDomain()

	// Create the task
	if err := h.taskUsecase.CreateTask(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create task"})
	}

	// Return the created task
	return c.Status(fiber.StatusCreated).JSON(dto.FromTaskDomain(task))
}

// GetTaskByID godoc
// @Summary Get task by ID
// @Description Retrieve task details by task ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} dto.TaskResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task id"})
	}

	task, err := h.taskUsecase.GetTaskByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.FromTaskDomain(task))
}

// GetTasksByColumnID godoc
// @Summary Get all tasks for a column
// @Description Retrieve all tasks associated with a column
// @Tags tasks
// @Accept json
// @Produce json
// @Param column_id query int true "Column ID"
// @Success 200 {array} dto.TaskResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /tasks [get]
func (h *TaskHandler) GetTasksByColumnID(c *fiber.Ctx) error {
	columnID, err := strconv.Atoi(c.Query("column_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid column id"})
	}

	tasks, err := h.taskUsecase.GetTasksByColumnID(uint(columnID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no tasks found"})
	}

	var taskResponses []dto.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, *dto.FromTaskDomain(task))
	}

	return c.Status(fiber.StatusOK).JSON(taskResponses)
}

// UpdateTask godoc
// @Summary Update a task
// @Description Update a task's details
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body dto.TaskUpdateRequest true "Task"
// @Success 200 {object} dto.TaskResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task id"})
	}

	// Get the existing task
	task, err := h.taskUsecase.GetTaskByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
	}

	// Parse the request
	var updateRequest struct {
		Name      string    `json:"name"`
		DueDate   time.Time `json:"due_date"`
		StartDate time.Time `json:"start_date"`
	}
	if err := c.BodyParser(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Update fields
	if updateRequest.Name != "" {
		task.Name = updateRequest.Name
	}
	if !updateRequest.DueDate.IsZero() {
		task.DueDate = updateRequest.DueDate
	}
	if !updateRequest.StartDate.IsZero() {
		task.StartDate = updateRequest.StartDate
	}
	log.Println(task.TaskID)
	log.Println(task.Position)
	log.Println(task.ColumnID)
	log.Println(task.Name)
	log.Println(task.StartDate)
	log.Println(task.DueDate)
	log.Println(task.CreatedBy)

	// Save the updated task
	if err := h.taskUsecase.UpdateTask(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not update task"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.FromTaskDomain(task))
}

// UpdateTaskPosition godoc
// @Summary Update a task's position
// @Description Move a task to a new position within its column
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param position body PositionUpdateRequest true "New Position"
// @Success 200 {object} dto.TaskResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /tasks/{id}/position [patch]
func (h *TaskHandler) UpdateTaskPosition(c *fiber.Ctx) error {
	taskID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task id"})
	}

	// Parse the request
	var updateRequest struct {
		Position int  `json:"position"`
		ColumnID uint `json:"column_id"`
	}
	if err := c.BodyParser(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Update the position
	if err := h.taskUsecase.UpdateTaskPosition(uint(taskID), updateRequest.Position, updateRequest.ColumnID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not update task position"})
	}

	// Get the updated task
	task, err := h.taskUsecase.GetTaskByID(uint(taskID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.FromTaskDomain(task))
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 204 {object} nil
// @Failure 404 {object} dto.ErrorResponse
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid task id"})
	}

	if err := h.taskUsecase.DeleteTask(uint(id)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
