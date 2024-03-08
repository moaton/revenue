package dto

type Response struct {
	ErrorCode   string      `json:"errorCode"`
	Description string      `json:"description"`
	Payload     interface{} `json:"payload"`
}
