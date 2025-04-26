package dto

// LoginRequest สำหรับรับข้อมูลเข้าสู่ระบบ
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// AuthResponse สำหรับส่งข้อมูลผู้ใช้และ token กลับ
type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
