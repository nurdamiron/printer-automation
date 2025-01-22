package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/nurdamiron/printer-automation/internal/handler"
    "github.com/nurdamiron/printer-automation/internal/repository"
    "github.com/nurdamiron/printer-automation/internal/service"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    // Загрузим переменные окружения, если нужно
    _ = godotenv.Load(".env")

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Формируем строку подключения
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPass, dbName,
    )

    // Инициализируем GORM
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Не удалось подключиться к базе данных: %v", err)
    }

    // Инициализация репозиториев
    userRepo := repository.NewUserRepository(db)
    jobRepo := repository.NewPrintJobRepository(db)

    // Инициализация сервисов
    userService := service.NewUserService(userRepo)
    jobService := service.NewPrintJobService(jobRepo)

    // Инициализация хендлеров
    userHandler := handler.NewUserHandler(userService)
    jobHandler := handler.NewPrintJobHandler(jobService)

    // Создаём Gin router
    r := gin.Default()

    // Пример middleware для CORS, логирования и т.п.
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    // Маршруты для User
    r.POST("/api/v1/users", userHandler.CreateUser)
    r.GET("/api/v1/users/:id", userHandler.GetUser)

    // Маршруты для Print Jobs
    r.POST("/api/v1/print-jobs", jobHandler.CreateJob)
    r.GET("/api/v1/print-jobs/:id", jobHandler.GetJob)

    // Простой ping endpoint
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    // Запуск сервера
    fmt.Println("Сервер запущен на http://localhost:8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Ошибка при запуске сервера: %v", err)
    }
}
