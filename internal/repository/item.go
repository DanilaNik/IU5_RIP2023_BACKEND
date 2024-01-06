package repository

import (
	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
)

func (r *Repository) GetItems(search string, host string) ([]*ds.Item, error) {
	var items []*ds.Item
	if search != "" {
		res := r.db.Where("status = $1", "enabled").Where("Name  ILIKE $2", "%"+search+"%").Find(&items)
		for _, item := range items {
			url := item.ImageURL
			item.ImageURL = "http://" + host + url
		}
		return items, res.Error
	}
	res := r.db.Where("status = ?", "enabled").Find(&items)
	for _, item := range items {
		url := item.ImageURL
		item.ImageURL = "http://" + host + url
	}
	return items, res.Error
}

func (r *Repository) GetItemByID(id int, host string) (*ds.Item, error) {
	item := &ds.Item{}
	err := r.db.First(item, "id=?", id).Error
	if err != nil {
		return &ds.Item{}, err
	}
	url := item.ImageURL
	item.ImageURL = "http://" + host + url
	return item, nil
}

func (r *Repository) DeleteItem(id int) error {
	err := r.db.Model(&ds.Item{}).Where("id = ?", id).Update("status", "deleted").Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) PostItem(item ds.Item) error {
	result := r.db.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) PutItem(item ds.Item, id int64) error {
	result := r.db.Where("id = ?", id).Updates(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
