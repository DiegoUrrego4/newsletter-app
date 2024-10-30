package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_Run(t *testing.T) {
	s := NewServer()
	router := s.registerRouter()

	// Creamos un servidor de prueba
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Realizamos una solicitud GET al endpoint /ping
	res, err := http.Get(ts.URL + "/ping")
	assert.NoError(t, err, "No se esperaba error al hacer la solicitud GET")

	// Verificamos que el código de estado sea 200 OK
	assert.Equal(t, http.StatusOK, res.StatusCode, "El código de estado debe ser 200 OK")

	// Leemos el cuerpo de la respuesta
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	assert.NoError(t, err, "No se esperaba error al leer el cuerpo de la respuesta")

	// Verificamos que el contenido del cuerpo sea 'pong'
	assert.Equal(t, "pong", string(body), "El cuerpo de la respuesta debe ser 'pong'")
}
