package entities

/**
 * Message - модель, представляющая сообщение в чате
 *
 * Id - уникальный идентификатор сообщения
 * Chat - ссылка на идентификатор чата, в который было отправлено сообщение
 * Author - ссылка на идентификатор отправителя сообщения
 * Text - текст отправленного сообщения
 * Created_At - дата создания
 */
type Message struct {
    Id         string
    Chat       string
    Author     string
    Text       string
    Created_At string
}

/**
 * NewMessage - конструктор структуры Message,
 * возвращающий экземпляр сообщения
 */
func NewMessage(id, chat, author, text, created_at string) *Message {
    return &Message{id, chat, author, text, created_at}
}