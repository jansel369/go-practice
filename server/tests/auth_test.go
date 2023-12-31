package test

import (
	"io"
	model "server/models"
	"testing"
)

func TestAuth(t *testing.T) {
	client := NewClient()

	// registerUser := map[string]any {
	// 	"name": "Test Andrew",
	// 	"password": "password",
	// 	"email": "andrew@test.com",
	// 	"age": 32,
	// }

	registerInput := model.UserRegisterInput{
		Name:     "Andrew",
		Password: "password",
		Email:    "andrew@test.com",
		Age:      32,
	}

	regRes, regErr := client.POST("/auth/register").Body().AsJSON(registerInput).Send()

	if regRes.IsError() {
		body, readErr := io.ReadAll(regRes.RawBody())
		if readErr != nil {
			t.Errorf("Error reading body %s", readErr)
		}
		t.Errorf("body %v: %v", regRes.StatusCode(), string(body))
	}

	// var body map[any]any

	// readErr := json.NewDecoder(regRes.RawResponse.Body).Decode(&body)

	if regErr != nil {
		t.Errorf("Post request error: %v", regErr)
	}

	t.Errorf("\n response %v", regRes.IsError())
}
