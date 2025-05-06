package handlers

import (
	"net/http"
	"onlinecourse/database"
	"onlinecourse/internal/models"

	"github.com/gin-gonic/gin"
)

// Gen APIKey
// func generateAPIKey() string {
// 	bytes := make([]byte, 16)
// 	rand.Read(bytes)
// 	return hex.EncodeToString(bytes)
// }

func Register(c *gin.Context) {
	affiliateID, ok1 := c.Get("affiliate_id")
	username, ok2 := c.Get("username")
	email, ok3 := c.Get("email")

	if !ok1 || !ok2 || !ok3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user context"})
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM affiliates WHERE affiliate_id=$1)", affiliateID).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if !exists {
		_, err := database.DB.Exec(`
			INSERT INTO affiliates (affiliate_id, affiliate_name, affiliate_email) 
			VALUES ($1, $2, $3)
		`, affiliateID, username, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Insert failed"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Affiliate registered successfully",
		"data": gin.H{
			"affiliate_id":    affiliateID,
			"affiliate_name":  username,
			"affiliate_email": email,
		},
	})
}

func Get_Url(c *gin.Context) {

}

func Url_Register(c *gin.Context) {
	var req models.AffUrl_req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// _, err := database.DB.Exec("INSERT INTO Affiliate_Url VALUES ($1, $2)", input.Email, input.Name)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "User created"})

	// Url_id        int    `json:"url_id"`
	// Affiliate_ID  string `json:"affiliate_id"`
	// Affiliate_Url string `json:"aff_url"`
	// Clicks        int    `json:"clicks"`
	// Parameter     string `json:"parameter"`
}
