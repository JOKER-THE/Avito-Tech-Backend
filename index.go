package main

import (
    "./entities"
    "database/sql"
    "encoding/json"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "net/http"
)

var database *sql.DB

/**
 * Основной метод Avito-Tech-Backend
 * Метод инициализации программы
 *
 */
func main() {

    /**
     * Подключение к базе данных
     *
     */
    db, err := sql.Open("mysql", "root:@/avito-tech")
     
    if err != nil {
        fmt.Println(err)
    }
    database = db
    defer db.Close()

    /**
     * Роутинг URL
     *
     */
    http.HandleFunc("/users/add", addUserHandler)
    http.HandleFunc("/chats/add", addChatHandler)
    http.HandleFunc("/chats/get", getChatHandler)
    http.HandleFunc("/messages/add", addMessageHandler)
    http.HandleFunc("/messages/get", getMessageHandler)

    /**
     * Старт сервера на 9000-порту
     *
     */
    http.ListenAndServe(":9000", nil)
}

/**
 * Добавить нового пользователя
 *
 */
func addUserHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        
    }
}

/**
 * Создать новый чат между пользователями
 *
 */
func addChatHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        
    }
}

/**
 * Получить список чатов конкретного пользователя
 *
 */
func getChatHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        
    }
}

/**
 * Отправить сообщение в чат от лица пользователя
 *
 */
func addMessageHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        
    }
}

/**
 * Получить список сообщений в конкретном чате
 *
 */
func getMessageHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        
    }
}