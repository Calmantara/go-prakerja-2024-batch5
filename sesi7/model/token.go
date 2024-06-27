package model

type (
	Token struct {
		AuthToken    string `json:"auth_token"`
		SessionToken string `json:"session_token"`
	}
)
