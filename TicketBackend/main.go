package main

import (
	"TicketBackend/api"
	"TicketBackend/sql"
	"TicketBackend/ticket"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	sql.ConnectSQL()
	//ticket.TableInit()

	ticket.Create("1234567890")
	status, err := ticket.CheckStatus("1234567890")
	if err != nil {
		fmt.Printf("Check Status Failed, Error: %v\n", err.Error())
	}
	fmt.Printf("Ticket Status: %v\n", status)

	//gin
	router := gin.Default()
	router.GET("/api/ticket/:ticketCode", api.TicketAPI)
	err2 := router.Run(":8080")
	if err != nil {
		panic("Start Router Failed, Error:" + err2.Error())
	}
}
