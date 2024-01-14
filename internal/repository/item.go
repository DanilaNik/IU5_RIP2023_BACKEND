package repository

import (
	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
)

func (r *Repository) GetItems(search string, material string, host string) ([]*ds.Item, error) {
	var items []*ds.Item
	if search != "" {
		if material != "" {
			res := r.db.Where("deleted_at IS NULL").Where("material = $1", material).Where("status = $2", "enabled").Where("name  ILIKE $3", "%"+search+"%").Order("id ASC").Find(&items)
			for _, item := range items {
				url := item.ImageURL
				item.ImageURL = "http://" + host + url
			}
			return items, res.Error
		} else {
			res := r.db.Where("deleted_at IS NULL").Where("status = $1", "enabled").Where("name  ILIKE $2", "%"+search+"%").Order("id ASC").Find(&items)
			for _, item := range items {
				url := item.ImageURL
				item.ImageURL = "http://" + host + url
			}
			return items, res.Error
		}
	} else {
		if material != "" {
			res := r.db.Where("material = $1", material).Where("status = $2", "enabled").Find(&items)
			for _, item := range items {
				url := item.ImageURL
				item.ImageURL = "http://" + host + url
			}
			return items, res.Error
		} else {
			res := r.db.Where("status = ?", "enabled").Find(&items)
			for _, item := range items {
				url := item.ImageURL
				item.ImageURL = "http://" + host + url
			}
			return items, res.Error
		}
	}

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
	err := r.db.Model(&ds.Item{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":     "deleted",
		"deleted_at": r.db.NowFunc(),
	}).Error
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
