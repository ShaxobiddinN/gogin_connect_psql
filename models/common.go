package models

type JSONResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type JSONErrorResponce struct {
	Error string `json:"error"`
}