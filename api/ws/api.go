package ws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/angadthandi/gocommerce/auth"
	"github.com/angadthandi/gocommerce/log"
	"github.com/angadthandi/gocommerce/registry"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type GenericWSRecieve struct {
	Api     string          `json:"api"`
	Message json.RawMessage `json:"message"`
}

type GenericWSResponse struct {
	Api     string      `json:"api"`
	Message interface{} `json:"message"`
}

// handler for ws/API
func API(
	w http.ResponseWriter,
	r *http.Request,
	dbRef *mongo.Database,
	reg *registry.Registry,
	caller registry.ClientID,
	jsonMsg json.RawMessage,
) {
	var (
		resp    GenericWSResponse
		recieve GenericWSRecieve
	)

	err := json.Unmarshal(jsonMsg, recieve)
	if err != nil {
		log.Errorf("rest/API JSON unmarshal error: %v", err)
		return
	}

	switch recieve.Api {
	case "auth":
		resp.Message = auth.Authenticate(recieve.Message)

	default:
		resp.Message = "Default JSON Message!"
	}

	// Response
	resp.Api = recieve.Api
	b, err := json.Marshal(resp)
	if err != nil {
		log.Errorf("rest/API JSON Marshal error: %v", err)
		return
	}

	log.Debugf("rest/API JSON Response: %v", string(b))
	fmt.Fprintf(w, "rest/API JSON Response: %v", string(b))
}
