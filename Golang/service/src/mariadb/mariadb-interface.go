package mariadb

type DatabaseInterface interface {
	GetDbVersion() string
}
