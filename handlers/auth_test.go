package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"dealls-dating-apps/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestSignUpHandler(t *testing.T) {
	// mock mongodb
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer func() {
		if mt.Client != nil {
			if err := mt.Client.Disconnect(context.Background()); err != nil {
				t.Fatalf("failed to disconnect mock client: %v", err)
			}
		}
	}()

	mt.Run("success signup", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		Initialize(mt.Client)

		user := models.User{
			Email:    "test@example.com",
			Password: "password",
		}
		jsonUser, _ := json.Marshal(user)

		req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonUser))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(SignUpHandler)

		// calling handler
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}

		var response map[string]string
		err = json.NewDecoder(rr.Body).Decode(&response)
		if err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		expected := map[string]string{"message": "user created"}
		if response["message"] != expected["message"] {
			t.Errorf("handler returned unexpected body: got %v want %v", response["message"], expected["message"])
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "dealls.users", mtest.FirstBatch, bson.D{
			{Key: "email", Value: user.Email},
			{Key: "password", Value: user.Password},
		}))

		var insertedUser models.User
		err = userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&insertedUser)
		if err != nil {
			t.Fatalf("user not inserted into database: %v", err)
		}
	})
}
