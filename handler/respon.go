package handler

type ResponseData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Id      interface{} `json:"id"`
	Data    interface{} `json:"data"`
	Count 	int64 `json:"count"`
}
