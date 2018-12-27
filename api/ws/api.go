package ws

import (
	"encoding/json"

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
	dbRef *mongo.Database,
	reg *registry.Registry,
	caller registry.ClientID,
	jsonMsg json.RawMessage,
) {
	var (
		resp                GenericWSResponse
		recieve             GenericWSRecieve
		sendMsgToAllClients bool
	)

	err := json.Unmarshal(jsonMsg, &recieve)
	if err != nil {
		log.Errorf("ws/API JSON unmarshal error: %v", err)
		return
	}

	switch recieve.Api {
	case "test":
		resp.Message = recieve.Message
		sendMsgToAllClients = true

	default:
		resp.Message = "Default JSON Message!"
	}

	// Response
	resp.Api = recieve.Api
	b, err := json.Marshal(resp)
	if err != nil {
		log.Errorf("ws/API JSON Marshal error: %v", err)
		return
	}

	log.Debugf("ws/API JSON Response: %v", string(b))
	if sendMsgToAllClients {
		reg.SendToAllClients(b)
	} else {
		reg.SendToCaller(caller, b)
	}
}
