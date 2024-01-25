package api

import (
	"strings"
	"testing"

	"bytes"

	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"encoding/json"
	"io"
	"net/http"

	"openapi/internal/infra/oapi_codegen/stockitem_api"
)

func TestPut(t *testing.T) {

	// When
	postRequestBody := &stockitem_api.PostStockItemJSONBody{
		Name: uuid.NewString(),
	}
	postRequestBodyJson, _ := json.Marshal(postRequestBody)
	client := http.Client{}
	postRequest, newReqErr := http.NewRequest(
		http.MethodPost,
		"http://localhost:3000/stock/items",
		bytes.NewBuffer(postRequestBodyJson))
	if newReqErr != nil {
		t.Fatal(newReqErr)
	}
	postRequest.Header.Set("Content-Type", "application/json")
	postResponse, reqErr := client.Do(postRequest)
	if reqErr != nil {
		t.Fatal(reqErr)
	}
	defer postResponse.Body.Close()
	resBodyByte, _ := io.ReadAll(postResponse.Body)
	postResponseBody := &stockitem_api.Created{}
	json.Unmarshal(resBodyByte, &postResponseBody)

	// Then
	if postResponse.StatusCode != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusCreated, postResponse.StatusCode)
	}

	if postResponseBody.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}

}

func TestPutValidation(t *testing.T) {

	// Generate a string of 101 characters
	longName := strings.Repeat("a", 101)

	requestBody := &stockitem_api.PostStockItemJSONBody{
		Name: longName,
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
