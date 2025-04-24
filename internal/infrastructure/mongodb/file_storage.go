package mongodb

import (
	"io"
	"mime/multipart"
	"prk/internal/domain/document"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FSRepository struct {
	database *mongo.Database
}

func NewFSRepository(db *mongo.Database) document.FileStorage {
	return &FSRepository{database: db}
}

func (rep *FSRepository) UploadFile(file multipart.File, filename string) (string, error) {
	bucket, err := gridfs.NewBucket(rep.database)
	if err != nil {
		return "", err
	}

	uploadStream, err := bucket.OpenUploadStream(
		filename,
		options.GridFSUpload().
			SetMetadata(bson.M{"uploaded_at": time.Now()}),
	)
	if err != nil {
		return "", err
	}
	defer uploadStream.Close()
	_, err = io.Copy(uploadStream, file)
	if err != nil {
		return "", err
	}
	fileID := uploadStream.FileID.(string)

	return fileID, nil
}
