package auth

import (
	"encoding/json"
	"time"

	"github.com/angadthandi/gocommerce/config"
	log "github.com/angadthandi/gocommerce/log"
	jwt "github.com/dgrijalva/jwt-go"
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

func CreateToken(userID int) UserToken {
	var ret UserToken

	expiresAt := time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix()

	// Create a new token object,
	// specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userid": userID,
			"nbf":    expiresAt,
		},
	)

	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err := token.SignedString(
		[]byte(config.JWTAuthSecret))
	if err != nil {
		log.Errorf("unable to create token: %v", err)
		return ret
	}

	ret.Token = tokenStr
	ret.ExpiresAt = int(expiresAt)

	return ret
}
