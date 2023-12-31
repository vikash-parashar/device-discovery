package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikash-parashar/asset-locator/db"
	"github.com/vikash-parashar/asset-locator/handlers"
	"github.com/vikash-parashar/asset-locator/middleware"
)

func SetupRoutes(r *gin.Engine, dbConn *db.DB) {
	// Unprotected routes
	r.GET("/", handlers.RenderIndexPage)
	r.GET("/signup", handlers.RenderIndexPage)
	r.GET("/login", handlers.RenderIndexPage)
	r.GET("/about", handlers.RenderAboutPage)
	r.GET("/help", handlers.RenderGetHelpPage)
	r.GET("/health-check", handlers.HealthCheck)
	r.POST("/signup", handlers.SignUp(dbConn))

	r.POST("/login", handlers.Login(dbConn))
	r.POST("/logout", handlers.Logout())
	r.GET("/forget-password-page", handlers.RenderForgotPasswordPage)
	r.POST("/forget-password", handlers.ForgotPassword(dbConn))
	r.GET("/reset-password", handlers.RenderResetPasswordPage)
	r.POST("/reset-password", handlers.ResetPassword(dbConn))

	// Protected routes
	protected := r.Group("/api/v1", middleware.AuthMiddleware("admin", "general"))

	// Homepage
	protected.GET("/homepage", handlers.RenderHomePage(dbConn))

	// for fetching disk details from external server
	protected.GET("/disk-details", handlers.FetchDisks)

	// User
	protected.GET("/get-current-user", handlers.GetCurrentUser(dbConn))

	// Location Details
	protected.GET("/location-details", handlers.GetLocationDetails(dbConn))
	protected.POST("/location-details", handlers.CreateNewLocationDetails(dbConn))
	protected.PATCH("/location-details/:id", handlers.UpdateDeviceLocationDetail(dbConn))
	protected.DELETE("/location-details/:id", handlers.DeleteDeviceLocationDetail(dbConn))
	protected.GET("/location-details/pdf", handlers.DownloadDeviceLocationDetailPDF(dbConn))
	protected.GET("/location-details/excel", handlers.DownloadDeviceLocationDetail(dbConn))

	// Owner Details
	protected.GET("/owner-details", handlers.GetOwnerDetails(dbConn))
	protected.POST("/owner-details", handlers.CreateNewOwnerDetails(dbConn))
	protected.PATCH("/owner-details/:id", handlers.UpdateDeviceAMCOwnerDetail(dbConn))
	protected.DELETE("/owner-details/:id", handlers.DeleteDeviceAMCOwnerDetail(dbConn))
	protected.GET("/owner-details/pdf", handlers.DownloadDeviceAMCOwnerDetailPDF(dbConn))
	protected.GET("/owner-details/excel", handlers.DownloadDeviceAMCOwnerDetail(dbConn))

	// Power Details
	protected.GET("/power-details", handlers.GetPowerDetails(dbConn))
	protected.POST("/power-details", handlers.CreateNewPowerDetails(dbConn))
	protected.PATCH("/power-details/:id", handlers.UpdateDevicePowerDetail(dbConn))
	protected.DELETE("/power-details/:id", handlers.DeleteDevicePowerDetail(dbConn))
	protected.GET("/power-details/pdf", handlers.DownloadDevicePowerDetailPDF(dbConn))
	protected.GET("/power-details/excel", handlers.DownloadDevicePowerDetail(dbConn))

	// Fiber Details
	protected.GET("/fiber-details", handlers.GetFiberDetails(dbConn))
	protected.GET("/fiber-details/:id", handlers.GetFiberDetailByID(dbConn))
	protected.POST("/fiber-details", handlers.CreateNewFiberDetails(dbConn))
	protected.PATCH("/fiber-details/:id", handlers.UpdateDeviceEthernetFiberDetail(dbConn))
	protected.DELETE("/fiber-details/:id", handlers.DeleteDeviceEthernetFiberDetail(dbConn))
	protected.GET("/fiber-details/pdf", handlers.DownloadDeviceEthernetFiberDetailPDF(dbConn))
	protected.GET("/fiber-details/excel", handlers.DownloadDeviceEthernetFiberDetail(dbConn))
}
