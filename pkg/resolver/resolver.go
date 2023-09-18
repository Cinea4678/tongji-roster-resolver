package resolver

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
)

const sheet = "Sheet1"

func ResolveFile(fileName string) string {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Fatalln("文件打开失败：", err.Error())
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("文件关闭时出错：", err.Error())
		}
	}()

	course := resolveCourseData(f)
	course.Students = resolveStudent(f)

	json, _ := jsoniter.MarshalToString(course)
	return json
}

func resolveCourseData(file *excelize.File) Course {
	courseName, _ := file.GetCellValue(sheet, "A2")
	courseName = string([]rune(courseName)[5:])

	class, _ := file.GetCellValue(sheet, "E2")
	class = string([]rune(class)[4:])

	courseCode, _ := file.GetCellValue(sheet, "U2")
	courseCode = string([]rune(courseCode)[5:])

	courseId, _ := file.GetCellValue(sheet, "A3")
	courseId = string([]rune(courseId)[5:])

	school, _ := file.GetCellValue(sheet, "E3")
	school = string([]rune(school)[3:])

	teacher, _ := file.GetCellValue(sheet, "U3")
	teacher = string([]rune(teacher)[5:])

	studentNumStr, _ := file.GetCellValue(sheet, "A4")
	studentNumStr = string([]rune(studentNumStr)[5:])
	studentNum, _ := strconv.Atoi(studentNumStr)

	time, _ := file.GetCellValue(sheet, "U3")
	time = string([]rune(time)[5:])

	return Course{
		Name:       courseName,
		Id:         courseId,
		Code:       courseCode,
		StudentNum: studentNum,
		Class:      class,
		School:     school,
		Teacher:    teacher,
		Time:       time,
		Students:   []Student{},
	}
}

func resolveStudentRow(row []string) Student {
	index, _ := strconv.Atoi(row[0])
	studentId, _ := strconv.Atoi(row[1])
	name := row[2]
	englishName := row[3]
	gender := row[4]
	grade, _ := strconv.Atoi(row[5])
	school := row[6]
	major := row[7]
	isInternational := false
	if row[8] == "是" {
		isInternational = true
	}

	return Student{
		Index:           index,
		StudentId:       studentId,
		Name:            name,
		EnglishName:     englishName,
		Gender:          gender,
		Grade:           grade,
		School:          school,
		Major:           major,
		IsInternational: isInternational,
	}
}

func resolveStudent(file *excelize.File) []Student {
	rows, err := file.GetRows(sheet)
	if err != nil {
		log.Fatalln("读取表格失败")
	}

	students := make([]Student, 0, len(rows)-6)
	for i := 7; i < len(rows); i++ {
		students = append(students, resolveStudentRow(rows[i]))
	}

	return students
}
