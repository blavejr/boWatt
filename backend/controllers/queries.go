package controllers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/blavejr/bowattServer/models"
	"github.com/blavejr/bowattServer/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QueriesHandler struct {
	QueriesCollection *mongo.Collection
	FilesCollection   *mongo.Collection
}

func NewQueriesHandler(queriesColl *mongo.Collection, filesCol *mongo.Collection) *QueriesHandler {
	return &QueriesHandler{
		QueriesCollection: queriesColl,
		FilesCollection:   filesCol,
	}
}

func (handler *QueriesHandler) HandleQuery(c *gin.Context) {
	slog.Info("Received query request")
	var req struct {
		FileHash string `json:"fileHash"` // changed from Filename to FileHash
		Query    string `json:"query"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Invalid request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "status": http.StatusBadRequest})
		return
	}

	// Look up the file content by hash
	var fileRecord models.FileMetadata
	user := c.MustGet("user").(models.User)
	slog.Info("Looking up file by hash", "fileHash", req.FileHash, "userId", user.ID)
	slog.Debug("Query request body", "body", req)
	err := handler.FilesCollection.FindOne(c, bson.M{
		"file_hash": req.FileHash,
		"user_id":   user.ID,
	}).Decode(&fileRecord)

	if err != nil {
		slog.Warn("File not found by hash", "fileHash", req.FileHash)
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found", "status": http.StatusNotFound})
		return
	}

	content := fileRecord.Content

	cacheKey := req.FileHash + ":" + req.Query
	if cached, found := utils.GetCachedQuery(cacheKey); found {
		slog.Info("Returning cached results", "cacheKey", cacheKey)
		c.JSON(http.StatusOK, gin.H{"results": cached, "status": http.StatusOK})
		return
	}

	snippets := utils.FuzzySearch(content, req.Query)
	utils.SetCachedQuery(cacheKey, snippets)

	// save query metadata into db for history
	record := models.QueryHistory{
		ID:        primitive.NewObjectID(),
		Query:     req.Query,
		UserId:    user.ID,
		FileHash:  req.FileHash,
		Timestamp: time.Now().Unix(),
	}

	_, err = handler.QueriesCollection.InsertOne(c, record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert failed", "status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": snippets, "status": http.StatusOK})
}

func (handler *QueriesHandler) QueryHistoryForUser(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	cursor, err := handler.QueriesCollection.Find(c, bson.M{"user_id": user.ID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch history", "status": http.StatusInternalServerError})
		return
	}
	defer cursor.Close(c)

	var queryHistory []models.QueryHistory
	if err := cursor.All(c, &queryHistory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode history", "status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, queryHistory)
}
