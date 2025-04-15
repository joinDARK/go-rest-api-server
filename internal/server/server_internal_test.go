package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go-rest-api/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestServer_HandlerHello(t *testing.T) {
	// Создаем новый сервер
	srv := New(config.New())
	srv.configureRouter() // Регистрируем маршруты
	
	// Создаем тестовый запрос и рекордер
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	assert.NoError(t, err, "Failed to create request") // Проверяем ошибку создания запроса
	
	// Act: Выполняем запрос через роутер сервера
	srv.router.ServeHTTP(rec, req)
	
	// Assert: Проверяем результаты
	assert.Equal(t, http.StatusOK, rec.Code, "Expected status OK")
    assert.Equal(t, "Hello, World!", rec.Body.String(), "Response body mismatch")
}