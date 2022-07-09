package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"schedule/config"
)

var Db *sql.DB

type Teacher struct {
	Id     int
	Name   string
	Sex    string
	WorkId string
	Phone  string
	Email  string
}

func Init() {
	conSte := "root:" + config.DatabasePasswd + "@tcp(127.0.0.1:3306)/schedule"
	fmt.Println(conSte)
	var err error
	Db, err = sql.Open("mysql", conSte)
	Db.SetMaxOpenConns(2000)
	Db.SetMaxIdleConns(1000)
	if err != nil {
		fmt.Print(err)
	}
}
func CloseConnection(rows *sql.Rows) {
	rows.Close()
}
func TeacherAdd(name, sex, workID, phone, email string) bool {

	_, err := Db.Exec("INSERT teacher  VALUES (null,?, ?, ?, ?, ?)", name, sex, phone, email, workID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func AllTeacher() []Teacher {
	rows, err := Db.Query("select * from teacher ")
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var teachers []Teacher
scan:
	if rows.Next() {
		teacher := new(Teacher)
		rows.Scan(&teacher.Id, &teacher.Name, &teacher.Sex, &teacher.Phone, &teacher.Email, &teacher.WorkId)

		teachers = append(teachers, *teacher)
		goto scan
	}
	return teachers
}
func ChangeTeacher(ID, name, sex, phone, email, workId string) error {
	_, err := Db.Exec("UPDATE `schedule`.`teacher` SET `name` = ?, `sex` = ?, `phone` = ?, `email` = ?, `workID` = ? WHERE `id` = ?", name, sex, phone, email, workId, ID)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
