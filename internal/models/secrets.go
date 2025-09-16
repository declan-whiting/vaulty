package models

type SecretModel struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type SecretListResponse struct {
	Secrets []SecretModel
}
