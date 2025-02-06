package dto

type ResponseEntity struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
