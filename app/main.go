package main

import (
	"fmt"

	"github.com/shibe/notion_hozonn_api/app/external_service_api"
	"github.com/shibe/notion_hozonn_api/app/external_service_api/notion/notion_db"
)

func main() {
    // レスポンスを表示
    nl, _ := external_service_api.GetNihonshuList()


    auth := "secret_KpJbwwAQ0CMnakH11vVIOIP0bmaE0OVwFRCmgtTwy9Z"
    dbid := "0855121019dc4ba998814515fdad9bc9"
    c := notion_db.NewNotionDBClient(auth, dbid)

    for i, n := range nl {
        o := notion_db.ConvertToNotionObject(n)
        response, err := c.Save(o)
        if err != nil {
            fmt.Printf("%s\n", err)
            return
        } else if response != "200 OK" {
            fmt.Printf("%s\n", response)
            return
        }
        if i == 1 {
            return
        }
    }
}