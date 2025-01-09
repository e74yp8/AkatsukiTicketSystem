package googlesheet

//source: https://developers.google.com/sheets/api/quickstart/go?hl=zh-tw

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
	"strconv"
)

// 指定要写入的 Spreadsheet ID 和 Sheet 名称
var spreadsheetId = "114qcgtr7YergFkRYKOotj2g3-8OSDY3ZY3YoDD-ZqZM"
var srv *sheets.Service

func init() {
	b, err := ioutil.ReadFile("./credentials.json")
	if err != nil {
		log.Fatalf("无法读取密钥文件：%v", err)
	}

	// 从密钥文件创建一个配置
	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("无法创建配置：%v", err)
	}

	// 使用配置创建一个客户端
	client := config.Client(context.Background())

	// 创建 Sheets 服务
	srv, err = sheets.New(client)
	if err != nil {
		log.Fatalf("无法创建 Sheets 服务：%v", err)
	}
}

func GetLimitId() (idLimit int) {
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, "Sheet1!B1").Do()
	if err != nil {
		panic("Unable to retrieve data from sheet: %v" + err.Error())
	}

	if len(resp.Values) == 0 {
		panic("No data found.")
	} else {
		idLimit, err = strconv.Atoi(resp.Values[0][0].(string))
		if err != nil {
			panic("Unable to prase data: %v" + err.Error())
		}
		return idLimit
	}

	return 0
}
