package api

import (
	"TicketBackend/entry"
	"github.com/gin-gonic/gin"
)

func TicketAPI(ctx *gin.Context) {
	err := entry.TicketCheck(ctx.Param("ticketCode"))
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(200, "Checked")
	}
}
