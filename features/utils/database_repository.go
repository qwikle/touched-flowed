package utils

type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
}

type Database interface {
	Query(query string, args ...interface{}) (Rows, error)
}
