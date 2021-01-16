package database

type Database interface {
	Collection(name string) (Collection, error)
}

type Collection interface {
}
