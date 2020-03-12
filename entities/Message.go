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
    Id         int32  `json:"id"`
    Chat       int32  `json:"chat"`
    Author     int32  `json:"author"`
    Text       string `json:"text"`
    Created_At string `json:"created_at"`
}

/**
 * NewMessage - конструктор структуры Message,
 * возвращающий экземпляр сообщения
 */
func NewMessage(id int32, chat int32, author int32, text, created_at string) *Message {
    return &Message{id, chat, author, text, created_at}
}