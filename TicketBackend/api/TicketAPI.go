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

//func TicketCreateAPI(ctx *gin.Context) {
//	t := entry.TimeSelect()
//	err := ticket.Create(ctx.Param("ticketCode"), t)
//	if err != nil {
//		ctx.JSON(400, err.Error())
//	} else {
//		ctx.JSON(200, t)
//	}
//}

func TicketCheckExistAPI(ctx *gin.Context) {
	exist, err := ticket.CheckExist(ctx.Param("ticketCode"))
	if err != nil {
		ctx.JSON(400, err.Error())
	} else if exist == false {
		ctx.JSON(200, "Not Exist")
	} else if exist == true {
		ctx.JSON(400, "Exist")
	}
}

type PostData struct {
	CodeSet []string `json:"CodeSet" binding:"required"`
}

func TicketCreateAPI(ctx *gin.Context) {
	var data PostData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	t := entry.TimeSelect()

	for _, code := range data.CodeSet {
		err := ticket.Create(code, t)
		if err != nil {
			ctx.JSON(400, err.Error())
		}
	}

	ctx.JSON(200, t)
}
