package routes

import (
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)

func TestAuthentication(t *testing.T) {
  t.Run("Should authenticate with valid account", func(t *testing.T) {
    json := strings.NewReader(`{"email":"aoshi@ansuzdev.com","password":"my_password"}`)
    writer := httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/authentication", json)
    ts.ServeHTTP(writer, request)
    if writer.Code != 200 {
      t.Errorf("Response code should be Ok, was: %d", writer.Code)
    }
  })

  t.Run("Should throw error with invalid email", func(t *testing.T) {
    json := strings.NewReader(`{"email":"invalid@ansuzdev.com","password":"my_password"}`)
    writer := httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/authentication", json)
    ts.ServeHTTP(writer, request)
    if writer.Code != 401 {
      t.Errorf("Response code should be 401, was: %d", writer.Code)
    }
  })

  t.Run("Should throw error with invalid password", func(t *testing.T) {
    json := strings.NewReader(`{"email":"aoshi@ansuzdev.com","password":"invalid_password"}`)
    writer := httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/authentication", json)
    ts.ServeHTTP(writer, request)
    if writer.Code != 401 {
      t.Errorf("Response code should be 401, was: %d", writer.Code)
    }
  })
}

func TestRegistration(t *testing.T) {
  t.Run("Should register with email and password", func(t *testing.T) {
    t.Skip("Skipped after passed")
    json := strings.NewReader(`{"email":"aoshi2@ansuzdev.com","password":"my_password"}`)
    writer := httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/registration", json)
    ts.ServeHTTP(writer, request)
    if writer.Code != 200 {
      t.Errorf("Response code should be Ok, was: %d", writer.Code)
    }
  })

  t.Run("Should throw error if email already exist", func(t *testing.T) {
    json := strings.NewReader(`{"email":"aoshi@ansuzdev.com","password":"my_password"}`)
    writer := httptest.NewRecorder()
    request, _ := http.NewRequest("POST", "/v1/registration", json)
    ts.ServeHTTP(writer, request)
    if writer.Code != 400 {
      t.Errorf("Response code should be 400, was: %d", writer.Code)
    }
  })
}
