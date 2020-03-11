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
    Id         string
    Name       string
    Users      string
    Created_At string
}

/**
 * NewChat - конструктор структуры Chat,
 * возвращающий экземпляр чата
 */
func NewChat(id, name, users, created_at string) *Chat {
    return &Chat{id, name, users, created_at}
}