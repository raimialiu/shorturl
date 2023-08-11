package constants

type Dialect string

const (
	undefined Dialect = ""
	MYSQL             = "mysql"
	POSTGRESS         = "postgress"
	SQLSERVER         = "sqlserver"
	SQLLITE           = "sqllite"
)

// 5W, GLG, LNX, CK,400 wasnet -> 10.45 -> 5W (IN PROGRESS)

func (d Dialect) String() string {
	enums := map[string]string{
		"mysql":     "mysql",
		"postgress": "postgress",
		"sqlserver": "sqlserver",
		"sqllite":   "sqllite",
	}

	return enums[string(d)]
}
