package database

import "challenge-serasa/api/mainframe"

type Database interface {
	Collection(name string) (Collection, error)
}

type Collection interface {
	SaveDocuments(negativations []mainframe.Negativation) error
	GetDocuments(value interface{}, field string) (*[]mainframe.Negativation, error)
}
