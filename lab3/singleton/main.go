package main

import (
	"fmt"
	"sync"
)

// Database — структура, экземпляр которой должен быть один
type Database struct {
	Connection string
}

var (
	instance *Database
	once     sync.Once
)

// GetInstance возвращает единственный экземпляр БД
func GetInstance() *Database {
	once.Do(func() {
		fmt.Println("Инициализация подключения к БД... (выполнится лишь 1 раз)")
		instance = &Database{Connection: "Server=localhost;Database=AdventureWorksLT2022;User Id=sa;Password=secret;"}
	})
	return instance
}

func main() {
	db1 := GetInstance()
	db2 := GetInstance()

	if db1 == db2 {
		fmt.Println("Успех: Обе переменные указывают на один и тот же адрес в памяти!")
	}
}
