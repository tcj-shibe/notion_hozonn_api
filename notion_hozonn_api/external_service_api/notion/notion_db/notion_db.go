package notion_db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type NOTION_COLUMN_TYPE string

const (
	NOTION_COLUMN_TYPE_TEXT   NOTION_COLUMN_TYPE = "text"
	NOTION_COLUMN_TYPE_NUMBER NOTION_COLUMN_TYPE = "number"
)

const (
	saveRecordEndPoint    = "https://api.notion.com/v1/pages"
	saveRecordContentType = "application/json"
	notion_version        = "2022-06-28"
)

type NotionColumn struct {
	IsTitle    bool
	ColumnName string
	ColumnType NOTION_COLUMN_TYPE
	Value      interface{}
}

type NotionObject map[string]interface{}

type implNotionDBClient struct {
	Authorization string
	DatabaseID    string
}

type NotionDBClient interface {
	Save(o NotionObject)(string, error)
}

func NewNotionDBClient(authorization string, databaseID string) NotionDBClient {
	return implNotionDBClient{
		Authorization: fmt.Sprintf("Bearer %s", authorization),
		DatabaseID:    databaseID,
	}
}

func (c implNotionDBClient) Save(o NotionObject)(string, error) {
	// 送信するJSONデータを作成
	requestData := map[string]interface{}{
		"parent": map[string]interface{}{
			"type":        "database_id",
			"database_id": c.DatabaseID,
		},
		"properties": o,
	}

	// JSONデータをバイト配列に変換
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	// リクエストを作成
	req, err := http.NewRequest("POST", saveRecordEndPoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err 
	}

	// リクエストヘッダーにContent-Typeを設定
	req.Header.Set("Authorization", c.Authorization)
	req.Header.Set("Notion-Version", notion_version)
	req.Header.Set("Content-Type", saveRecordContentType)

	// HTTPクライアントを作成
	client := &http.Client{}

	// リクエストを送信
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	s := resp.Status
	defer resp.Body.Close()
	return s, nil
}
