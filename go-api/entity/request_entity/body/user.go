package body

type PatchUserStatusBody struct {
	StampId string `json:"stampId" binding:"required"`
}