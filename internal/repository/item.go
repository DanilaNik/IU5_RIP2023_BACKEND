package repository

import (
	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
)

func (r *Repository) GetItems(search string) ([]*ds.Item, error) {
	var items []*ds.Item
	if search != "" {
		res := r.db.Where("status = $1", "enabled").Where("Name  ILIKE $2", "%"+search+"%").Find(&items)
		return items, res.Error
	}
	res := r.db.Where("status = ?", "enabled").Find(&items)
	return items, res.Error
}

func (r *Repository) GetItemByID(id int) (*ds.Item, error) {
	item := &ds.Item{}
	err := r.db.First(item, "id=?", id).Error
	if err != nil {
		return &ds.Item{}, err
	}
	return item, nil
}

func (r *Repository) DeleteItem(id int) error {
	err := r.db.Exec("UPDATE items SET status ='deleted' WHERE id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
