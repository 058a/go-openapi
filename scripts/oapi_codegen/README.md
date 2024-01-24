~/go/bin/oapi-codegen -generate types,server,spec -package hello_api ./../../api/hello.yaml > ./../../internal/infra/oapi_codegen/hello_api/hello_api.go && go mod tidy

~/go/bin/oapi-codegen -generate types,server,spec -package stockitem_api ./../../api/stockitem.yaml > ./../../internal/infra/oapi_codegen/stockitem_api/stockitem_api.go && go mod tidy