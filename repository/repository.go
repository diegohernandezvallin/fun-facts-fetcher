package repository

type Repository interface {
	PutItem(funFactItem FunFactItem) error
}

type FunFactItem struct {
	LastTimePolled int64
	FunFact        string
}
