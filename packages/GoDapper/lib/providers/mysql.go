package providers

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	constants2 "shorturl/packages/GoDapper/constants"
	"shorturl/packages/GoDapper/lib"
	"shorturl/packages/GoDapper/lib/helpers"
	GoMapper "shorturl/packages/GoDapper/lib/interface"
	"shorturl/packages/GoDapper/utilities"
)

type MysqlDbProvider struct {
	DbConfig MysqlDbConfig
	State    string
	Object   *sql.DB
}

type MysqlDbConfig struct {
	User     string
	Password string
	DbName   string
	DbPort   int32
	DbHost   string
}

func New(config MysqlDbConfig) MysqlDbProvider {
	return MysqlDbProvider{DbConfig: config}
}

func FormatSqlConnectionString(config *MysqlDbConfig) string {
	if config.DbPort == 0 {
		config.DbPort = 3306
	}
	cfg := mysql.Config{
		User:   config.User,
		Passwd: config.Password,
		Net:    "tcp",
		Addr:   utilities.ConcatStrings(config.DbHost, ":", string(config.DbPort)),
		DBName: config.DbName,
	}

	return cfg.FormatDSN()
}

// INTERFACE_METHODS
func (m *MysqlDbProvider) Close() {
	if m.State != constants2.OPEN.String() {
		panic("db connection not opened, please open")
	}

	if m.Object == nil {
		panic("db object is null")
	}

	if err := m.Object.Close(); err != nil {
		panic(err)
	}
}

func (m *MysqlDbProvider) DbObject() *sql.DB {
	return m.Object
}

func (m *MysqlDbProvider) QueryClient() GoMapper.RowQuery {
	return lib.NewQuerClient(m.DbObject())
}

func (g *MysqlDbProvider) Open() (status bool, provider GoMapper.DbProvider, err error) {
	var connectionString string = FormatSqlConnectionString(&g.DbConfig)
	db := helpers.ConnectToDb(constants2.MYSQL, connectionString)
	g.State = constants2.OPEN.String()
	g.Object = db
	//provider := *(*GoMapper.DbProvider)g

	return true, g, nil
}
