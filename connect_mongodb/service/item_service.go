package service

import (
	"connect_mongodb_gorm/database"
	"connect_mongodb_gorm/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type ItemService interface {
	FindAll() ([]model.GroceryItem, error)
	FindById(id string) (*model.GroceryItem, error)
	Create(item *model.GroceryItem) error
	Update(id string, item *model.GroceryItem) error
	Delete(id string) error
}
type ItemServiceImpl struct{}

func (i ItemServiceImpl) FindAll() ([]model.GroceryItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.DB.Collection("GroceryItem").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []model.GroceryItem
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (i ItemServiceImpl) FindById(id string) (*model.GroceryItem, error) {
	//TODO implement me
	panic("implement me")
}

func (i ItemServiceImpl) Create(item *model.GroceryItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := database.DB.Collection("GroceryItem").InsertOne(ctx, item)
	if err != nil {
		return err
	}
	return nil
}

func (i ItemServiceImpl) Update(id string, item *model.GroceryItem) error {
	//TODO implement me
	panic("implement me")
}

func (i ItemServiceImpl) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewItemService() ItemService {
	return &ItemServiceImpl{}

}
