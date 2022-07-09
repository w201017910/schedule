package database

import (
	"fmt"
	"strconv"
)

type College struct {
	Id    int
	Name  string
	Name1 string
}
type Room struct {
	Id          int
	Name        string
	Number      int
	CollegeId   string
	CollegeName string
	CampusName  string
}

func AllCollege() []College {
	rows, err := Db.Query("SELECT college.id, college.`name`, campus.`name` as name1 FROM  college INNER JOIN campus ON college.campusId = campus.id")
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var rooms []College
scan:
	if rows.Next() {
		room := new(College)
		rows.Scan(&room.Id, &room.Name, &room.Name1)

		rooms = append(rooms, *room)
		goto scan
	}
	return rooms
}
func AddRoom(name string, number string, college string) error {
	_, err := Db.Exec("INSERT INTO classroom(`name`, `number`, `collegeId`) VALUES (?, ?, ?)", name, number, college)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func AllRoom() []Room {
	rows, err := Db.Query("SELECT classroom.id,  classroom.`name`,  classroom.number, classroom.collegeId,  college.`name` AS collegeName,  campus.`name` AS campusName FROM classroom INNER JOIN college ON  classroom.collegeId = college.id INNER JOIN campus ON  college.campusId = campus.id")
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var rooms []Room
scan:
	if rows.Next() {
		room := new(Room)
		rows.Scan(&room.Id, &room.Name, &room.Number, &room.CollegeId, &room.CollegeName, &room.CampusName)

		rooms = append(rooms, *room)
		goto scan
	}
	return rooms
}
func ChangeRoom(ID, name, number, collegeId string) error {
	_, err := Db.Exec("UPDATE `schedule`.`classroom` SET `name` = ?, `number` = ?, `collegeId` = ? WHERE `id` = ?", name, number, collegeId, ID)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func DelRoom(ID string) error {
	_, err := Db.Exec("DELETE FROM `schedule`.`classroom` WHERE `id` = ?", ID)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func SearchRoom(cla []string, roomId string) []string {
	collegeId, num := SearchCollege(cla)
	sql := "SELECT id FROM classroom WHERE classroom.number >= ? AND classroom.collegeId IN ("
	fmt.Println(collegeId)
	for i := 0; i < len(collegeId)-1; i++ {

		sql = sql + strconv.Itoa(collegeId[i]) + ","
	}
	sql = sql + strconv.Itoa(collegeId[len(collegeId)-1]) + ")"
	fmt.Println(sql)
	rows, err := Db.Query(sql, num)
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var rooms []string
scan:
	if rows.Next() {
		var room string
		rows.Scan(&room)
		if room == roomId {
			goto scan
		}
		rooms = append(rooms, room)
		goto scan
	}
	return rooms
}
func SearchCollege(cla []string) ([]int, int) {
	var c []int
	num := 0
	for _, v := range cla {
		rows, err := Db.Query("SELECT collegeId,number FROM class WHERE id=?", v)
		defer CloseConnection(rows)
		if err != nil {
			fmt.Println(err)
		}

		if rows.Next() {
			var test int
			var num1 int
			rows.Scan(&test, &num1)
			num = num + num1
			c = append(c, test)

		}
	}

	return c, num
}
