package routes

import (
	"clicknext-backend/internal/delivery/http/handlers"
	"clicknext-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, h *handlers.Handler) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", h.AuthHandler.Register)
	auth.Post("/login", h.AuthHandler.Login)

	user := api.Group("/users")
	user.Use(middleware.Protected())
	user.Get("/:id", middleware.CheckOwnership(), h.UserHandler.GetUserByID)

	board := api.Group("boards")
	board.Use(middleware.Protected())
	board.Post("/", h.BoardHandler.CreateBoard)
	board.Get("/created", h.BoardHandler.GetMyCreatedBoards)
	board.Get("/joined", h.BoardHandler.GetMyJoinedBoards)
	board.Get("/my", h.BoardHandler.GetAllMyBoards)
	board.Get("/:id", h.BoardHandler.GetBoardByID)

	board.Get("/:id/members", h.BoardMemmberHandler.GetBoardMembers)

	invite := api.Group("/invites")
	invite.Use(middleware.Protected())
	invite.Post("/", h.InviteHandler.CreateInvite)
	invite.Post("/:id/accept", h.InviteHandler.AcceptInvite)
	invite.Get("/:id", h.InviteHandler.GetInviteByID)
	invite.Get("/", h.InviteHandler.GetInvitesByBoardID)

	column := api.Group("/columns")
	column.Use(middleware.Protected())
	column.Post("/", h.ColumnHandler.CreateColumn)
	column.Get("/:id", h.ColumnHandler.GetColumnByID)
	column.Get("/", h.ColumnHandler.GetColumnsByBoardID)
	column.Put("/:id", h.ColumnHandler.UpdateColumn)
	column.Patch("/:id/position", h.ColumnHandler.UpdateColumnPosition)
	column.Delete("/:id", h.ColumnHandler.DeleteColumn)

	task := api.Group("/tasks")
	task.Use(middleware.Protected())
	task.Post("/", h.TaskHandler.CreateTask)
	task.Get("/:id", h.TaskHandler.GetTaskByID)
	task.Get("/", h.TaskHandler.GetTasksByColumnID)
	task.Put("/:id", h.TaskHandler.UpdateTask)
	task.Patch("/:id/position", h.TaskHandler.UpdateTaskPosition)
	task.Delete("/:id", h.TaskHandler.DeleteTask)

	assignment := api.Group("/assignments")
	assignment.Use(middleware.Protected())
	assignment.Post("/", h.AssigneeTaskHandler.AssignTask)
	assignment.Delete("/", h.AssigneeTaskHandler.UnassignTask)
	assignment.Get("/task", h.AssigneeTaskHandler.GetAssigneesByTaskID)
	assignment.Get("/assignee", h.AssigneeTaskHandler.GetTasksByAssigneeID)

	tag := api.Group("/tags")
	tag.Use(middleware.Protected())
	tag.Post("/", h.TagHandler.CreateTag)
	tag.Get("/:id", h.TagHandler.GetTagByID)
	tag.Get("/", h.TagHandler.GetAllTags)
	tag.Get("/by-user", h.TagHandler.GetTagsByUser)
	tag.Put("/:id", h.TagHandler.UpdateTag)
	tag.Delete("/:id", h.TagHandler.DeleteTag)

	taskTags := api.Group("/tasks/tags")
	taskTags.Use(middleware.Protected())
	taskTags.Post("/", h.TaskTagHandler.CreateTaskTag)
	taskTags.Get("/:id", h.TaskTagHandler.GetTaskTagByID)
	taskTags.Get("/by-task", h.TaskTagHandler.GetTaskTagsByTaskID)
	taskTags.Get("/by-tag", h.TaskTagHandler.GetTaskTagsByTagID)
	taskTags.Delete("/:id", h.TaskTagHandler.DeleteTaskTag)
	taskTags.Delete("/", h.TaskTagHandler.DeleteTaskTagByTaskAndTag)
	taskTags.Delete("/by-task", h.TaskTagHandler.DeleteTaskTagsByTaskID)
}
