package main

import (
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
)

var (
    ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
    ErrorFailedToFetchRecord     = "failed to fetch record"
    ErrorInvalidUserData         = "invalid user data"
    ErrorInvalidEmail            = "invalid email"
    ErrorCouldNotMarshalItem     = "could not marshal item"
    ErrorCouldNotDeleteItem      = "could not delete item"
    ErrorCouldNotDynamoPutItem   = "could not dynamo put item error"
    ErrorUserAlreadyExists       = "user.User already exists"
    ErrorUserDoesNotExists       = "user.User does not exist"
)

type MyEvent struct {
    Name string `json:"What is your name?"`
    Age int     `json:"How old are you?"`
}

type MyResponse struct {
    Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
    return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
    lambda.Start(HandleLambdaEvent)
}
