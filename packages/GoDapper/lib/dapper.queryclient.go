package lib

import "database/sql"

type QueryClient struct {
	dbObject *sql.DB
}

func NewQuerClient(db *sql.DB) QueryClient {
	return QueryClient{dbObject: db}
}

func (q QueryClient) QueryFirst(query string) (data interface{}, err error) {

}

func (q QueryClient) Query(query string) {

}
func (q QueryClient) ExecuteNonQuery(query string) {

}
