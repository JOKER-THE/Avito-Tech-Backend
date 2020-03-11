package main

import (
    "fmt"
    "net/http"
)

/**
 * Основной метод Avito-Tech-Backend
 * Метод инициализации программы
 *
 */
func main() {
    /**
     * Роутинг URL
     *
     */
    http.HandleFunc("/", indexHandler)
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
 * Главная страница
 *
 */
func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "{\"message\": \"Hello, Avito!\"}")
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