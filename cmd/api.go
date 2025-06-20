package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UsageLog struct {
	UserID    string `json:"user_id" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
	Photos    int    `json:"photos"`
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/logs", saveUsageLog)
	router.GET("/usage/:userId", getMonthlyUsage)
}

func saveUsageLog(c *gin.Context) {
	var input UsageLog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB 연결 실패: " + err.Error()})
		return
	}
	defer db.Close()

	if err := InsertUsageLog(db, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "로그 저장 실패: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "로그 저장 완료"})
}

func getMonthlyUsage(c *gin.Context) {
	userID := c.Param("userId")
	yearStr := c.Query("year")
	monthStr := c.Query("month")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "year 파라미터가 필요하며 숫자여야 합니다"})
		return
	}
	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "month 파라미터는 1~12 범위의 숫자여야 합니다"})
		return
	}

	db, err := InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB 연결 실패: " + err.Error()})
		return
	}
	defer db.Close()

	usage, err := GetMonthlyUsage(db, userID, year, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "사용량 조회 실패: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":      userID,
		"year":         year,
		"month":        month,
		"used_minutes": usage.TotalDurationMinutes,
		"photo_count":  usage.TotalPhotos,
	})
}
