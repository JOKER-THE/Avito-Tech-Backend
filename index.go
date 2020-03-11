package main

import (
    "fmt"
    "net/http"
)

/**
 * Главная страница
 *
 */
func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Avito!")
}

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

    /**
     * Старт сервера на 9000-порту
     *
     */
    http.ListenAndServe(":9000", nil)
}