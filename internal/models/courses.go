package models

type Courses struct {
	Course_ID         int     `json:"course_id"`   // ใช้ string เพราะว่าเป็น UUID
	Course_Name       string  `json:"course_name"` // ชื่อคอร์ส
	Course_Desc       string  `json:"course_desc"` // คำอธิบาย
	Thumbnail_Url     string  `json:"thumbnail_url"`
	Course_Type       string  `json:"course_type"`       // ประเภทคอร์ส (เช่น Programming, Design, etc.)
	Course_Instructor string  `json:"course_instructor"` // ชื่อผู้สอน
	Profile_Url       string  `json:"profile_url"`
	Course_Price      float64 `json:"course_price"`     // ราคา
	Duration          string  `json:"duration"`         // ระยะเวลา
	Rating            float64 `json:"rating"`           // คะแนนเฉลี่ย (1-5)
	Num_reviews       int     `json:"num_reviews"`      // จำนวนรีวิว
	Enrollment_count  int     `json:"enrollment_count"` // จำนวนคนที่ลงทะเบียน
	Detail_url        string  `json:"detail_url"`       // URL ของรายละเอียดคอร์ส
	Created_at        string  `json:"created_at"`
	Updated_at        string  `json:"updated_at"`
}

type CType struct {
	// Course_ID   int    `json:"course_id"`
	Course_Type string `json:"course_type"` // ประเภทคอร์ส (เช่น Programming, Design, etc.)
}
type CIntse struct {
	// Course_ID         int    `json:"course_id"`
	Course_Instructor string `json:"course_instructor"`
}

// course_id SERIAL PRIMARY KEY,
// course_name VARCHAR(255) NOT NULL,
// description TEXT,
// thumbnail_url VARCHAR(255),
// instructor_name VARCHAR(255),
// profile_url VARCHAR(255),
// duration VARCHAR(255),
// price DECIMAL(10,2),
// detail_url VARCHAR(255),
// rating DECIMAL(2,1),
// num_reviews INT,
// enrollment_count INT,
// created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
// updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
