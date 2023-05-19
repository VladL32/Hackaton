package trash

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

type repo struct {
}

type PostRepository interface {
	Add(*Post)
	FindAll() ([]Post, error)
}

type Post struct {
	GasValue    int64       `json:"gasValue"`
	Humidity    float64     `json:"humidity"`
	Pressure    int64       `json:"pressure"`
	Temperature float64     `json:"temperature"`
	UserId      interface{} `json:"user_id"`
	WaterValue  int64       `json:"waterValue"`
	Time        string      `json:"time"`
}

const (
	projectId      string = "hackathon-1018f" //hackathon-1018f
	collectionName string = "users"
)

func NewPostRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *Post) (*Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("failed to create a firestore client! : %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"GasValue":    post.GasValue,
		"Humidity":    post.Humidity,
		"Pressure":    post.Pressure,
		"Temperature": post.Temperature,
		"UserId":      post.UserId,
		"WaterValue":  post.WaterValue,
		"Time":        post.Time,
	})
	if err != nil {
		log.Fatalf("An error appeared: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("failed to create a firestore client! : %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []Post
	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()
		if err != nil {
			break
		}
		post := Post{
			GasValue:    doc.Data()["gasValue"].(int64),
			Humidity:    doc.Data()["humidity"].(float64),
			Pressure:    doc.Data()["pressure"].(int64),
			Temperature: doc.Data()["temperature"].(float64),
			UserId:      doc.Data()["user_id"].(interface{}),
			WaterValue:  doc.Data()["waterValue"].(int64),
			Time:        doc.Data()["time"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (*repo) Add(post *Post) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("failed to create a firestore client! : %v", err)
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"gasValue":    post.GasValue,
		"humidity":    post.Humidity,
		"pressure":    post.Pressure,
		"temperature": post.Temperature,
		"user_id":     post.UserId,
		"waterValue":  post.WaterValue,
		"time":        post.Time,
	})
}
