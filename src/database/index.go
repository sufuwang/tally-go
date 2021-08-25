package database

import (
	mysql "tally-go/src/database/mysql"
)

func LinkDataBase() {
	mysql.Link()
}
