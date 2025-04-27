package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

// runMigrations симулирует выполнение миграций.
// В реальном проекте можно использовать специализированные библиотеки.
func runMigrations(db *sql.DB) {
	log.Println("Running automatic migrations...")
	// Здесь можно выполнить SQL-скрипты или вызвать миграционную утилиту.
	log.Println("Migrations applied successfully.")
}

func main() {
	// Чтение конфигурации из переменных окружения
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Устанавливаем соединение с базой данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	defer db.Close()

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		log.Fatalf("Unable to ping DB: %v", err)
	}

	// Выполняем миграции
	runMigrations(db)

	// Настраиваем сервер для отдачи статики из volume
	fs := http.FileServer(http.Dir("/app/static"))
	http.Handle("/", fs)

	srv := &http.Server{
		Addr: ":" + port,
	}

	// Запуск сервера в горутине
	go func() {
		log.Printf("Starting server on port %s...", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Организуем graceful shutdown при получении сигнала завершения (Ctrl+C, SIGTERM)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	log.Println("Shutting down server...")

	// Создаем контекст с таймаутом для завершения текущих операций
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped.")
}
