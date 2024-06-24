package main

import (
	"TicketBackend/api"
	"TicketBackend/sql"
	"TicketBackend/ticket"
	"github.com/gin-gonic/gin"
)

func main() {
	sql.ConnectSQL()
	ticket.TableInit()

	//gin
	router := gin.Default()
	router.GET("/api/ticket/:ticketCode", api.TicketAPI)
	err2 := router.Run(":8080")
	if err2 != nil {
		panic("Start Router Failed, Error:" + err2.Error())
	}
}
