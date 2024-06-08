package api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/zacksfF/Hotel-Reservation-Backend/types"
)

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	// userHandler := NewUserHandler(tdb.UserStore)
	// app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "zakaria.saiff@mm.com",
		Password:  "password12345",
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	var user types.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		t.Fatal(err)
	}
	if user.ID.IsZero() {
		t.Error("Expected user.ID to be not empty")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Error("Expected user.EncryptedPassword to be not returned")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("Expected %s, got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("Expected %s, got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("Expected %s, got %s", params.Email, user.Email)
	}
}