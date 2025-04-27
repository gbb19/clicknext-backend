package http

import (
	"clicknext-backend/internal/delivery/http/dto"
	"clicknext-backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type InviteHandler struct {
	inviteUsecase *usecase.InviteUsecase
}

func NewInviteHandler(inviteUsecase *usecase.InviteUsecase) *InviteHandler {
	return &InviteHandler{inviteUsecase: inviteUsecase}
}

// CreateInvite godoc
// @Summary Create a new invite
// @Description Create an invite to a board
// @Tags invites
// @Accept json
// @Produce json
// @Param invite body dto.InviteCreateRequest true "Invite"
// @Success 201 {object} dto.InviteResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /invites [post]
func (h *InviteHandler) CreateInvite(c *fiber.Ctx) error {
	// ดึง userID จาก Locals
	userID := c.Locals("userID").(uint)

	// แปลงข้อมูลจาก request
	var inviteRequest dto.InviteCreateRequest
	if err := c.BodyParser(&inviteRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// ใช้ userID ที่ดึงมาจาก Locals แทน inviterID ในคำเชิญ
	invite := inviteRequest.ToInviteDomain()
	invite.InviterID = userID // ตั้งค่าผู้สร้างคำเชิญเป็น user ที่ล็อกอิน

	// สร้างคำเชิญในฐานข้อมูล
	if err := h.inviteUsecase.CreateInvite(invite); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create invite"})
	}

	// ส่งข้อมูลคำเชิญที่ถูกสร้างกลับไป
	return c.Status(fiber.StatusCreated).JSON(dto.FromInviteDomain(invite))
}

// GetInviteByID godoc
// @Summary Get invite by ID
// @Description Retrieve invite details by invite ID
// @Tags invites
// @Accept json
// @Produce json
// @Param id path int true "Invite ID"
// @Success 200 {object} dto.InviteResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /invites/{id} [get]
func (h *InviteHandler) GetInviteByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid invite id"})
	}

	invite, err := h.inviteUsecase.GetInviteByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "invite not found"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.FromInviteDomain(invite))
}

// GetInvitesByBoardID godoc
// @Summary Get all invites for a board
// @Description Retrieve all invites associated with a board
// @Tags invites
// @Accept json
// @Produce json
// @Param board_id path int true "Board ID"
// @Success 200 {array} dto.InviteResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /invites/board/{board_id} [get]
func (h *InviteHandler) GetInvitesByBoardID(c *fiber.Ctx) error {
	boardID, err := c.ParamsInt("board_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid board id"})
	}

	invites, err := h.inviteUsecase.GetInvitesByBoardID(uint(boardID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no invites found"})
	}

	var inviteResponses []dto.InviteResponse
	for _, invite := range invites {
		inviteResponses = append(inviteResponses, *dto.FromInviteDomain(invite))
	}

	return c.Status(fiber.StatusOK).JSON(inviteResponses)
}

func (h *InviteHandler) AcceptInvite(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	// ดึง inviteID จาก path parameter
	inviteID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid invite id"})
	}

	// ยอมรับคำเชิญและเพิ่มสมาชิกใน board
	if err := h.inviteUsecase.AcceptInvite(uint(inviteID), userID); err != nil {
		if err.Error() == "not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "invite not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not accept invite"})
	}

	// ส่งข้อมูล invite ที่อัปเดตสถานะแล้วกลับไป
	invite, _ := h.inviteUsecase.GetInviteByID(uint(inviteID))
	return c.Status(fiber.StatusOK).JSON(dto.FromInviteDomain(invite))
}
