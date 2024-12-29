package main

import (
	"TicketBackend/api"
	"TicketBackend/sql"
	"TicketBackend/ticket"
	"TicketBackend/timeSlot"
	"github.com/gin-gonic/gin"
)

func main() {
	sql.ConnectSQL()
	ticket.TableInit()
	timeSlot.TableInit()
	//var loc, _ = time.LoadLocation("Asia/Hong_Kong")
	//timeSlot.CountSet(time.Date(2025, 1, 19, 12, 0, 0, 0, loc), 10000)

	//gin
	router := gin.Default()
	router.GET("/ticket/check/:ticketCode", api.TicketCheckAPI)
	router.GET("/ticket/check_exist/:ticketCode", api.TicketCheckExistAPI)
	router.POST("/ticket/create", api.TicketCreateAPI)

	err2 := router.Run(":8080")
	if err2 != nil {
		panic("Start Router Failed, Error:" + err2.Error())
	}

}
