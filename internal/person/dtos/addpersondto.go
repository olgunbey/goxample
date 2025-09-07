package dtos

type AddPersonResponseDto struct {
	Id           int64  `json:"id"`
	Message      string `json:"message"`
	Successfully bool   `json:"successfully"`
}

type AddPersonRequestDto struct {
	Name string
	Age  int
}
