package controllers

import (
    "net/http"

    "GoGym/models"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type ActionController struct {
    DB *gorm.DB
}

func (ctrl *ActionController) CreateAction(c *gin.Context) {
    var action models.Action
    if err := c.ShouldBindJSON(&action); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ctrl.DB.Create(&action).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create action"})
        return
    }

    c.JSON(http.StatusCreated, action)
}

func (ctrl *ActionController) GetAction(c *gin.Context) {
    id := c.Param("id")
    var action models.Action

    if err := ctrl.DB.First(&action, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Action not found"})
        return
    }

    c.JSON(http.StatusOK, action)
}

func (ctrl *ActionController) GetAllActions(c *gin.Context) {
    var actions []models.Action
    if err := ctrl.DB.Find(&actions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve actions"})
        return
    }

    c.JSON(http.StatusOK, actions)
}

func (ctrl *ActionController) UpdateAction(c *gin.Context) {
    id := c.Param("id")
    var action models.Action

    if err := ctrl.DB.First(&action, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Action not found"})
        return
    }

    if err := c.ShouldBindJSON(&action); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ctrl.DB.Save(&action).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update action"})
        return
    }

    c.JSON(http.StatusOK, action)
}

func (ctrl *ActionController) DeleteAction(c *gin.Context) {
    id := c.Param("id")
    if err := ctrl.DB.Delete(&models.Action{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete action"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Action deleted successfully"})
}