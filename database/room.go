package database

import "fmt"

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
