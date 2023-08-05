package packages

import (
	sql "database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"shorturl/packages/constants"
	"shorturl/packages/utilities"
)

type GoDapper struct {
	PingDB          bool
	AutoCloseDB     bool
	Config          *GoDapperConfig
	object          *sql.DB
	ConnectionState constants.SqlState
}

type GoDapperConfig struct {
	Dialect     constants.Dialect
	User        string
	Password    string
	DbName      string
	DbPort      int32
	DbHost      string
	Address     string
	NetworkType *string
}

func NewGoDapper(config *GoDapperConfig) GoDapper {
	return GoDapper{
		Config:      config,
		PingDB:      true,
		AutoCloseDB: false,
	}
}

func (g *GoDapper) GetSqlDialect() constants.Dialect {
	return g.Config.Dialect
}

func (g *GoDapper) FormatSqlConnectionString() string {
	if g.Config.Dialect == constants.MYSQL {
		cfg := mysql.Config{
			User:   g.Config.User,
			Passwd: g.Config.Password,
			Net:    "tcp",
			Addr:   utilities.ConcatStrings(g.Config.DbHost, ":", string(g.Config.DbPort)),
			DBName: g.Config.DbName,
		}

		return cfg.FormatDSN()
	}

	panic("string")
}

func _connectToDb(dialect string, connectionString string) (db *sql.DB) {
	db, err := sql.Open(dialect, connectionString)
	if err != nil {
		fmt.Println(db)
		return db
	}
	panic(err)
}

func (g GoDapper) Close() {
	if g.object == nil {
		panic("db not in opened state, error.... ")
	}

	if g.ConnectionState != constants.OPEN {
		panic("!oops, it seems the db connection is not opened...., ")
	}

	g.object.Close()
}

func (g *GoDapper) Open() (db *sql.DB) {
	var connectionString string = g.FormatSqlConnectionString()
	db = _connectToDb(g.GetSqlDialect().String(), connectionString)
	canCloseDb := g.AutoCloseDB == true

	if canCloseDb {
		g.ConnectionState = constants.OPEN
		g.object = db
		defer g.Close()
	}

	return db
}

type GoDbConnection interface {
	Open(dbConfig interface{})
	Close(instance *GoDapper)
	Ping()
	TryPing()
}
