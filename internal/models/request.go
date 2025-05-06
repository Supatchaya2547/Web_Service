package models

// รวม request GET ต่างๆ

type Click_logs struct {
	Affiliate_ID string `json:"affiliate_id"`
	Course_ID    int    `json:"course_id"`
	Click        int    `json:"click_date"`
}

// log การ request

type RequestLog struct {
	AffiliateID string `json:"affiliate_id"`
	Method      string `json:"method"`
	Action      string `json:"action"`    // search, click
	Parameter   string `json:"parameter"` // คำค้นหา
	Timestamp   string `json:"timestamp"` // วันและเวลา
}
