package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_Run(t *testing.T) {
	s := NewServer()
	router := s.registerRouter()

	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/ping")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode, "El código de estado debe ser 200 OK")
}
