package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteJSON(interface{}) error
}
