package auth

import (
	"encoding/json"

	log "github.com/angadthandi/gocommerce/log"
	jwt "github.com/dgrijalva/jwt-go"
)

type AuthRecieve struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func Authenticate(
	jsonMsg json.RawMessage,
) json.RawMessage {
	var (
		ret     json.RawMessage
		resp    AuthResponse
		recieve AuthRecieve
	)

	err := json.Unmarshal(jsonMsg, recieve)
	if err != nil {
		log.Errorf("authenticate JSON unmarshal error: %v", err)
		return nil
	}

	var m jwt.SigningMethod
	resp.Token = jwt.New(m).Raw

	// Response
	ret, err = json.Marshal(resp)
	if err != nil {
		log.Errorf("authenticate JSON Marshal error: %v", err)
		return nil
	}

	return ret
}
