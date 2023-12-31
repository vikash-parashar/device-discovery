package handlers

import (
	"net/http"

	"github.com/vikash-parashar/asset-locator/db"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"status": "success", "message": "application is running"})
	c.HTML(http.StatusOK, "healthcheck.html", nil)
}

func RenderIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func RenderForgotPasswordPage(c *gin.Context) {
	c.HTML(http.StatusOK, "forgot_password.html", nil)
}
func RenderAboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", nil)
}
func RenderGetHelpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "help.html", nil)
}
func RenderHomePage(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", nil)
	}
}
