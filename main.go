package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"miniProject2/dto"
	"net/http"
)

func main() {
	// Koneksi ke basis data MySQL
	dsn := "mysql:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Op(dsn), &gorm.Config{})
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Inisialisasi router
	router := gin.Default()

	// Handler untuk endpoint register customer
	router.POST("/register", func(c *gin.Context) {
		// Ambil data dari body request
		var customer dto.Customer
		if err := c.ShouldBindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Lakukan validasi data customer

		// Simpan data customer ke basis data
		if err := db.Create(&customer).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register customer"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Customer registered successfully"})
	})

	// Handler untuk endpoint register admin
	router.POST("/register/admin", func(c *gin.Context) {
		// Ambil data dari body request
		var admin dto.Actor
		if err := c.ShouldBindJSON(&admin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Hash password admin
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		admin.Password = string(hashedPassword)

		// Simpan data admin ke basis data
		if err := db.Create(&admin).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register admin"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Admin registered successfully"})
	})

	// Handler untuk endpoint approve/reject admin registration
	router.POST("/admin/approval", func(c *gin.Context) {
		// Ambil data dari body request
		var approval dto.RegisterApproval
		if err := c.ShouldBindJSON(&approval); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Lakukan logika untuk approve/reject admin registration

		c.JSON(http.StatusOK, gin.H{"message": "Admin registration approved/rejected"})
	})

	// Handler untuk endpoint melihat approval request
	router.GET("/admin/approval", func(c *gin.Context) {
		// Lakukan logika untuk melihat approval request

		// Contoh response
		approvals := []dto.RegisterApproval{
			{ID: 1, AdminID: 1, SuperAdminID: 1, Status: "Pending"},
			{ID: 2, AdminID: 2, SuperAdminID: 1, Status: "Pending"},
		}

		c.JSON(http.StatusOK, gin.H{"approvals": approvals})
	})

	// Handler untuk endpoint login admin
	router.POST("/login", func(c *gin.Context) {
		// Ambil data dari body request
		var loginData struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Lakukan logika untuk login admin

		// Contoh response
		token := "example_token"

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// Handler untuk endpoint menghapus data customer
	router.DELETE("/customer/:id", func(c *gin.Context) {
		// Ambil ID customer dari path parameter
		id := c.Param("id")

		// Lakukan logika untuk menghapus data customer

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Customer with ID %s deleted", id)})
	})

	// Handler untuk endpoint menghapus data admin
	router.DELETE("/admin/:id", func(c *gin.Context) {
		// Ambil ID admin dari path parameter
		id := c.Param("id")

		// Lakukan logika untuk menghapus data admin

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Admin with ID %s deleted", id)})
	})

	// Handler untuk endpoint mengaktifkan/deaktifkan admin
	router.PUT("/admin/:id/activate", func(c *gin.Context) {
		// Ambil ID admin dari path parameter
		id := c.Param("id")

		// Lakukan logika untuk mengaktifkan/deaktifkan admin

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Admin with ID %s activated/deactivated", id)})
	})

	// Handler untuk endpoint mendapatkan data customer dengan parameter pencarian dan paginasi
	router.GET("/customers", func(c *gin.Context) {
		// Ambil query parameter pencarian
		// Ambil query parameter paginasi

		// Lakukan logika untuk mendapatkan data customer

		// Contoh response
		customers := []dto.Customer{
			{ID: 1, FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Avatar: "avatar1.jpg"},
			{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane.smith@example.com", Avatar: "avatar2.jpg"},
		}

		c.JSON(http.StatusOK, gin.H{"customers": customers})
	})

	// Handler untuk endpoint mendapatkan data admin dengan parameter pencarian dan paginasi
	router.GET("/admins", func(c *gin.Context) {

		// Contoh response
		admins := []dto.Actor{
			{ID: 1, Username: "admin1", Password: "hashed_password1", RoleID: 1},
			{ID: 2, Username: "admin2", Password: "hashed_password2", RoleID: 1},
		}

		c.JSON(http.StatusOK, gin.H{"admins": admins})
	})

	// Handler untuk endpoint melakukan sinkronisasi data customer dari external API
	router.GET("/customers/sync", func(c *gin.Context) {
		// Lakukan logika untuk sinkronisasi data customer

		c.JSON(http.StatusOK, gin.H{"message": "Customer data synchronized"})
	})

	// Jalankan server HTTP
	router.Run(":8080")

}
