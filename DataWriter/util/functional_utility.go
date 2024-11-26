package util

type Consumer[T any] interface {
	Consume(T)
}
