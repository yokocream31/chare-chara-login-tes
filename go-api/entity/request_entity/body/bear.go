package body

import (
	"time"
)

type PostSendBearBody struct {
	Message string `json:"Message" binding:"required"`
}

type GetHistoryBody struct {
	Start time.Time `json:"Start"`
}