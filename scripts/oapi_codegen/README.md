~/go/bin/oapi-codegen -generate types,server,spec -package openapi ./../../api/hello.yaml > ./../../internal/infra/oapi_codegen/hello/hello.go && go mod tidy

~/go/bin/oapi-codegen -generate types,server,spec -package openapi ./../../api/stockitem.yaml > ./../../internal/infra/oapi_codegen/stockitem/stockitem.go && go mod tidy