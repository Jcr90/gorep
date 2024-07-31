/*

package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"gotest.tools/v3/assert"
)

func setupTestEcho() *echo.Echo {
	e := echo.New()
	repo := NewUserRepository(db) // Asegúrate de tener una base de datos de prueba configurada.
	handler := NewHandler(repo)
	e.POST("/users", handler.CreateUser)
	e.GET("/users/:id", handler.GetUser)
	e.PUT("/users/:id", handler.UpdateUser)
	e.DELETE("/users/:id", handler.DeleteUser)
	return e
}

func TestCreateUser(t *testing.T) {
	e := setupTestEcho()

	// Test case: Successful creation
	userJSON := `{"name":"John Doe","email":"john.doe@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NilError(t, e.Validate(user))
	e.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusCreated)
	assert.Assert(t, strings.Contains(rec.Body.String(), "John Doe"))

	// Test case: Missing name
	userJSON = `{"email":"john.doe@example.com"}`
	req = httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	assert.NilError(t, e.Validate(user))
	e.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusBadRequest)
}

func TestGetUser(t *testing.T) {
	e := setupTestEcho()

	// Create user first
	userJSON := `{"name":"John Doe","email":"john.doe@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	e.Validate(user)
	e.ServeHTTP(rec, req)

	// Get user
	req = httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NilError(t, e.Validate(user))
	e.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Assert(t, strings.Contains(rec.Body.String(), "John Doe"))

	// Test case: User not found
	req = httptest.NewRequest(http.MethodGet, "/users/999", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("999")

	assert.NilError(t, e.Validate(user))
	e.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusNotFound)
}

func TestUpdateUser(t *testing.T) {
	e := setupTestEcho()

	// Create user first
	userJSON := `{"name":"John Doe","email":"john.doe@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	e.Validate(user)
	e.ServeHTTP(rec, req)

	// Update user
	userJSON = `{"name":"Jane Doe","email":"jane.doe@example.com"}`
	req = httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NilError(t, e.Validate(user))
	e.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Assert(t, strings.Contains(rec.Body.String(), "Jane Doe"))

	// Test case: User not found
	req = httptest.NewRequest(http.MethodPut, "/users/999", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("999")

	assert.NilError(t, e.Validate(user))
	e.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusNotFound)
}

func TestDeleteUser(t *testing.T) {
	e := setupTestEcho()

	// Create user first
	userJSON := `{"name":"John Doe","email":"john.doe@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	e.Validate(user)
	e.ServeHTTP(rec, req)

	// Delete user
	req = httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NilError(t, e.Validate(user))
	e.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusNoContent)

	// Test case: User not found
	req = httptest.NewRequest(http.MethodDelete, "/users/999", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("999")

	assert.NilError(t, e.Validate(user))
	e.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusNotFound)
}


*/