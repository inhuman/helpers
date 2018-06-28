package http_test

import (
	"testing"
	"errors"
	ht "github.com/inhuman/helpers/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestCheckErrorHTTP(t *testing.T) {

	err := getError()

	w := httptest.NewRecorder()

	expected := `{"error":"test error"}`

	isError := ht.CheckErrorHTTP(err, w, 500)

	if w.Body.String() != expected {
		t.Errorf("writer returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}

	if status := w.Code; status != http.StatusInternalServerError {
		t.Errorf("writer returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	assert.Equal(t, true, isError)
}

func TestCheckErrorHTTPNoError(t *testing.T) {

	err := getNilError()

	w := httptest.NewRecorder()

	expected := ``

	isError := ht.CheckErrorHTTP(err, w, 500)

	if w.Body.String() != expected {
		t.Errorf("writer returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}

	if status := w.Code; status != http.StatusOK {
		t.Errorf("writer returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	assert.Equal(t, false, isError)
}

func getNilError() error {
	return nil
}

func getError() error {
	return errors.New("test error")
}
