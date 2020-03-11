package entities

/**
 * User - модель, представляющая пользователя
 *
 * Id - уникальный идентификатор пользователя
 * Username - уникальное имя пользователя
 * Created_At - время создания пользователя
 */
type User struct {
    Id         string `json:"id"`
    Username   string `json:"username"`
    Created_At string `json:"created_at"`
}

/**
 * NewUser - конструктор структуры User,
 * возвращающий экземпляр пользователя
 */
func NewUser(id, username, created_at string) *User {
    return &User{id, username, created_at}
}