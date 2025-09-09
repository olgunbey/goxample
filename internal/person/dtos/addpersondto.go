package dtos

type AddPersonResponseDto struct {
	Id           int64  `json:"id"`
	Message      string `json:"message"`
	Successfully bool   `json:"successfully"`
	StatusCode   int    `json:"status_code"`
}

type AddPersonRequestDto struct {
	Name string
	Age  int
}
