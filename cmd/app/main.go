package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ThembinkosiThemba/dynamodb-go-crud/config"
	"github.com/ThembinkosiThemba/dynamodb-go-crud/internal/repository/adapter"
	"github.com/ThembinkosiThemba/dynamodb-go-crud/internal/repository/instance"
	"github.com/ThembinkosiThemba/dynamodb-go-crud/internal/routes"
	"github.com/ThembinkosiThemba/dynamodb-go-crud/internal/rules"
	RulesProduct "github.com/ThembinkosiThemba/dynamodb-go-crud/internal/rules/product"
	"github.com/ThembinkosiThemba/dynamodb-go-crud/utils/logger"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	configs := config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)

	logger.INFO("waiting for the service to start...", nil)

	errors := Migrate(connection)
	if len(errors) >= 0 {
		for _, err := range errors {
			logger.PANIC("error on migration:...", err)
		}
	}

	logger.PANIC("", CheckTables(connection))

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	logger.INFO("service is running on port", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}

func Migrate(connection *dynamodb.DynamoDB) []error {
	var errors []error
	callMigrateAndAppendError(&errors, connection, &RulesProduct.Rules{})
	return errors
}

func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface) {
	err := rule.Migrate(connection)
	if err != nil {
		*errors = append(*errors, err)
	}
}

func CheckTables(connection *dynamodb.DynamoDB) error{
	res, err := connection.ListTables(&dynamodb.ListTablesInput{})
	if res != nil {
		if len(res.TableNames) == 0 {
			logger.INFO("tables not found:", nil)
		}
		for _, tableName := range res.TableNames{
			logger.INFO("table found:", *tableName)
		}
	}
	return err
}
