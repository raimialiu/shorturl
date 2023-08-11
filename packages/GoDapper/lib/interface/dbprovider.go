package GoMapper

import (
	"database/sql"
)

type DbProvider interface {
	Open() (status bool, provider DbProvider, err error)
	Close()
	DbObject() *sql.DB
	QueryClient() RowQuery
}
