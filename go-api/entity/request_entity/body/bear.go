package body

type PostSendBearBody struct {
	Message string `json:"Message" binding:"required"`
}