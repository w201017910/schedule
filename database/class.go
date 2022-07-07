package database

import "fmt"

type ClassIn struct {
	Id          int
	Name        string
	Number      int
	CollegeId   int
	CollegeName string
	CampusName  string
}

func AddClass(name, number, college string) error {
	_, err := Db.Exec("INSERT INTO `schedule`.`class`(`name`, `number`, `collegeId`) VALUES ( ?, ?, ?)", name, number, college)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func AllClass() []ClassIn {
	rows, err := Db.Query("SELECT class.id,  class.`name`,  class.number,  class.collegeId,  college.`name` AS collegeName,  campus.`name` AS campusName FROM class INNER JOIN college ON  class.collegeId = college.id INNER JOIN campus ON  college.campusId = campus.id")
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var Classes []ClassIn
scan:
	if rows.Next() {
		classes := new(ClassIn)
		rows.Scan(&classes.Id, &classes.Name, &classes.Number, &classes.CollegeId, &classes.CollegeName, &classes.CampusName)

		Classes = append(Classes, *classes)
		goto scan
	}
	return Classes
}
func ChangeClass(id, name, number, college string) error {
	_, err := Db.Exec("UPDATE `schedule`.`class` SET `name` = ?, `number` = ?, `collegeId` = ?  WHERE `id` = ?", name, number, college, id)
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
