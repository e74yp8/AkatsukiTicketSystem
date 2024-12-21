package api

import (
	"TicketBackend/entry"
	"TicketBackend/ticket"
	"github.com/gin-gonic/gin"
)

func TicketCheckAPI(ctx *gin.Context) {
	err := entry.TicketCheck(ctx.Param("ticketCode"))
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(200, "Checked")
	}
}

func TicketCreateAPI(ctx *gin.Context) {
	err := ticket.Create(ctx.Param("ticketCode"))
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(200, "Created")
	}
}
