package external_service_api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/shibe/notion_hozonn_api/domain/model/nihonshu"
)

func GetNihonshuList() ([]nihonshu.Nihonshu, error) {
	// APIエンドポイント
	url := "https://muro.sakenowa.com/sakenowa-data/api/brands"

	// HTTP GETリクエストを作成
	resp, err := http.Get(url)
	if err != nil {
		return []nihonshu.Nihonshu{}, errors.Wrap(err, "request error.")
	}
	defer resp.Body.Close()

	// レスポンスボディを読み取り
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []nihonshu.Nihonshu{}, errors.Wrap(err, "failed to read response body.")
	}
	var data NihonshuResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []nihonshu.Nihonshu{}, errors.Wrap(err, "failed to Unmarshal response body.")
	}
	// FIXME need to throw
	return data.NihonshuList, nil
}

type NihonshuResponse struct {
	Copyright    string              `json:"copyright"`
	NihonshuList []nihonshu.Nihonshu `json:"brands"`
}
