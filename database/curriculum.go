package database

import "fmt"

func AddCurriculum(name, teacher, timeLong, totalTime, semester, cla, forbid string) error {
	_, err := Db.Exec("INSERT INTO `schedule`.`curriculum`(`name`, `teacherId`, `timeLong`, `totalTime`, `semester`, `cla`, `forbiddenTime`) VALUES (?, ?, ?, ?, ?, ?, ?)", name, teacher, timeLong, totalTime, semester, cla, forbid)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
