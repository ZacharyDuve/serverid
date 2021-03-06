package serverid

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	sIdResourcePath string = "/serverid"
)

type jsonServerId struct {
	ServerId string `json:"server-id"`
}

func GetHandlerFuncFromServerIdService(sIdSvc ServerIdService) (path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	sIdJSON := &jsonServerId{ServerId: sIdSvc.GetServerId().String()}
	return sIdResourcePath, func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		err := encoder.Encode(sIdJSON)
		if err != nil {
			log.Println(err)
		}
	}
}
