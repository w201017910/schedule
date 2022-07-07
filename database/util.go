package database

import "fmt"

type Course struct {
	Id            int
	CurId         int
	ClassTime     int
	ClassId       int
	RoomId        int
	Name          string
	teacherId     int
	timeLong      int
	WeekTime      int
	semester      string
	cla           string
	forbiddenTime string
	originWeek    int
	firstRoom     int
}

func AllCourseByClass(ClassId int) []Course {
	rows, err := Db.Query("SELECT DISTINCT course.id,  course.curid,  course.classTime,  course.classId,  course.roomId,  curriculum.`name`,  curriculum.teacherId,  curriculum.timeLong,  curriculum.totalTime, curriculum.semester,  curriculum.cla,  curriculum.forbiddenTime,  curriculum.originWeek,  curriculum.roomId AS firstRoom FROM course INNER JOIN curriculum ON  course.curid = curriculum.id WHERE course.classId = ? ", ClassId)
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var courses []Course
scan:
	if rows.Next() {
		course := new(Course)
		rows.Scan(&course.Id, &course.CurId, &course.ClassTime, &course.ClassId, &course.RoomId, &course.Name, course.teacherId, course.timeLong, course.WeekTime, course.semester, course.cla, course.forbiddenTime, course.originWeek, course.firstRoom)

		courses = append(courses, *course)
		goto scan
	}
	return courses
}
func CreateByClass(ClassId int) [16][20]Course {
	courses := AllCourseByClass(ClassId)
	var classObject [16][20]Course
	for i := 0; i < len(courses); i++ {
		for j := courses[i].originWeek - 1; j < courses[i].originWeek-1+courses[i].timeLong; j++ {
			classObject[j][courses[i].ClassTime] = courses[i]
		}
	}
	return classObject
}
func AllCourseByTeacher(teacherId int) []Course {
	rows, err := Db.Query("SELECT DISTINCT course.id,  course.curid,  course.classTime,  course.classId,  course.roomId,  curriculum.`name`,  curriculum.teacherId,  curriculum.timeLong,  curriculum.totalTime, curriculum.semester,  curriculum.cla,  curriculum.forbiddenTime,  curriculum.originWeek,  curriculum.roomId AS firstRoom FROM course INNER JOIN curriculum ON  course.curid = curriculum.id WHERE curriculum.teacherId = ? ", teacherId)
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var courses []Course
scan:
	if rows.Next() {
		course := new(Course)
		rows.Scan(&course.Id, &course.CurId, &course.ClassTime, &course.ClassId, &course.RoomId, &course.Name, course.teacherId, course.timeLong, course.WeekTime, course.semester, course.cla, course.forbiddenTime, course.originWeek, course.firstRoom)

		courses = append(courses, *course)
		goto scan
	}
	return courses
}
func CreateByTeacher(teacherId int) [16][20]Course {
	courses := AllCourseByTeacher(teacherId)
	var classObject [16][20]Course
	for i := 0; i < len(courses); i++ {
		for j := courses[i].originWeek - 1; j < courses[i].originWeek-1+courses[i].timeLong; j++ {
			classObject[j][courses[i].ClassTime] = courses[i]
		}
	}
	return classObject
}
func AllCourseByRoom(roomId int) []Course {
	rows, err := Db.Query("SELECT DISTINCT course.id,  course.curid,  course.classTime,  course.classId,  course.roomId,  curriculum.`name`,  curriculum.teacherId,  curriculum.timeLong,  curriculum.totalTime, curriculum.semester,  curriculum.cla,  curriculum.forbiddenTime,  curriculum.originWeek,  curriculum.roomId AS firstRoom FROM course INNER JOIN curriculum ON  course.curid = curriculum.id WHERE course.roomId = 3 ", roomId)
	defer CloseConnection(rows)
	if err != nil {
		fmt.Println(err)
	}
	var courses []Course
scan:
	if rows.Next() {
		course := new(Course)
		rows.Scan(&course.Id, &course.CurId, &course.ClassTime, &course.ClassId, &course.RoomId, &course.Name, course.teacherId, course.timeLong, course.WeekTime, course.semester, course.cla, course.forbiddenTime, course.originWeek, course.firstRoom)

		courses = append(courses, *course)
		goto scan
	}
	return courses
}
func CreateByRoom(roomId int) [16][20]Course {
	courses := AllCourseByRoom(roomId)
	var classObject [16][20]Course
	for i := 0; i < len(courses); i++ {
		for j := courses[i].originWeek - 1; j < courses[i].originWeek-1+courses[i].timeLong; j++ {
			classObject[j][courses[i].ClassTime] = courses[i]
		}
	}
	return classObject
}
