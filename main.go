package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ricardobaumann/go-reloading-proxy/loader"
)

var lazyLoader loader.LazyLoader

func init() {
	lazyLoader = loader.LazyLoaderImpl{
		UrlLoader: loader.UrlLoaderImpl{
			BasePath: "test",
		},
		Cache: loader.DummyMapCache{
			Repo: make(map[string]string),
		},
	}
}

//Handler handles incoming api gateway requests
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := strings.TrimPrefix(request.Path, "/reports/")
	fmt.Printf("ID %v", id)
	fmt.Println("Requested id : ", id)

	value := lazyLoader.LazyLoad(id)

	return events.APIGatewayProxyResponse{Body: value, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
