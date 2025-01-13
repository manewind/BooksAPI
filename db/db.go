package db

import "context"


type Databse[T any] interface{
	Connect() error
	Close() error
	Query(ctx context.Context,query string, args... any)([]T, error)
	Exec(ctx context.Context,query string, args... any)(int64, error)
}
