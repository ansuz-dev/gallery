package routes

import (
  "encoding/json"
  "gallery/models"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)

func TestGetAccount(t *testing.T) {
  // get token first
  postData := strings.NewReader(`{"email":"aoshi@ansuzdev.com","password":"my_password"}`)
  writer := httptest.NewRecorder()
  request, _ := http.NewRequest("POST", "/v1/authentication", postData)
  ts.ServeHTTP(writer, request)
  if writer.Code != 200 {
    t.Errorf("Response code should be Ok, was: %d", writer.Code)
  }

  token := writer.Body.String()

  t.Run("Shoud get account for authorized user", func(t *testing.T) {
    writer := httptest.NewRecorder()
    request, _ := http.NewRequest("GET", "/v1/account", nil)
    request.Header.Set("Authorization", "Bearer "+token)
    ts.ServeHTTP(writer, request)
    if writer.Code != 200 {
      t.Errorf("Response code should be Ok, was: %d", writer.Code)
    }

    var account models.Account
    err := json.Unmarshal(writer.Body.Bytes(), &account)
    if err != nil {
      t.Errorf("Response should be an account object")
    }

    if account.Email != "aoshi@ansuzdev.com" {
      t.Errorf("Returned account should be aoshi@ansuzdev.com, was: %s", account.Email)
    }
  })

  t.Run("Shoud throw error if not authorized", func(t *testing.T) {
    writer := httptest.NewRecorder()
    request, _ := http.NewRequest("GET", "/v1/account", nil)
    ts.ServeHTTP(writer, request)
    if writer.Code != 401 {
      t.Errorf("Response code should be 401, was: %d", writer.Code)
    }
  })
}

func TestUpdateAccount(t *testing.T) {
  // get token first
  postData := strings.NewReader(`{"email":"aoshi@ansuzdev.com","password":"my_password"}`)
  writer := httptest.NewRecorder()
  request, _ := http.NewRequest("POST", "/v1/authentication", postData)
  ts.ServeHTTP(writer, request)
  if writer.Code != 200 {
    t.Errorf("Response code should be Ok, was: %d", writer.Code)
  }

  token := writer.Body.String()

  t.Run("Shoud update account for authorized user", func(t *testing.T) {
    t.Skip("Skipped after passed")
    postData := strings.NewReader(`{"email":"updated@ansuzdev.com"}`)
    writer := httptest.NewRecorder()
    request, _ := http.NewRequest("PUT", "/v1/account", postData)
    request.Header.Set("Authorization", "Bearer "+token)
    ts.ServeHTTP(writer, request)
    if writer.Code != 200 {
      t.Errorf("Response code should be Ok, was: %d", writer.Code)
    }
  })

  t.Run("Shoud throw error if not authorized", func(t *testing.T) {
    postData := strings.NewReader(`{"email":"updated@ansuzdev.com"}`)
    writer := httptest.NewRecorder()
    request, _ := http.NewRequest("PUT", "/v1/account", postData)
    ts.ServeHTTP(writer, request)
    if writer.Code != 401 {
      t.Errorf("Response code should be 401, was: %d", writer.Code)
    }
  })
}
