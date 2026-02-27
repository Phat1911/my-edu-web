package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendDirectMessage(c *gin.Context) {
	senderID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req struct {
		ReceiverID int    `json:"receiver_id"`
		Content    string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Content = strings.TrimSpace(req.Content)
	if req.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "content cannot be empty"})
		return
	}
	if req.ReceiverID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid receiver_id"})
		return
	}
	if req.ReceiverID == senderID.(int) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot send message to yourself"})
		return
	}

	receiver, err := h.db.GetUserByID(c.Request.Context(), req.ReceiverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	if receiver == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "receiver not found"})
		return
	}

	msg, err := h.db.SaveMessage(c.Request.Context(), senderID.(int), req.ReceiverID, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save message"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": msg})
}

func (h *Handler) GetDirectMessages(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	otherUserIDStr := c.Param("other_user_id")
	otherUserID, err := strconv.Atoi(otherUserIDStr)
	if err != nil || otherUserID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid other_user_id"})
		return
	}

	msgs, err := h.db.GetConversation(c.Request.Context(), userID.(int), otherUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": msgs})
}

func (h *Handler) GetUsers(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	users, err := h.db.GetUserList(c.Request.Context(), userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
