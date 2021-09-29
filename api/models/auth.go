package models

type Auth struct {
	Token string `json:"token" binding:"required"`
}

type AuthGoogle struct {
	Email string `json:"email"`
	Exp   string `json:"exp"`
	Error string `json:"error"`
}

type AccessDetails struct {
	AccessUuid string
	Email      string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AuthReq struct {
	AccessToken  string `json:"access_token" `
	RefreshToken string `json:"refresh_token"`
}
