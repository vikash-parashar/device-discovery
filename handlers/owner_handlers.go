package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/vikash-parashar/asset-locator/db"
	"github.com/vikash-parashar/asset-locator/logger"
	"github.com/vikash-parashar/asset-locator/models"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
)

// GetOwnerDetails handles the GET request to retrieve owner details.
func GetOwnerDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Fetching owner details from the database")
		data, err := db.GetAllDeviceAMCOwnerDetail()
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
		logger.InfoLogger.Println("Owner details fetched successfully.")
		c.HTML(http.StatusOK, "owner_details.html", data)
	}
}

func CreateNewOwnerDetails(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.DeviceAMCOwnerDetail

		logger.InfoLogger.Println("Creating new owner details")

		// Retrieve form data
		idStr := c.PostForm("id")
		serialNumber := c.PostForm("serial_number")
		deviceMakeModel := c.PostForm("device_make_model")
		model := c.PostForm("model")
		poNumber := c.PostForm("po_number")
		poOrderDateStr := c.PostForm("po_order_date")
		eoslDateStr := c.PostForm("eosl_date")
		amcStartDateStr := c.PostForm("amc_start_date")
		amcEndDateStr := c.PostForm("amc_end_date")
		deviceOwner := c.PostForm("device_owner")

		// Parse the date strings into time.Time values
		poOrderDate, _ := time.Parse("2006-01-02", poOrderDateStr)
		eoslDate, _ := time.Parse("2006-01-02", eoslDateStr)
		amcStartDate, _ := time.Parse("2006-01-02", amcStartDateStr)
		amcEndDate, _ := time.Parse("2006-01-02", amcEndDateStr)

		// Convert ID to int
		id, _ := strconv.Atoi(idStr)

		// Assign the values to the DeviceAMCOwnerDetail struct
		data = models.DeviceAMCOwnerDetail{
			Id:              id,
			SerialNumber:    serialNumber,
			DeviceMakeModel: deviceMakeModel,
			Model:           model,
			PONumber:        poNumber,
			POOrderDate:     poOrderDate,
			EOSLDate:        eoslDate,
			AMCStartDate:    amcStartDate,
			AMCEndDate:      amcEndDate,
			DeviceOwner:     deviceOwner,
		}

		if err := db.CreateDeviceAMCOwnerDetail(&data); err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
		logger.InfoLogger.Println("New owner details created successfully.")
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Entry Added Successfully"},
		)
	}
}

// UpdateDeviceAMCOwnerDetail updates a DeviceAMCOwnerDetail record based on its ID.
func UpdateDeviceAMCOwnerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Updating Device AMC Owner Detail")

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			logger.ErrorLogger.Println("Invalid ID:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		type RequestData struct {
			SerialNumber    string `json:"serial_number"`
			DeviceMakeModel string `json:"device_make_model"`
			Model           string `json:"model"`
			PONumber        string `json:"po_number"`
			POOrderDate     string `json:"po_order_date"`
			EOSLDate        string `json:"eosl_date"`
			AMCStartDate    string `json:"amc_start_date"`
			AMCEndDate      string `json:"amc_end_date"`
			DeviceOwner     string `json:"device_owner"`
		}

		var requestData RequestData
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		// Define the layout that matches the date format
		layout := "2006-01-02"

		// Parse the date string into a time.Time value
		pD, err := time.Parse(layout, requestData.POOrderDate)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Parse the date string into a time.Time value
		eD, err := time.Parse(layout, requestData.AMCEndDate)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Parse the date string into a time.Time value
		aD, err := time.Parse(layout, requestData.AMCStartDate)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Parse the date string into a time.Time value
		aDD, err := time.Parse(layout, requestData.AMCEndDate)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		updatedData := &models.DeviceAMCOwnerDetail{
			Id:              id,
			SerialNumber:    requestData.SerialNumber,
			DeviceMakeModel: requestData.DeviceMakeModel,
			Model:           requestData.Model,
			PONumber:        requestData.PONumber,
			POOrderDate:     pD,
			EOSLDate:        eD,
			AMCStartDate:    aD,
			AMCEndDate:      aDD,
			DeviceOwner:     requestData.DeviceOwner,
		}

		if err := db.UpdateDeviceAMCOwnerDetail(id, updatedData); err != nil {
			logger.ErrorLogger.Println("Failed to update Device AMC Owner Detail:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Device AMC Owner Detail"})
			return
		}

		logger.InfoLogger.Println("Device AMC Owner Detail updated successfully.")
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Device AMC Owner Detail updated successfully"})
	}
}

// DeleteDeviceAMCOwnerDetailHandler deletes a DeviceAMCOwnerDetail record based on its ID.
func DeleteDeviceAMCOwnerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Deleting Device AMC Owner Detail")

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			logger.ErrorLogger.Println("Invalid ID:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := db.DeleteDeviceAMCOwnerDetail(id); err != nil {
			logger.ErrorLogger.Println("Failed to delete Device AMC Owner Detail:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Device AMC Owner Detail"})
			return
		}

		logger.InfoLogger.Println("Device AMC Owner Detail deleted successfully.")
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Device AMC Owner Detail deleted successfully"})
	}
}

func DownloadDeviceAMCOwnerDetail(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Downloading Device AMC Owner Detail in Excel format")
		// Query the database for DeviceAMCOwnerDetail data
		rows, err := db.Query("SELECT * FROM device_amc_owner")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to query the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a new Excel file
		file := xlsx.NewFile()
		sheet, err := file.AddSheet("DeviceAMCOwnerDetails")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to create Excel sheet", http.StatusInternalServerError)
			return
		}

		// Add header row
		headerRow := sheet.AddRow()
		headerRow.AddCell().SetString("ID")
		headerRow.AddCell().SetString("Serial Number")
		headerRow.AddCell().SetString("Device Make Model")
		headerRow.AddCell().SetString("Model")
		headerRow.AddCell().SetString("PO Number")
		headerRow.AddCell().SetString("PO Order Date")
		headerRow.AddCell().SetString("EOSL Date")
		headerRow.AddCell().SetString("AMC Start Date")
		headerRow.AddCell().SetString("AMC End Date")
		headerRow.AddCell().SetString("Device Owner")

		// Add data rows from the database
		for rows.Next() {
			var device models.DeviceAMCOwnerDetail
			if err := rows.Scan(&device.Id, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.PONumber, &device.POOrderDate, &device.EOSLDate, &device.AMCStartDate, &device.AMCEndDate, &device.DeviceOwner); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}
			dataRow := sheet.AddRow()
			dataRow.AddCell().SetInt(device.Id)
			dataRow.AddCell().SetString(device.SerialNumber)
			dataRow.AddCell().SetString(device.DeviceMakeModel)
			dataRow.AddCell().SetString(device.Model)
			dataRow.AddCell().SetString(device.PONumber)
			dataRow.AddCell().SetDate(device.POOrderDate)
			dataRow.AddCell().SetDate(device.EOSLDate)
			dataRow.AddCell().SetDate(device.AMCStartDate)
			dataRow.AddCell().SetDate(device.AMCEndDate)
			dataRow.AddCell().SetString(device.DeviceOwner)
		}

		// Save the Excel file to the response
		c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=DeviceAMCOwnerDetails.xlsx")
		err = file.Write(c.Writer)
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to write Excel file to response", http.StatusInternalServerError)
		}
		logger.InfoLogger.Println("Device AMC Owner Detail downloaded successfully in Excel format.")
	}
}

func DownloadDeviceAMCOwnerDetailPDF(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.InfoLogger.Println("Downloading Device AMC Owner Detail in PDF format")
		// Query the database for DeviceAMCOwnerDetail data
		rows, err := db.Query("SELECT * FROM device_amc_owner")
		if err != nil {
			log.Fatal(err)
			http.Error(c.Writer, "Failed to query the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a new PDF document
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()

		// Set font and text size
		pdf.SetFont("Arial", "", 12)

		// Add table headers
		headers := []string{"ID", "Serial Number", "Device Make Model", "Model", "PO Number", "PO Order Date", "EOSL Date", "AMC Start Date", "AMC End Date", "Device Owner"}
		for _, header := range headers {
			pdf.CellFormat(40, 10, header, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)

		// Add data rows from the database
		for rows.Next() {
			var device models.DeviceAMCOwnerDetail
			if err := rows.Scan(&device.Id, &device.SerialNumber, &device.DeviceMakeModel, &device.Model, &device.PONumber, &device.POOrderDate, &device.EOSLDate, &device.AMCStartDate, &device.AMCEndDate, &device.DeviceOwner); err != nil {
				log.Fatal(err)
				http.Error(c.Writer, "Failed to scan database row", http.StatusInternalServerError)
				return
			}

			data := []string{
				fmt.Sprint(device.Id),
				device.SerialNumber,
				device.DeviceMakeModel,
				device.Model,
				device.PONumber,
				device.POOrderDate.Format("2006-01-02"),
				device.EOSLDate.Format("2006-01-02"),
				device.AMCStartDate.Format("2006-01-02"),
				device.AMCEndDate.Format("2006-01-02"),
				device.DeviceOwner,
			}

			for _, str := range data {
				pdf.CellFormat(40, 10, str, "1", 0, "C", false, 0, "")
			}
			pdf.Ln(-1)
		}

		// Create the PDF file
		pdf.Output(c.Writer)

		// Set response headers
		c.Header("Content-Type", "application/pdf")
		c.Header("Content-Disposition", "attachment; filename=DeviceAMCOwnerDetails.pdf")
		logger.InfoLogger.Println("Device AMC Owner Detail downloaded successfully in PDF format.")
	}
}
