package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/toanmars/crud-generator/models"
	"github.com/toanmars/crud-generator/services"
)

type StaffHandler struct {
	service *services.StaffService
}

func NewStaffHandler(service *services.StaffService) *StaffHandler {
	return &StaffHandler{service: service}
}

func (h *StaffHandler) Create(c *gin.Context) {
	var model models.Staff
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&model); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model)
}

func (h *StaffHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	model, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model)
}

func (h *StaffHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var model models.Staff
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.ID = uint(id)
	if err := h.service.Update(&model); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model)
}

func (h *StaffHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
