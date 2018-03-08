package golang_helpers

import (
	"database/sql"
	"time"
	"log"
)

var DbConn *sql.DB

/**
	Function to Open a Connection Pool
*/
func DbOpenConn() (*sql.DB, error) {
	connstr := "db_username:db_password@tcp(db_server:3306)/db_name?charset=utf8"
	DbConn, err := sql.Open("mysql", connstr)
	CheckErr(err)

	/*
	We may or may not use these settings but these setting helped in fixing two MySQL related bugs I faced while development
	- Too Many Connections
	- Aborted Connection
	*/
	//No Idle connection
	DbConn.SetMaxIdleConns(0)

	//Should be higher for production code. But also change the same in MySQL settings (/etc/mysql/my.cnf)
	DbConn.SetMaxOpenConns(100)

	//Connection dies after 5 secs
	DbConn.SetConnMaxLifetime(time.Second * 5)

	return DbConn, err
}

//In case the Connection died, create a new connection
func FixConn() (*sql.DB) {
	checkping := DbConn.Ping()

	var err error
	if checkping != nil {
		//if error
		log.Println("DbConn.Ping failed here. Creating a new connection.")
		DbConn, err = DbOpenConn()
		CheckErr(err)
		return DbConn
	}
	return DbConn
}

func DbExecute(db *sql.DB, query string, params ...interface{}) (sql.Result, error) {
	db = FixConn()

	//Prepared statements
	stmt, err := db.Prepare(query)
	CheckErr(err)

	//Never use .Query as it will not close the connection
	res, err := stmt.Exec(params...)
	CheckErr(err)

	//Always close the stmt or rows (you may use defer too)
	stmt.Close()
	return res, err
}

func DbInsert(db *sql.DB, query string, params ...interface{}) (int64, error) {
	db = FixConn()

	res, err := DbExecute(db, query, params...)
	CheckErr(err)

	//return Last Insert Id from MySQL
	id, err := res.LastInsertId()
	CheckErr(err)

	return id, err
}

func DbUpdate(db *sql.DB, query string, params ...interface{}) (int64, error) {
	db = FixConn()

	res, err := DbExecute(db, query, params...)
	CheckErr(err)

	affect, err := res.RowsAffected()
	CheckErr(err)

	return affect, err
}

func DbDelete(db *sql.DB, query string, params ...interface{}) (int64, error) {
	var affect int64
	var err error
	db = FixConn()

	affect, err = DbUpdate(db, query, params...)
	CheckErr(err)

	return affect, err
}

//Fetch rows but without any parameters
func DbQueryGetRows(db *sql.DB, query string) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error

	db = FixConn()
	rows, err = db.Query(query)
	CheckErr(err)

	return rows, err
}

//Fetch rows but WITH parameters
func DbQueryGetRowsParams(db *sql.DB, query string, params ...interface{}) (*sql.Rows, error) {
	db = FixConn()

	stmt, err1 := db.Prepare(query)
	CheckErr(err1)

	rows, err2 := stmt.Query(params...)
	CheckErr(err2)

	stmt.Close()
	return rows, err2
}

func DbQueryGetRowsParamsSingle(db *sql.DB, query string, params ...interface{}) (*sql.Rows, error) {
	var err error

	db = FixConn()
	CheckErr(err)

	stmt, err := db.Prepare(query)
	CheckErr(err)

	rows, err := stmt.Query(params...)
	CheckErr(err)

	stmt.Close()
	return rows, err
}

func FetchRowsByID(db *sql.DB, query string, param string) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error

	db = FixConn()
	CheckErr(err)

	rows, err = db.Query(query, param)
	CheckErr(err)

	return rows, err
}

func DbCloseConn(db *sql.DB) {
	db.Close()
}

func CountRows(rows *sql.Rows) int {
	count := 0
	for rows.Next() {
		count++
	}
	return count
}

func CheckErr(err error) {
	if err != nil {
		log.Panic(err)
		return
	}
}

func SelectExample() {

	db := DbConn

	id := "10"
	sel_db, err := DbQueryGetRowsParams(db, "SELECT name, age FROM mytable where id=?", id)
	CheckErr(err)
	defer sel_db.Close()

	for sel_db.Next() {
		var age int
		var name string
		err = sel_db.Scan(&name, &age)
		CheckErr(err)
	}
}

func InsertExample() {

	db := DbConn

	name := "My Name"
	age := "20"
	_, err := DbInsert(db, "INSERT INTO mytable VALUES(null,?,?)", name, age)
	CheckErr(err)
}

func UpdateExample() {

	db := DbConn

	name := "My Name"
	age := "20"
	id := "1"
	_, err := DbUpdate(db, "update alp_sets set name=?, age=? where id = ?",  name, age, id)
	CheckErr(err)
}

//You may run this in Main method of your project
func RunInMainMethod() {
	if DbConn == nil {
		log.Println("Creating new Connection. DbConn is Dead")
		var err error
		DbConn, err = DbOpenConn()
		CheckErr(err)
	}
}