package helpers

type Response struct {
	Errors  []string    `json:"errors"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
