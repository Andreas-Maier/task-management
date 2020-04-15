package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// Task - Model of a basic task
type Task struct {
	ID    primitive.ObjectID
	Title string
	Body  string
}
