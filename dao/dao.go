package dao

import "database/sql"

type userDao struct {
	db *sql.DB
}

type UserDaO interface {
}

func NewUserDao(dbuser, dbpassword, dbname, host string, port int) UserDaO {
	// Connect using the credientails
	return &userDao{db: nil}
}
