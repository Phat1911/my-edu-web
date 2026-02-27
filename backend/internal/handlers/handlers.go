package handlers

import (
	"context"
	"edu-web-backend/internal/models"
	"edu-web-backend/internal/repository"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *repository.DB
}

func NewHandler(db *repository.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetVideos(c *gin.Context) {
	videos, err := h.db.GetAllVideos(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if videos == nil {
		videos = []models.Video{}
	}
	c.JSON(http.StatusOK, gin.H{"data": videos, "total": len(videos)})
}

func (h *Handler) GetAudios(c *gin.Context) {
	audios, err := h.db.GetAllAudios(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if audios == nil {
		audios = []models.Audio{}
	}
	c.JSON(http.StatusOK, gin.H{"data": audios, "total": len(audios)})
}

func (h *Handler) GetQRCodes(c *gin.Context) {
	qrs, err := h.db.GetAllQRCodes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if qrs == nil {
		qrs = []models.QRCode{}
	}
	c.JSON(http.StatusOK, gin.H{"data": qrs, "total": len(qrs)})
}

func (h *Handler) GenerateQR(c *gin.Context) {
	var req struct {
		Label     string `json:"label" binding:"required"`
		TargetURL string `json:"target_url" binding:"required"`
		Type      string `json:"type"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Type == "" {
		req.Type = "general"
	}

	qrData := generateQRBase64(req.TargetURL)
	if qrData == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "QR generation failed"})
		return
	}

	qr, err := h.db.SaveQRCode(c.Request.Context(), req.Label, req.TargetURL, req.Type, qrData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": qr})
}

func (h *Handler) GetChatHistory(c *gin.Context) {
	sessionID := c.Param("session_id")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id required"})
		return
	}
	msgs, err := h.db.GetChatHistory(c.Request.Context(), sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if msgs == nil {
		msgs = []models.ChatMessage{}
	}
	c.JSON(http.StatusOK, gin.H{"data": msgs})
}

func (h *Handler) SendChat(c *gin.Context) {
	var req struct {
		SessionID string `json:"session_id" binding:"required"`
		Message   string `json:"message" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.SessionID = strings.TrimSpace(req.SessionID)
	req.Message = strings.TrimSpace(req.Message)
	if len(req.SessionID) > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id too long (max 100)"})
		return
	}
	if len(req.Message) > 2000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "message too long (max 2000)"})
		return
	}
	if req.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "message cannot be empty"})
		return
	}

	if err := h.db.SaveChatMessage(c.Request.Context(), req.SessionID, "user", req.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := buildAIResponse(c.Request.Context(), req.Message, h.db)

	// Use context.WithoutCancel so a client disconnect does not orphan the assistant message.
	saveCtx := context.WithoutCancel(c.Request.Context())
	if err := h.db.SaveChatMessage(saveCtx, req.SessionID, "assistant", response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response":   response,
		"session_id": req.SessionID,
	})
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "EduWeb API is running"})
}
