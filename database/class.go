package database

import "fmt"

type ClassIn struct {
	Id            int
	Name          string
	Number        int
	CollegeId     int
	CollegeName   string
	CampusName    string
	DefaultRoomId int
	RoomName      string
	RoomNumber    int
}

func AddClass(name, number, college, room string) error {
	_, err := Db.Exec("INSERT INTO `schedule`.`class`(`name`, `number`, `collegeId`, `defaultRoomId`) VALUES (?, ?, ?, ?)", name, number, college, room)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func AllClass() []ClassIn {
	rows, err := Db.Query("SELECT class.id,  class.`name`,  class.number,  class.collegeId,  college.`name` AS collegename,  campus.`name` AS campusname,  class.defaultRoomId,  classroom.`name` AS roomname,  classroom.number AS roomnumber FROM class INNER JOIN college ON  class.collegeId = college.id INNER JOIN classroom ON  class.defaultRoomId = classroom.id INNER JOIN campus ON college.campusId = campus.id ")
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var Classes []ClassIn
scan:
	if rows.Next() {
		classes := new(ClassIn)
		rows.Scan(&classes.Id, &classes.Name, &classes.Number, &classes.CollegeId, &classes.CollegeName, &classes.CampusName, &classes.DefaultRoomId, &classes.RoomName, &classes.RoomNumber)

		Classes = append(Classes, *classes)
		goto scan
	}
	return Classes
}
func ChangeClass(id, name, number, college, room string) error {
	_, err := Db.Exec("UPDATE `schedule`.`class` SET `name` = ?, `number` = ?, `collegeId` = ?, `defaultRoomId` = ? WHERE `id` = ?", name, number, college, room, id)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func DelClass(id string) error {
	_, err := Db.Exec("DELETE FROM `schedule`.`class` WHERE `id` = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
