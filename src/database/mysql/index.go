package mysql

// Docker Mysql 镜像启动命令
// docker run -d -p 3306:3306
// 		-e MYSQL_USER="root"
// 		-e MYSQL_PASSWORD="Wo123456"
// 		-e MYSQL_ROOT_PASSWORD="Wo123456"
// 		--name tally-mysql
// 		mysql/mysql-server --character-set-server=utf8 --collation-server=utf8_general_ci

import (
	"database/sql"
	"fmt"
	"time"

	tool "tally-go/src/tool"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlDb *sql.DB
var MysqlDbErr error

const (
	USER_NAME  = "root"
	PASS_WORD  = "Wo123456"
	HOST       = "localhost"
	PORT       = "3306"
	DATABASE   = "tally"
	Table_User = "users"
)

func Link() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE)
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	if MysqlDbErr != nil {
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}
	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)
	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}
}

func QueryByField(field string) map[string]bool {
	names := make(map[string]bool)
	sql := fmt.Sprintf("SELECT %s FROM %s", field, Table_User)
	rows, _ := MysqlDb.Query(sql)
	var name string
	for rows.Next() {
		rows.Scan(&name)
		names[name] = true
	}
	return names
}

func RegisterUser(userInfo tool.TypeUserInfo) {
	sql := fmt.Sprintf(
		"INSERT INTO %s (nickName, password, signature, token) VALUES ('%s', '%s', '%s', '%s')",
		Table_User,
		userInfo.NickName,
		userInfo.Password,
		userInfo.Signature,
		userInfo.Token,
	)
	res, err := MysqlDb.Exec(sql)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("RegisterUser: %d", count)
}

func QueryUserInfo(fieldKey string, fieldValue string) (ok bool, userInfo tool.TypeUserInfoMysql) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE %s='%s'", Table_User, fieldKey, fieldValue)
	rows, err := MysqlDb.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	for rows.Next() {
		rows.Scan(&userInfo.Id, &userInfo.NickName, &userInfo.Password, &userInfo.Signature, &userInfo.Token)
		count += 1
	}
	ok = count > 0
	return
}
