package main

import (
    "./entities"
    "database/sql"
    "encoding/json"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "net/http"
    "time"
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
        var data map[string]interface{}
        json.NewDecoder(r.Body).Decode(&data)
        user := data["username"]
        created := time.Now().Unix()
        
        result, err := database.Exec(`
            INSERT INTO users (username, created_at) VALUES (?, ?)
        `, user, created)
        if err != nil {
            panic(err.Error())
        }

        id, err := result.LastInsertId()
        if err != nil{
            panic(err)
        }

        json.NewEncoder(w).Encode(id)
    }
}

/**
 * Создать новый чат между пользователями
 *
 */
func addChatHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var data map[string]interface{}
        json.NewDecoder(r.Body).Decode(&data)
        name := data["name"]
        created := time.Now().Unix()
        users := data["users"].([]interface{})

        result, err := database.Exec(`
            INSERT INTO chats (name, created_at) VALUES (?, ?)
        `, name, created)
        if err != nil {
            panic(err.Error())
        }

        id, err := result.LastInsertId()
        if err != nil{
            panic(err)
        }

        for _, v := range users {
            database.Exec(`
                INSERT INTO chats_users (user_id, chat_id) VALUES (?, ?)
            `, v, id)
        }

        json.NewEncoder(w).Encode(id)
    }
}

/**
 * Получить список чатов конкретного пользователя
 *
 */
func getChatHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var chats []entities.Chat
        var data map[string]interface{}

        json.NewDecoder(r.Body).Decode(&data)

        user := data["user"]
        
        result, err := database.Query(`
            SELECT chats.id, chats.name, users.users, chats.created_at
            FROM chats_users
            INNER JOIN chats ON chats_users.chat_id = chats.id
            INNER JOIN messages ON chats.id = messages.chat
            INNER JOIN (
                SELECT 
                chats.id, GROUP_CONCAT(DISTINCT users.id ORDER BY users.id ASC SEPARATOR ", ") AS users 
                FROM chats 
                LEFT OUTER JOIN chats_users ON chats.id = chats_users.chat_id 
                LEFT OUTER JOIN users ON chats_users.user_id = users.id 
                GROUP BY chats.id
            ) AS users ON chats_users.chat_id = users.id
            WHERE user_id = ?
            GROUP BY chats_users.chat_id
            ORDER BY messages.created_at DESC
        `, user)
        if err != nil {
            panic(err.Error())
        }

        defer result.Close()

        for result.Next() {
            var chat entities.Chat
            err := result.Scan(&chat.Id, &chat.Name, &chat.Users, &chat.Created_At)
            if err != nil {
                panic(err.Error())
            }
            chats = append(chats, chat)
        }

        json.NewEncoder(w).Encode(chats)
    }
}

/**
 * Отправить сообщение в чат от лица пользователя
 *
 */
func addMessageHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var data map[string]interface{}
        json.NewDecoder(r.Body).Decode(&data)
        chat := data["chat"]
        author := data["author"]
        text := data["text"]
        created := time.Now().Unix()

        result, err := database.Exec(`
            INSERT INTO messages (chat, author, text, created_at) VALUES (?, ?, ?, ?)
        `, chat, author, text, created)
        if err != nil {
            panic(err.Error())
        }

        id, err := result.LastInsertId()
        if err != nil{
            panic(err)
        }

        json.NewEncoder(w).Encode(id)
    }
}

/**
 * Получить список сообщений в конкретном чате
 *
 */
func getMessageHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var messages []entities.Message
        var data map[string]interface{}

        json.NewDecoder(r.Body).Decode(&data)

        chat := data["chat"]
        
        result, err := database.Query(`
            SELECT id, chat, author, text, created_at
            FROM messages
            WHERE chat = ?
            ORDER BY created_at ASC
        `, chat)
        if err != nil {
            panic(err.Error())
        }

        defer result.Close()

        for result.Next() {
            var message entities.Message
            err := result.Scan(&message.Id, &message.Chat, &message.Author, &message.Text, &message.Created_At)
            if err != nil {
                panic(err.Error())
            }
            messages = append(messages, message)
        }

        json.NewEncoder(w).Encode(messages)
    }
}