package entities

/**
 * Chat - модель, представляющая чат
 *
 * Id - уникальный идентификатор чата
 * Name - уникальное имя чата
 * Users - список пользователей в чате
 * Created_At - время создания
 */
type Chat struct {
    Id         string `json:"id"`
    Name       string `json:"name"`
    Users      string `json:"users"`
    Created_At string `json:"created_at"`
}

/**
 * NewChat - конструктор структуры Chat,
 * возвращающий экземпляр чата
 */
func NewChat(id, name, users, created_at string) *Chat {
    return &Chat{id, name, users, created_at}
}