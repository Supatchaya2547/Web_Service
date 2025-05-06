package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"onlinecourse/database"
	"onlinecourse/internal/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetType(c *gin.Context) {
	rows, err := database.DB.Query("select DISTINCT course_type from courses")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch data"})
		log.Println("error desc", err)
		return
	}
	defer rows.Close()

	// Initialize เป็น slice ว่างแทนที่จะเป็น nil
	results := make([]models.CType, 0)
	for rows.Next() {
		var course models.CType
		if err := rows.Scan(&course.Course_Type); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data scan error"})
			return
		}
		results = append(results, course)
	}

	c.JSON(http.StatusOK, results)
}
func GetInts(c *gin.Context) {
	rows, err := database.DB.Query("select DISTINCT course_instructor from courses")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch data"})
		log.Println("error desc", err)
		return
	}
	defer rows.Close()

	// Initialize เป็น slice ว่างแทนที่จะเป็น nil
	results := make([]models.CIntse, 0)
	for rows.Next() {
		var course models.CIntse
		if err := rows.Scan(&course.Course_Instructor); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data scan error"})
			return
		}
		results = append(results, course)
	}

	c.JSON(http.StatusOK, results)
}

// ข้อ 3
func GetData(c *gin.Context) {
	//รับ Query search
	// ค้นหาได้ 3 ค่า คือ ชื่อคอส วิชา คนสอน
	// search := ""
	// req := "SELECT * FROM courses WHERE 1=1"
	// var param string
	// if search = c.Query("n"); search != "" {
	// 	param = "'%" + search + "%'"
	// 	req += fmt.Sprintf(" and Course_Name ILIKE %s", param)
	// }
	// if search = c.Query("t"); search != "" {
	// 	param = "'%" + search + "%'"
	// 	req += fmt.Sprintf(" and Course_Type ILIKE %s", param)
	// }
	// if search = c.Query("i"); search != "" {
	// 	param = "'%" + search + "%'"
	// 	req += fmt.Sprintf(" and Course_Instructor ILIKE %s", param)
	// }

	// rows, err := database.DB.Query(req)
	// fmt.Println("🍤 request: ", req)

	baseSQL := `
	SELECT c.course_id, c.course_name, c.course_desc, c.thumbnail_url, 
		   c.course_type, c.course_instructor, c.profile_url, c.course_price, 
		   c.duration, c.rating, c.num_reviews, c.enrollment_count, 
		   c.created_at, c.updated_at, curl.detail_url
	FROM courses c 
	JOIN courses_url curl ON c.course_id = curl.course_id
	WHERE 1=1
	`

	filters := []string{}
	args := []interface{}{}
	i := 1

	if name := c.Query("n"); name != "" {
		filters = append(filters, fmt.Sprintf("AND c.course_name ILIKE $%d", i))
		args = append(args, "%"+name+"%")
		i++
	}
	if typ := c.Query("t"); typ != "" {
		filters = append(filters, fmt.Sprintf("AND c.course_type ILIKE $%d", i))
		args = append(args, "%"+typ+"%")
		i++
	}
	if instructor := c.Query("i"); instructor != "" {
		filters = append(filters, fmt.Sprintf("AND c.course_instructor ILIKE $%d", i))
		args = append(args, "%"+instructor+"%")
		i++
	}

	query := baseSQL + " " + strings.Join(filters, " ")

	rows, err := database.DB.Query(query, args...)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch data"})
		log.Println("error desc", err)
		return
	}
	defer rows.Close()

	// Initialize เป็น slice ว่างแทนที่จะเป็น nil
	results := make([]models.Courses, 0)
	for rows.Next() {
		var course models.Courses
		if err := rows.Scan(&course.Course_ID, &course.Course_Name, &course.Course_Desc, &course.Thumbnail_Url,
			&course.Course_Type, &course.Course_Instructor, &course.Profile_Url, &course.Course_Price,
			&course.Duration, &course.Rating, &course.Num_reviews, &course.Enrollment_count,
			&course.Created_at, &course.Updated_at, &course.Detail_url); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data scan error"})
			return
		}
		results = append(results, course)
	}

	c.JSON(http.StatusOK, results)
}

func GetAllCourses(c *gin.Context) {

	rows, err := database.DB.Query("select c.course_id, c.course_name, c.course_desc, c.thumbnail_url, c.course_type, c.course_instructor, c.profile_url, c.course_price, c.duration, c.rating, c.num_reviews, c.enrollment_count, c.created_at, c.updated_at, curl.detail_url from courses c join courses_url curl on c.course_id = curl.course_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถดึงข้อมูล courses ได้"})
		return
	}
	defer rows.Close()
	var courses []models.Courses
	for rows.Next() {
		var course models.Courses
		err := rows.Scan(&course.Course_ID, &course.Course_Name, &course.Course_Desc, &course.Thumbnail_Url,
			&course.Course_Type, &course.Course_Instructor, &course.Profile_Url, &course.Course_Price,
			&course.Duration, &course.Rating, &course.Num_reviews, &course.Enrollment_count,
			&course.Created_at, &course.Updated_at, &course.Detail_url)
		if err != nil {
			log.Println("Scan Error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการอ่านข้อมูล"})
			return
		}
		courses = append(courses, course)
	}
	// ตรวจสอบ error หลังวน loop เสร็จ
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการอ่านข้อมูล (rows.Err)"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func GetCourseByID(c *gin.Context) {
	courseID := c.Param("course_id")
	if courseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "course_id is required"})
		return
	}

	var course models.Courses
	query := `
	SELECT c.course_id, c.course_name, c.course_desc, c.thumbnail_url, 
	       c.course_type, c.course_instructor, c.profile_url, c.course_price, 
	       c.duration, c.rating, c.num_reviews, c.enrollment_count, 
	       c.created_at, c.updated_at, curl.detail_url 
	FROM courses c 
	JOIN courses_url curl ON c.course_id = curl.course_id 
	WHERE c.course_id = $1
	`

	err := database.DB.QueryRow(query, courseID).Scan(
		&course.Course_ID,
		&course.Course_Name,
		&course.Course_Desc,
		&course.Thumbnail_Url,
		&course.Course_Type,
		&course.Course_Instructor,
		&course.Profile_Url,
		&course.Course_Price,
		&course.Duration,
		&course.Rating,
		&course.Num_reviews,
		&course.Enrollment_count,
		&course.Created_at,
		&course.Updated_at,
		&course.Detail_url,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
			return
		}
		log.Println("QueryRow Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "เกิดข้อผิดพลาดในการอ่านข้อมูล"})
		return
	}

	c.JSON(http.StatusOK, course)
}
