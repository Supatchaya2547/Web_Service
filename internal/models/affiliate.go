package models

// ใส่หลัง json ให้เหมือนใน DB
type Affiliates struct {
	Affiliate_ID    string `json:"affiliate_id"`
	Affiliate_Name  string `json:"affiliate_name"`
	Affiliate_Email string `json:"affiliate_email"`
}

type Affiliate_Url struct {
	Url_id        int    `json:"url_id"`
	Affiliate_ID  string `json:"affiliate_id"`
	Affiliate_Url string `json:"aff_url"`
	Clicks        int    `json:"clicks"`
}

type AffUrl_req struct {
	Affiliate_Url string `json:"aff_url"`
	Action        string `json:"action"`
}
