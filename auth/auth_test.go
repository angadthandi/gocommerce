package auth

import (
	"encoding/json"
	"fmt"
	"testing"

	log "github.com/angadthandi/gocommerce/log"
)

func TestAuthenticate(t *testing.T) {
	q := AuthRecieve{
		Username: "test",
		Password: "test",
	}

	// Test Request
	req, err := json.Marshal(q)
	if err != nil {
		log.Errorf("test authenticate JSON Marshal error: %v",
			err)
		return
	}

	b, err := Authenticate(req)
	if err != nil {
		log.Errorf("test authenticate error: %v", err)
		return
	}

	var response AuthResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		log.Errorf("test authenticate JSON unmarshal error: %v",
			err)
		return
	}

	fmt.Printf("Authenticate response: %v", response)
}

func TestCreateToken(t *testing.T) {
	userID := 10

	utoken := CreateToken(userID)

	fmt.Printf("CreateToken response: %v", utoken)
}

func BenchmarkCreateToken(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CreateToken(n)
	}
}
