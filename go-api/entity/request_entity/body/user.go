package body

type PatchUserStatusBody struct {
	StampId uint `json:"stampId" binding:"required"`
}