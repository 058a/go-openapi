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
	requestBody := &stockitem_api.PostStockItemJSONBody{
		Name: uuid.NewString(),
	}
	requestBodyJson, _ := json.Marshal(requestBody)
	client := http.Client{}
	request, newReqErr := http.NewRequest(
		http.MethodPost,
		"http://localhost:3000/stock/items",
		bytes.NewBuffer(requestBodyJson))
	if newReqErr != nil {
		t.Fatal(newReqErr)
	}
	request.Header.Set("Content-Type", "application/json")
	response, reqErr := client.Do(request)
	if reqErr != nil {
		t.Fatal(reqErr)
	}
	defer response.Body.Close()
	resBodyByte, _ := io.ReadAll(response.Body)
	actualResponse := &stockitem_api.Created{}
	json.Unmarshal(resBodyByte, &actualResponse)

	// Then
	if response.StatusCode != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusCreated, response.StatusCode)
	}

	if actualResponse.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}

}

func TestPostStockItemValidation(t *testing.T) {

	// When
	// Generate a string of 101 characters
	length101name := ""
	for i := 0; i < 101; i++ {
		length101name += "a"
	}

	requestBody := &stockitem_api.PostStockItemJSONBody{
		Name: length101name,
	}
	requestBodyJson, _ := json.Marshal(requestBody)
	client := http.Client{}
	request, newReqErr := http.NewRequest(
		http.MethodPost,
		"http://localhost:3000/stock/items",
		bytes.NewBuffer(requestBodyJson))
	if newReqErr != nil {
		t.Fatal(newReqErr)
	}
	request.Header.Set("Content-Type", "application/json")
	response, reqErr := client.Do(request)
	if reqErr != nil {
		t.Fatal(reqErr)
	}
	defer response.Body.Close()
	resBodyByte, _ := io.ReadAll(response.Body)
	actualResponse := &stockitem_api.Created{}
	json.Unmarshal(resBodyByte, &actualResponse)

	// Then
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("want %d, got %d", http.StatusBadRequest, response.StatusCode)
	}
}
