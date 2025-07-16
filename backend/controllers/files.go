package controllers

import (
	"io"
	"net/http"
	"time"

	"github.com/blavejr/bowattServer/models"
	"github.com/blavejr/bowattServer/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FilesHandler struct {
	FileCollection *mongo.Collection
}

func NewFilesHandler(fileColl *mongo.Collection) *FilesHandler {
	return &FilesHandler{FileCollection: fileColl}
}

func (handler *FilesHandler) HandleFileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Open file stream
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	// Hash the file
	fileHash, err := utils.HashFileContent(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash file"})
		return
	}

	user := c.MustGet("user").(models.User)

	var existing models.FileMetadata
	err = handler.FileCollection.FindOne(c, bson.M{
		"file_hash": fileHash,
		"user_id":   user.ID,
	}).Decode(&existing)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "You already uploaded this file", "id": fileHash})
		return
	}

	src2, _ := file.Open()
	defer src2.Close()
	contentBytes, _ := io.ReadAll(src2)

	record := models.FileMetadata{
		ID:        primitive.NewObjectID(),
		UserId:    user.ID,
		Name:      file.Filename,
		FileHash:  fileHash,
		Content:   string(contentBytes),
		Timestamp: time.Now().Unix(),
	}

	_, err = handler.FileCollection.InsertOne(c, record)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded", "id": fileHash})
}

func (handler *FilesHandler) ListFilesForUser(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	// Find files for this user
	cursor, err := handler.FileCollection.Find(c, bson.M{"user_id": user.ID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}
	defer cursor.Close(c)

	var files []models.FileMetadata
	if err := cursor.All(c, &files); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode files"})
		return
	}

	c.JSON(http.StatusOK, files)
}

func (handler *FilesHandler) HandleMultipleFileUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid multipart form"})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
		return
	}

	// each file gets a status
	user := c.MustGet("user").(models.User)
	var uploaded []string
	var skipped []string
	var failed []string

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			failed = append(failed, file.Filename)
			continue
		}

		fileHash, err := utils.HashFileContent(src)
		src.Close()
		if err != nil {
			failed = append(failed, file.Filename)
			continue
		}

		var existing models.FileMetadata
		err = handler.FileCollection.FindOne(c, bson.M{
			"file_hash": fileHash,
			"user_id":   user.ID,
		}).Decode(&existing)

		if err == nil {
			skipped = append(skipped, file.Filename)
			continue
		}

		src2, _ := file.Open()
		contentBytes, _ := io.ReadAll(src2)
		src2.Close()

		record := models.FileMetadata{
			ID:        primitive.NewObjectID(),
			UserId:    user.ID,
			Name:      file.Filename,
			FileHash:  fileHash,
			Content:   string(contentBytes),
			Timestamp: time.Now().Unix(),
		}

		_, err = handler.FileCollection.InsertOne(c, record)
		if err != nil {
			failed = append(failed, file.Filename)
			continue
		}

		uploaded = append(uploaded, file.Filename)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Upload completed",
		"uploaded": uploaded,
		"skipped":  skipped,
		"failed":   failed,
	})
}
