package triveutil

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

var (
	// ErrorLogger contains the standard logging template for errors
	ErrorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
	// SuccessLogger contains the standard logging template for successes
	SuccessLogger = log.New(os.Stderr, "SUCCESS ", log.Llongfile)
)

// StandardResponse contains the typical response structure of a Trive API
type StandardResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// ServerError generate the standard template for an internal server error
func ServerError(err error) (events.APIGatewayProxyResponse, error) {
	ErrorLogger.Println(err.Error())
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       StatusResponse(http.StatusInternalServerError),
	}, nil
}

// ClientError generates an AWS API Gateway formatted response for client-related errors
func ClientError(status int) (events.APIGatewayProxyResponse, error) {
	ErrorLogger.Println(fmt.Sprintf("There was a client-related error (%d)", status))
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       StatusResponse(status),
	}, nil
}

// StatusResponse generates the standard, Trive API response for a http status
func StatusResponse(status int) string {
	responseBody := &StandardResponse{
		Status:  status,
		Message: http.StatusText(status),
	}
	js, err := json.Marshal(responseBody)
	if err != nil {
		return http.StatusText(status)
	}
	return string(js)
}
