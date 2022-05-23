package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

var app *fiber.App
var fiberLambda *fiberadapter.FiberLambda

func init() {
	log.Printf("Fiber cold start")
	app = fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	fiberLambda = fiberadapter.New(app)
}

func main() {
	if isRunningOnLambda() {
		lambda.Start(Handler)
	} else {
		if err := app.Listen(":3000"); err != nil {
			log.Fatal(err)
		}
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func isRunningOnLambda() bool {
	runtime := os.Getenv("AWS_LAMBDA_RUNTIME_API")

	return runtime != ""
}
