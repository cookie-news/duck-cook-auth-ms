package api_helper

type HTTPError struct {
	Code    int    `json:"code" example:"400" format:"int64"`
	Message string `json:"message" example:"status bad request"`
}