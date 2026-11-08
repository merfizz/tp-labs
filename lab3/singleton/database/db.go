package database

import (
	"fmt"
	"sync"
	"time"
)

// dbConnection — неэкспортируемая структура (аналог private),
// хранящая состояние подключения.
type dbConnection struct {
	connectionString string
	connectedAt      time.Time
}

// глобальные переменные
var (
	// instance — единственный экземпляр
	instance *dbConnection
	// once — мьютекс, гарантирующий, что инициализация пройдет строго 1 раз
	once sync.Once
)

// GetInstance — глобальная точка доступа к БД.
func GetInstance() *dbConnection {
	once.Do(func() {
		fmt.Println("[db] Инициализация подключения к базе данных... (выполняется лишь однажды)")

		// Имитация задержки при подключении к тяжелой БД
		time.Sleep(500 * time.Millisecond)

		instance = &dbConnection{
			// Знакомая тебе схема базы
			connectionString: "Server=localhost;Database=AdventureWorksLT2022;User Id=sa;Password=secret;",
			connectedAt:      time.Now(),
		}
		fmt.Println("[db] Пул соединений успешно создан!")
	})

	return instance
}

// Query — публичный метод для работы с базой
func (db *dbConnection) Query(sql string) {
	fmt.Printf("[db] Выполнение запроса: '%s' к БД: %s\n", sql, db.connectionString)
}
