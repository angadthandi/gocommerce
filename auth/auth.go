package auth

import (
	"encoding/json"

	log "github.com/angadthandi/gocommerce/log"
)

type AuthRecieve struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	UserToken
}

type UserToken struct {
	Token     string `json:"token"`
	ExpiresAt int    `json:"expiresAt"`
}

func Authenticate(
	jsonMsg json.RawMessage,
) (json.RawMessage, error) {
	var (
		ret     json.RawMessage
		resp    AuthResponse
		recieve AuthRecieve
		err     error
		userID  int
	)

	err = json.Unmarshal(jsonMsg, &recieve)
	if err != nil {
		log.Errorf("authenticate JSON unmarshal error: %v", err)
		return nil, err
	}

	userID = ValidateDBUser(recieve.Username, recieve.Password)

	if userID != 0 {
		resp.UserToken = CreateToken(userID)
	}

	// Response
	ret, err = json.Marshal(resp)
	if err != nil {
		log.Errorf("authenticate JSON Marshal error: %v", err)
		return nil, err
	}

	return ret, err
}

func ValidateDBUser(
	username string,
	password string,
) int {
	var userID int

	// TODO check DB
	// recieve.Username
	// recieve.Password

	// STUB
	userID = 1

	return userID
}
