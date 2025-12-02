package interfaces

type Storage interface {
	Save(string) error
	Load(string) error
	Set(string, interface{})
	Get(string) (interface{}, bool)
	Delete(string)
}
