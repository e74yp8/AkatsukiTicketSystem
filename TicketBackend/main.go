package main

import (
	"TicketBackend/api"
	"TicketBackend/config"
	"TicketBackend/sql"
	"TicketBackend/ticket"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func main() {
	sql.ConnectSQL()
	ticket.TableInit()

	cfg := config.LoadConfig()

	//gin
	router := gin.Default()
	router.GET("/ticket/check/:ticketCode", api.TicketCheckAPI)
	router.GET("/ticket/check_exist/:ticketCode", api.TicketCheckExistAPI)
	router.GET("/ticket/create/:ticketCode/:email", api.TicketCreateAPI)
	router.GET("/ticket/checkid", api.CheckLimitIdAPI)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	err2 := router.Run(":" + strconv.Itoa(cfg.ServicePort))
	if err2 != nil {
		panic("Start Router Failed, Error:" + err2.Error())
	}
}
