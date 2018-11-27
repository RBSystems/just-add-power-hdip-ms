package main

import (
	"net/http"
	"os"

	"github.com/byuoitav/common"
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/v2/auth"
	"github.com/byuoitav/just-add-power-hdip-ms/handlers"
)

func main() {
	port := ":8022"
	router := common.NewRouter()

	log.L.Debugf("Tied to a room system: %v", os.Getenv("ROOM_SYSTEM"))

	write := router.Group("", auth.AuthorizeRequest("write-state", "room", auth.LookupResourceFromAddress))
	read := router.Group("", auth.AuthorizeRequest("read-state", "room", auth.LookupResourceFromAddress))

	//Functionality endpoints
	write.GET("/input/:transmitter/:receiver", handlers.SetReceiverToTransmissionChannel)

	//Status endpoints
	read.GET("/input/get/:address", handlers.GetTransmissionChannel)

	//Configuration endpoints
	write.PUT("/configure/:transmitter", handlers.SetTransmitterChannel)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
