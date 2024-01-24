package stockitem

import (
	"testing"

	"bytes"

	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"encoding/json"
	"io"
	"net/http"

	"openapi/internal/infra/oapi_codegen/stockitem_api"
)

func TestPostStockItem(t *testing.T) {

	// When
	request := &stockitem_api.PostStockItemJSONBody{
		Name: uuid.NewString(),
	}
	requestJson, _ := json.Marshal(request)
	client := http.Client{}
	req, newReqErr := http.NewRequest(
		http.MethodPost,
		"http://localhost:3000/stock/items",
		bytes.NewBuffer(requestJson))
	if newReqErr != nil {
		t.Fatal(newReqErr)
	}
	req.Header.Set("Content-Type", "application/json")
	res, reqErr := client.Do(req)
	if reqErr != nil {
		t.Fatal(reqErr)
	}
	defer res.Body.Close()
	resBodyByte, _ := io.ReadAll(res.Body)
	actualCreatedResponse := &stockitem_api.Created{}
	json.Unmarshal(resBodyByte, &actualCreatedResponse)

	// Then
	if res.StatusCode != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusCreated, res.StatusCode)
	}

	if actualCreatedResponse.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}

}
