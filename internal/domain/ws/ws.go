package ws

type Ws interface {
	WriteJSON(v interface{}) error
	Close() error
}
