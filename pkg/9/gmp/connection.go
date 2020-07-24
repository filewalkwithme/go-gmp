package gmp

type Connection interface {
	Execute(command interface{}, response interface{}) error
	Close() error
}
