package database

import "fmt"

type Curriculum struct {
	Id            int
	Name          string
	TeacherId     int
	TimeLong      int
	TotalTime     int
	Semester      string
	Cla           string
	ForbiddenTime string
	OriginWeek    int
	RoomId        int
	interval      int
}

func AddCurriculum(name, teacher, timeLong, totalTime, semester, cla, forbid, origin, roomId string, interval int) error {
	_, err := Db.Exec("INSERT INTO `schedule`.`curriculum`(`name`, `teacherId`, `timeLong`, `totalTime`, `semester`, `cla`, `forbiddenTime`, `originWeek`,roomId,`interval`) VALUES (?, ?, ?, ?, ?, ?, ?, ?,?,?)", name, teacher, timeLong, totalTime, semester, cla, forbid, origin, roomId, interval)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func QueryCurriculum(id int) Curriculum {
	rows, err := Db.Query("SELECT * FROM curriculum WHERE curriculum.id = ?", id)
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var curriculum Curriculum

	if rows.Next() {

		rows.Scan(&curriculum.Id, &curriculum.Name, &curriculum.TeacherId, &curriculum.TimeLong, &curriculum.TotalTime, &curriculum.Semester, &curriculum.Cla, &curriculum.ForbiddenTime, &curriculum.OriginWeek, &curriculum.RoomId, &curriculum.interval)

	}
	return curriculum
}
func LastIdByCourse() int {
	rows, err := Db.Query("SELECT MAX(id) from curriculum")
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var course int

	if rows.Next() {

		rows.Scan(&course)
	}

	return course
}
func DelCurriculum(id int) error {
	_, err := Db.Exec("DELETE FROM curriculum WHERE id=?", id)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
