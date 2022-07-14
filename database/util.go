package database

import "fmt"

type Course struct {
	Id            int
	CurId         int
	ClassTime     int
	ClassId       int
	RoomId        int
	mul           int
	Name          string
	TeacherId     int
	timeLong      int
	WeekTime      int
	semester      string
	cla           string
	forbiddenTime string
	originWeek    int
	firstRoom     int
	Interval      int
	Exist         bool
}

func AllCourseByClass(ClassId, semester string) []Course {
	rows, err := Db.Query("SELECT DISTINCT course.id,  course.curid,  course.classTime,  course.classId,  course.roomId, course.mul , curriculum.`name`,  curriculum.teacherId,  curriculum.timeLong,  curriculum.totalTime, curriculum.semester,  curriculum.cla,  curriculum.forbiddenTime,  curriculum.originWeek,  curriculum.roomId AS firstRoom,curriculum.interval FROM course INNER JOIN curriculum ON  course.curid = curriculum.id WHERE course.classId = ? and curriculum.semester=?", ClassId, semester)
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var courses []Course
scan:
	if rows.Next() {
		course := new(Course)
		rows.Scan(&course.Id, &course.CurId, &course.ClassTime, &course.ClassId, &course.RoomId, &course.mul, &course.Name, &course.TeacherId, &course.timeLong, &course.WeekTime, &course.semester, &course.cla, &course.forbiddenTime, &course.originWeek, &course.firstRoom, &course.Interval)
		course.Exist = true
		courses = append(courses, *course)
		goto scan
	}
	return courses
}
func CreateByClass(ClassId, semester string) [16][20]Course {
	courses := AllCourseByClass(ClassId, semester)
	var classObject [16][20]Course
	for i := 0; i < len(courses); i++ {
		for j := courses[i].originWeek - 1; j < courses[i].originWeek-1+courses[i].timeLong; j = j + 1 + courses[i].Interval {
			classObject[j][courses[i].ClassTime] = courses[i]
		}
	}
	return classObject
}
func AllCourseByTeacher(teacherId, semester string) []Course {
	rows, err := Db.Query("SELECT DISTINCT course.id,  course.curid,  course.classTime,  course.classId,  course.roomId,course.mul,  curriculum.`name`,  curriculum.teacherId,  curriculum.timeLong,  curriculum.totalTime, curriculum.semester,  curriculum.cla,  curriculum.forbiddenTime,  curriculum.originWeek,  curriculum.roomId AS firstRoom,curriculum.interval FROM course INNER JOIN curriculum ON  course.curid = curriculum.id WHERE curriculum.teacherId = ? and curriculum.semester=?", teacherId, semester)
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var courses []Course
scan:
	if rows.Next() {
		course := new(Course)
		rows.Scan(&course.Id, &course.CurId, &course.ClassTime, &course.ClassId, &course.RoomId, &course.mul, &course.Name, &course.TeacherId, &course.timeLong, &course.WeekTime, &course.semester, &course.cla, &course.forbiddenTime, &course.originWeek, &course.firstRoom, &course.Interval)
		course.Exist = true
		courses = append(courses, *course)
		goto scan
	}
	return courses
}
func CreateByTeacher(teacherId, semester string) [16][20]Course {
	courses := AllCourseByTeacher(teacherId, semester)
	var classObject [16][20]Course
	for i := 0; i < len(courses); i++ {
		for j := courses[i].originWeek - 1; j < courses[i].originWeek-1+courses[i].timeLong; j = j + 1 + courses[i].Interval {
			classObject[j][courses[i].ClassTime] = courses[i]
		}
	}
	return classObject
}
func AllCourseByRoom(roomId, semester string) []Course {
	rows, err := Db.Query("SELECT DISTINCT course.id,  course.curid,  course.classTime,  course.classId,  course.roomId, course.mul, curriculum.`name`,  curriculum.teacherId,  curriculum.timeLong,  curriculum.totalTime, curriculum.semester,  curriculum.cla,  curriculum.forbiddenTime,  curriculum.originWeek,  curriculum.roomId AS firstRoom,curriculum.interval FROM course INNER JOIN curriculum ON  course.curid = curriculum.id WHERE course.roomId = ? and curriculum.semester=?", roomId, semester)
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var courses []Course
scan:
	if rows.Next() {
		course := new(Course)
		rows.Scan(&course.Id, &course.CurId, &course.ClassTime, &course.ClassId, &course.RoomId, &course.mul, &course.Name, &course.TeacherId, &course.timeLong, &course.WeekTime, &course.semester, &course.cla, &course.forbiddenTime, &course.originWeek, &course.firstRoom, &course.Interval)
		course.Exist = true
		courses = append(courses, *course)
		goto scan
	}
	return courses
}
func CreateByRoom(roomId, semester string) [16][20]Course {
	courses := AllCourseByRoom(roomId, semester)
	var classObject [16][20]Course
	for i := 0; i < len(courses); i++ {
		for j := courses[i].originWeek - 1; j < courses[i].originWeek-1+courses[i].timeLong; j = j + 1 + courses[i].Interval {
			classObject[j][courses[i].ClassTime] = courses[i]
		}
	}
	return classObject
}
func AddCourse(curId string, classTime int, classId string, roomId string, mul int) error {
	_, err := Db.Exec("INSERT INTO `schedule`.`course`(`curid`, `classTime`, `classId`, `roomId`, `mul`) VALUES (?,?,?,?,?)", curId, classTime, classId, roomId, mul)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
