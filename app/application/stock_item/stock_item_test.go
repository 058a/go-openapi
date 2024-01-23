package stock_item

import (
	"bytes"
	"testing"

	"github.com/google/uuid"

	"encoding/json"
	"io"
	"net/http"
)

// http://localhost:3000/ にGETでアクセスし、戻り値を検証する
func TestStockItem(t *testing.T) {
	request := new(PostStockItemJSONBody)
	request.Name = uuid.NewString()
	requestJson, _ := json.Marshal(request)
	res, err := http.Post("http://localhost:3000/stock/items",
		"application/json",
		bytes.NewBuffer(requestJson))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusCreated, res.StatusCode)
	}

	resBodyByte, _ := io.ReadAll(res.Body)
	var actual = &CreatedResponse{}
	json.Unmarshal(resBodyByte, &actual)

	if actual.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}
}
