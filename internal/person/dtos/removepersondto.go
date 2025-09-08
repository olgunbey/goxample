package dtos

type RemovePersonGetByIdResponseDto struct {
	Message      string `json:"message"`
	Successfully bool   `json:"successfully"`
	Id           int    `json:"id"`
	StatusCode   int    `json:"status_code"`
}
