package models

type Response struct {
	Data    string `json:"Data"`
	Message string `json:"Message"`
	Status  int    `json:"Status"`
	Success bool   `json:"Success"`
}
