package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"schedule/config"
	"schedule/database"
	"strconv"
	"strings"
)

func Login(c *gin.Context) {
	account := c.PostForm("account")
	passwd := c.PostForm("passwd")

	if account == config.Account && passwd == config.PassWord {
		c.SetCookie("name", account, 1000, "/", "localhost", false, true)

		c.JSON(200, true)
		return
	}

	c.JSON(200, false)
	return
}
func AddTeacher(c *gin.Context) {
	name := c.PostForm("name")
	sex := c.PostForm("sex")
	workID := c.PostForm("workID")
	phone := c.PostForm("phone")
	email := c.PostForm("email")

	judge := database.TeacherAdd(name, sex, workID, phone, email)
	if judge {
		c.JSON(200, true)
		return
	}
	c.JSON(200, false)
	return
}
func AllTeacher(c *gin.Context) {
	i := database.AllTeacher()

	c.JSON(200, i)
}
func ChangeTeachers(c *gin.Context) {

	err := database.ChangeTeacher(c.PostForm("Id"), c.PostForm("Name"), c.PostForm("Sex"), c.PostForm("Phone"), c.PostForm("Email"), c.PostForm("WorkId"))
	if err != nil {
		c.JSON(200, "修改失败")
		return
	}
	c.JSON(200, "修改成功")
}
func AddRoom(c *gin.Context) {
	name := c.PostForm("name")
	college := c.PostForm("college")
	number := c.PostForm("number")
	collegeId := strings.Split(college, "-")[0]
	err := database.AddRoom(name, number, collegeId)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
func ChangeRoom(c *gin.Context) {
	id := c.PostForm("Id")
	name := c.PostForm("Name")
	number := c.PostForm("Number")
	CollegeName := c.PostForm("CollegeName")
	collegeId := strings.Split(CollegeName, "-")[0]
	err := database.ChangeRoom(id, name, number, collegeId)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
func DelRoom(c *gin.Context) {
	id := c.PostForm("Id")
	err := database.DelRoom(id)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
func AddClass(c *gin.Context) {
	name := c.PostForm("name")
	number := c.PostForm("number")
	college := c.PostForm("college")

	collegeId := strings.Split(college, "-")[0]

	err := database.AddClass(name, number, collegeId)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
func ChangeClass(c *gin.Context) {
	id := c.PostForm("Id")
	name := c.PostForm("Name")
	number := c.PostForm("Number")
	college := c.PostForm("college")

	collegeId := strings.Split(college, "-")[0]

	err := database.ChangeClass(id, name, number, collegeId)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
func DelClass(c *gin.Context) {
	id := c.PostForm("Id")
	err := database.DelClass(id)
	if err != nil {
		c.JSON(200, false)
		return
	}
	c.JSON(200, true)
}
func AddCurriculum(c *gin.Context) {
	ban := c.PostForm("ban")[1:]
	className := c.PostForm("className")[1:]
	name := c.PostForm("name")
	teacher := c.PostForm("teacher")
	timeLong := c.PostForm("timeLong")
	totalTime := c.PostForm("totalTime")
	year := c.PostForm("year")
	origin := c.PostForm("origin")
	room := c.PostForm("room")
	interval, _ := strconv.Atoi(c.PostForm("interval"))
	roomId := strings.Split(room, "-")[0]
	err := database.AddCurriculum(name, teacher, timeLong, totalTime, year, className, ban, origin, roomId, interval)
	courseId := database.LastIdByCourse()

	classes := strings.Split(className, "-")

	TimeLong, _ := strconv.Atoi(timeLong)
	weekTime, _ := strconv.Atoi(totalTime)
	originWeek, _ := strconv.Atoi(origin)

	if err != nil {
		c.JSON(200, false)
		return
	}
	if len(classes) > 1 {
		jj := mulSort(strconv.Itoa(courseId), name, teacher, TimeLong, weekTime, year, originWeek, roomId, interval, classes, ban)
		if !jj {
			err = database.DelCurriculum(courseId)
		}
		c.JSON(200, jj)
	} else {
		jj := Sort(strconv.Itoa(courseId), name, teacher, TimeLong, weekTime, year, originWeek, roomId, interval, classes, ban)
		if !jj {
			err = database.DelCurriculum(courseId)
		}
		c.JSON(200, jj)
	}
	c.JSON(200, false)
}
func Sort(courseId, name, teacher string, timeLong, weekTime int, year string, originWeek int, roomId string, interval int, classes []string, forbidden string) bool {
	count := 0

	var cc []CC
	var cla [16][20]database.Course
	cla = database.CreateByClass(classes[0], year)
	var rooms []string
	rooms = append(rooms, roomId)
	rooms = append(rooms, database.SearchRoom(classes, roomId)...)
	room := database.CreateByRoom(roomId, year)
	tea := database.CreateByTeacher(teacher, year)
	for _, v := range rooms {
		for j := 0; j < 20; j++ {
			if strings.Contains(forbidden, strconv.Itoa(j)) {
				continue
			}

			judge := false
			for i := originWeek - 1; i < originWeek+timeLong-1; i = i + interval + 1 {
				fmt.Println("i:", i, "j:", j, "room:", room[i][j].Exist, "tea:", tea[i][j].Exist, "cla:", cla[i][j].Exist)
				if (room[i][j].Exist || tea[i][j].Exist) || cla[i][j].Exist {

					judge = true
					break
				}
			}
			if !judge {
				cc = append(cc, CC{
					courseId: courseId,
					clasTime: j,
					classes:  classes[0],
					RoomId:   v,
					mul:      1,
				})

				count++
				if count == weekTime {
					goto scan
				}
			}
		}
	}
scan:

	if count == weekTime {

		for _, v := range cc {
			database.AddCourse(v.courseId, v.clasTime, v.classes, v.RoomId, v.mul)
		}
		return true
	} else {

	}
	return false
}

type CC struct {
	courseId string
	clasTime int
	classes  string
	RoomId   string
	mul      int
}

func mulSort(courseId, name, teacher string, timeLong, weekTime int, year string, originWeek int, roomId string, interval int, classes []string, forbidden string) bool {
	count := 0

	var cc []CC
	var cla [][16][20]database.Course
	for i := 0; i < len(classes); i++ {
		cla = append(cla, database.CreateByClass(classes[i], year))
	}
	var rooms []string
	rooms = append(rooms, roomId)
	rooms = append(rooms, database.SearchRoom(classes, roomId)...)
	room := database.CreateByRoom(roomId, year)
	tea := database.CreateByTeacher(teacher, year)

	for _, v := range rooms {
		for j := 0; j < 20; j++ {
			if strings.Contains(forbidden, strconv.Itoa(j)) {
				continue
			}

			judge := false
			for i := originWeek - 1; i < originWeek+timeLong-1; i = i + interval + 1 {
				fmt.Println("i:", i, "j:", j, "room:", room[i][j].Exist, "tea:", tea[i][j].Exist, "cla:", Judge(cla, i, j))
				if (room[i][j].Exist || tea[i][j].Exist) || Judge(cla, i, j) {

					judge = true
					break
				}
			}
			if !judge {
				for i := 0; i < len(classes); i++ {

					cc = append(cc, CC{
						courseId: courseId,
						clasTime: j,
						classes:  classes[i],
						RoomId:   v,
						mul:      1,
					})

				}
				count++
				if count == weekTime {
					goto scan
				}
			}
		}
	}
scan:

	if count == weekTime {

		for _, v := range cc {
			database.AddCourse(v.courseId, v.clasTime, v.classes, v.RoomId, v.mul)
		}
		return true
	} else {

	}
	return false

}
func Judge(a [][16][20]database.Course, i, j int) bool {
	d := false
	for c := 0; c < len(a); c++ {
		if a[c][i][j].Exist {
			d = true
			break
		}
	}
	return d
}
