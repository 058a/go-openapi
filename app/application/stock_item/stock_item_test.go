package stock_item

import (
	"testing"
	"time"

	"bytes"
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/google/uuid"

	"encoding/json"
	"io"
	"net/http"

	repository_models "openapi/repository/models"
)

func TestPostStockItem(t *testing.T) {
	// Given
	executedAt := time.Now().UTC()

	// When
	request := new(PostStockItemJSONBody)
	request.Name = uuid.NewString()
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
	actualCreatedResponse := &StockItem{}
	json.Unmarshal(resBodyByte, &actualCreatedResponse)

	// Then
	if res.StatusCode != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusCreated, res.StatusCode)
	}

	if actualCreatedResponse.Id == uuid.Nil {
		t.Errorf("expected not empty, actual empty")
	}

	// Database Test
	dbDriver := "postgres"
	dsn := "host=openapi-db port=5432 user=user password=password dbname=openapi sslmode=disable"

	db, openErr := sql.Open(dbDriver, dsn)
	if openErr != nil {
		t.Fatal(openErr)
	}
	defer db.Close()

	existed, findErr := repository_models.FindStockItem(context.Background(), db, actualCreatedResponse.Id.String())
	if findErr != nil {
		t.Errorf(findErr.Error())
	}

	if existed.Name != request.Name {
		t.Errorf("expected %s, actual %s", request.Name, existed.Name)
	}

	if existed.CreatedAt.Before(executedAt) {
		t.Errorf("CreatedAt %s is before executedAt %s", existed.CreatedAt, executedAt)
	}

	if existed.UpdatedAt.Before(executedAt) {
		t.Errorf("UpdatedAt %s is before executedAt %s ", existed.UpdatedAt, executedAt)
	}
}

func TestPutStockItem(t *testing.T) {
	// SetUp
	client := http.Client{}

	// Given
	postRequestBody := PostStockItemJSONBody{
		Name: uuid.NewString(),
	}
	postRequestBodyJson, _ := json.Marshal(postRequestBody)
	postRequest, newReqErr := http.NewRequest(
		http.MethodPost,
		"http://localhost:3000/stock/items",
		bytes.NewBuffer(postRequestBodyJson))
	if newReqErr != nil {
		t.Fatal(newReqErr)
	}
	postRequest.Header.Set("Content-Type", "application/json")
	res, reqErr := client.Do(postRequest)
	if reqErr != nil {
		t.Fatal(reqErr)
	}
	defer res.Body.Close()
	resBodyByte, _ := io.ReadAll(res.Body)
	postResponse := &StockItem{}
	json.Unmarshal(resBodyByte, &postResponse)

	// When
	putRequestBody := PutStockItemJSONBody{
		Name: uuid.NewString(),
	}
	putRequestJson, _ := json.Marshal(putRequestBody)
	putRequest, newReqErr := http.NewRequest(
		http.MethodPut,
		"http://localhost:3000/stock/items/{"+postResponse.Id.String()+"}",
		bytes.NewBuffer(putRequestJson))
	if newReqErr != nil {
		t.Fatal(newReqErr)
	}
	putRequest.Header.Set("Content-Type", "application/json")
	putResponse, reqErr := client.Do(putRequest)
	if reqErr != nil {
		t.Fatal(reqErr)
	}
	defer res.Body.Close()

	// Then
	if putResponse.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, res.StatusCode)
	}

	// Database Test
	dbDriver := "postgres"
	dsn := "host=openapi-db port=5432 user=user password=password dbname=openapi sslmode=disable"

	db, openErr := sql.Open(dbDriver, dsn)
	if openErr != nil {
		t.Fatal(openErr)
	}
	defer db.Close()

	existed, findErr := repository_models.FindStockItem(context.Background(), db, postResponse.Id.String())
	if findErr != nil {
		t.Errorf(findErr.Error())
	}

	if existed.Name != postRequestBody.Name {
		t.Errorf("expected %s, actual %s", postRequestBody.Name, existed.Name)
	}

	if existed.CreatedAt.Equal(postResponse.CreatedAt) {
		t.Errorf("CreatedAt %s is postResponse.CreatedAt %s", existed.CreatedAt, postResponse.CreatedAt)
	}

	if existed.UpdatedAt.Before(postResponse.UpdatedAt) {
		t.Errorf("UpdatedAt %s is postResponse.UpdatedAt %s ", existed.UpdatedAt, postResponse.UpdatedAt)
	}
}
