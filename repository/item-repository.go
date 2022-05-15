package repository

import (
  "errors"
  "main/input"
  "main/models"
  "main/db"
)

type ItemRepository struct {
}

func (c *ItemRepository) Create(body input.ItemInput) (models.Item, error) {

  item := models.Item{Title: body.Title, Description: body.Description, Count: *body.Count };

	db.GetDB().Create(&item);

  return item, nil;
}

func (c *ItemRepository) Edit(key string, body input.ItemInput) (*models.Item, error) {
	var item models.Item;

  result := db.GetDB().First(&item, "id = ?", key);

  if result.Error != nil {
		return nil, errors.New("Not Found.");
	}

  item.Title = body.Title;
  item.Description = body.Description;
  item.Count = *body.Count;

  db.GetDB().Save(&item);

  return &item, nil;
}

func (c *ItemRepository) GetAll(deleted bool) []models.Item {
  var items []models.Item;

	db.GetDB().Where("deleted = ?", deleted).Find(&items);

  return items;
}

func (c *ItemRepository) Delete(key string, comments string) (error) {
  var item models.Item;

  result := db.GetDB().Where("deleted = ?", false).First(&item, "id = ?", key);

  if result.Error != nil {
		return errors.New("Not Found.");
	}

  item.Deleted = true;
  item.DeletionComments = comments;

  db.GetDB().Save(item);
  
  return nil;
}

func (c *ItemRepository) UnDelete(key string) (*models.Item, error) {
  var item models.Item;

  result := db.GetDB().Where("deleted = ?", true).First(&item, "id = ?", key);

  if result.Error != nil {
		return nil, errors.New("Not Found.");
	}

  item.Deleted = false;
  item.DeletionComments = "";

  db.GetDB().Save(item);
  
  return &item, nil;
}
