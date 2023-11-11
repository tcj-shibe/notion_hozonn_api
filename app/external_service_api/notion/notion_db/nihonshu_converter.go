package notion_db

import (
	"github.com/shibe/notion_hozonn_api/app/domain/model/nihonshu"
)

// ConvertToNotionObject
// NotionObjectって言いながらめちゃくちゃ複雑で
// どうやってタイプごとにまとめればいいのか測りかねてる
func ConvertToNotionObject (n nihonshu.Nihonshu) NotionObject {
	r := NotionObject{
		"名前": map[string]interface{}{
			"type": "title",
			"title": []map[string]interface{}{
				{
					"type": "text",
					"text": map[string]interface{}{
						"content": n.Name,
					},
				},
			},
		},
		"ID": map[string]interface{}{
			"type": "number",
			"number": n.ID,
		},
		"蔵元ID": map[string]interface{}{
			"type": "number",
			"number": n.BreweryID,
		},
	}
	return r
}