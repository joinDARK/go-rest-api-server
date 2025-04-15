package repositories

import "database/sql"

// Этот интерфейс создан против циклических зависимостей
type storeInterface interface {
    DB() *sql.DB // Метод для получения соединения с базой данных
}