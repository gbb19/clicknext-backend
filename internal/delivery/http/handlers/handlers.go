package handlers

import (
	handler "clicknext-backend/internal/delivery/http"
	"clicknext-backend/internal/repository/postgres"
	"clicknext-backend/internal/usecase"

	"gorm.io/gorm"
)

type Handler struct {
	UserHandler         *handler.UserHandler
	AuthHandler         *handler.AuthHandler
	BoardHandler        *handler.BoardHandler
	BoardMemmberHandler *handler.BoardMemberHandler
	InviteHandler       *handler.InviteHandler
	ColumnHandler       *handler.ColumnHandler
	TaskHandler         *handler.TaskHandler
	AssigneeTaskHandler *handler.AssigneeTaskHandler
	TagHandler          *handler.TagHandler
}

func NewHandler(db *gorm.DB) *Handler {
	userRepo := postgres.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	authUsecase := usecase.NewAuthUseCase(userRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	boardRepo := postgres.NewBoardRepository(db)
	boardUsecase := usecase.NewBoardUsecase(boardRepo)
	boardHandler := handler.NewBoardHandler(boardUsecase)

	boardMemberRepo := postgres.NewBoardMemberRepository(db)
	boardMemberUsecase := usecase.NewBoardMemberUsecase(boardMemberRepo)
	boardMemberHandler := handler.NewBoardMemberHandler(boardMemberUsecase)

	inviteRepo := postgres.NewInviteRepository(db)
	inviteUsecase := usecase.NewInviteUsecase(inviteRepo, boardMemberRepo)
	inviteHandler := handler.NewInviteHandler(inviteUsecase)

	columnRepo := postgres.NewColumnRepository(db)
	columnUsecase := usecase.NewColumnUsecase(columnRepo)
	columnHandler := handler.NewColumnHandler(columnUsecase)

	taskRepo := postgres.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	assigneeTaskRepo := postgres.NewAssigneeTaskRepository(db)
	assigneeTaskUsecase := usecase.NewAssigneeTaskUsecase(assigneeTaskRepo)
	assigneeTaskHandler := handler.NewAssigneeTaskHandler(assigneeTaskUsecase)

	tagRepo := postgres.NewTagRepository(db)
	tagUsecase := usecase.NewTagUsecase(tagRepo)
	tagHandler := handler.NewTagHandler(tagUsecase)

	return &Handler{
		UserHandler:         userHandler,
		AuthHandler:         authHandler,
		BoardHandler:        boardHandler,
		BoardMemmberHandler: boardMemberHandler,
		InviteHandler:       inviteHandler,
		ColumnHandler:       columnHandler,
		TaskHandler:         taskHandler,
		AssigneeTaskHandler: assigneeTaskHandler,
		TagHandler:          tagHandler,
	}
}
