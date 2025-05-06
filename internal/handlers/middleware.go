// ยังไม่แก้
package handlers

// ต้องมี Middle -> Best practice
import (
	"database/sql"
	"fmt"
	"net/http"
	"onlinecourse/database"
	"time"

	"github.com/gin-gonic/gin"
)

// ข้อ 5 request logs ----> ยังไม่ทดลอง
func RequestLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		search := ""
		s := ""
		if s = c.Query("n"); s != "" {
			search += fmt.Sprint("s= " + s + ", ")
		}
		if s = c.Query("t"); s != "" {
			search += fmt.Sprint("t= " + s + ", ")
		}
		if s = c.Query("i"); s != "" {
			search += fmt.Sprint("i= " + s)
		}
		if search == "" {
			search = "All Course"
		}
		str_id, exists := c.Get("affiliate_id")
		// fmt.Printf("value: %v\n type: %T\n", str_id, str_id)

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		_, err := database.DB.Exec(
			"INSERT INTO request_logs (affiliate_id,method, action,parameter,timestamp)  VALUES ($1 ,$2, $3, $4, $5)",
			str_id, c.Request.Method, c.Request.URL.Path, search, time.Now())
		// c.Request.Method

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Next()
	}
}

// ข้อ 7
func ClickLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		course_id := c.Param("id")
		url := c.Query("url")
		act := c.Query("act")

		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required (Your Website URL)"})
			return
		}
		if act == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "action is required (Ex click, search)"})
			return
		}
		// fmt.Println("⭐ req.URL : ", req.Affiliate_Url)
		// fmt.Println("⭐ req.Action : ", req.Action)
		// ตรวจสอบ affiliate_id แล้วเก็บค่าเอาไว้เตรียมบันทึกลง DB
		var id string
		r := database.DB.QueryRow(`select affiliate_id from affiliate_url where aff_url = $1 `, url)
		err := r.Scan(&id)
		// fmt.Printf("id:%s\n", id)
		if err == sql.ErrNoRows {
			fmt.Println("ไม่พบผู้ใช้ที่มี ID:", id)
			c.JSON(http.StatusBadRequest, gin.H{"error": "URL ของคุณผิด หรือไม่มีในระบบ"})
			return
		}

		// บันทึก click log
		_, err = database.DB.Exec(
			`INSERT INTO click_logs (affiliate_id,course_id) values ($1,$2)`,
			id, course_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		// เพิ่มจำนวน click ทั้งหมด
		_, err = database.DB.Exec(`UPDATE affiliate_url SET clicks = clicks + 1 WHERE affiliate_id = $1`, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update click count"})
			return
		}
		c.Next()
	}

}
