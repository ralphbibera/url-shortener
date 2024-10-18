package logger

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/uuid"
)

type person struct {
	Id       *string   `json:"Id,omitempty"`
	Name     string    `json:"Name,omitempty"`
	Accounts []account `json:"Accounts,omitempty"`
}

type account struct {
	Number *int `json:"Number,omitempty"`
}

func TestStructuredLog(t *testing.T) {
	person := person{
		Id:   aws.String(uuid.NewString()),
		Name: "Ralph",
		Accounts: []account{
			{
				Number: aws.Int(1),
			},
		},
	}
	Log(person, "Input:")
}
