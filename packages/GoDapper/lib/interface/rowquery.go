package GoMapper

type RowQuery interface {
	QueryFirst(query string) (data interface{}, err error)
	Query(query string)
	ExecuteNonQuery(query string)
}

// WAKE UP MAN (THE ERA (DEVELOPMENT AND SERIOUS WORK) HAS BEGUN)
// SERIOUS WORK AND MORE WORK....SEE YOU AT WORK
