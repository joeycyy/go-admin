package store

type UserStore interface {
	Create() error
	Update() error
	Delete() error
	DeleteCollection() error
	Get() error
	List() error
}
