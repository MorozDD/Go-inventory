package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inventory/api/handlers"
	"inventory/models"
	"net/http"
	"strconv"
)

func InitializeRoutes(r *gin.Engine) {

	r.POST("/part", func(c *gin.Context) {
		var part models.Part
		if err := c.ShouldBindJSON(&part); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := handlers.CreatePart(&part); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, part)
	})

	// Get a part by ID
	r.GET("/part/:id", func(c *gin.Context) {
		id := c.Param("id")
		idUint, err := strconv.ParseUint(id, 10, 0)
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}
		part, err := handlers.GetPartByID(idUint)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Part not found"})
			return
		}
		c.JSON(http.StatusOK, part)
	})

	// Get a part by ID and version
	r.GET("/part/:id/:version", func(c *gin.Context) {
		id := c.Param("id")
		version := c.Param("version")

		idUint, err := strconv.ParseUint(id, 10, 0)
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}

		versionInt, err := strconv.Atoi(version)
		if err != nil {
			fmt.Println("Invalid version:", err)
			return
		}

		part, err := handlers.GetPartByVersion(uint(idUint), versionInt)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Part not found"})
			return
		}

		c.JSON(http.StatusOK, part)
	})

	// Update a part by ID
	r.PUT("/part/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedPart models.Part
		if err := c.ShouldBindJSON(&updatedPart); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		idUint, err := strconv.ParseUint(id, 10, 0)
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}

		existingPart, err := handlers.GetPartByID(idUint)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Part not found"})
			return
		}

		existingPart.Name = updatedPart.Name
		existingPart.SKU = updatedPart.SKU
		existingPart.Description = updatedPart.Description
		existingPart.Price = updatedPart.Price
		existingPart.Location = updatedPart.Location
		existingPart.ShipmentPackaging = updatedPart.ShipmentPackaging
		existingPart.Metadata = updatedPart.Metadata
		existingPart.Attributes = updatedPart.Attributes

		if err := handlers.UpdatePart(existingPart); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, existingPart)
	})

	// Delete a part by ID
	r.DELETE("/part/:id", func(c *gin.Context) {
		id := c.Param("id")

		idUint, err := strconv.ParseUint(id, 10, 0)
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}
		existingPart, err := handlers.GetPartByID(idUint)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Part not found"})
			return
		}

		if err := handlers.DeletePart(existingPart); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Part deleted successfully"})
	})
}
