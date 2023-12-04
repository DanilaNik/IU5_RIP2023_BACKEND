package repository

import "github.com/DanilaNik/IU5_RIP2023/internal/ds"

func (r *Repository) GetUsers() (*[]ds.User, error) {
	var items []ds.User
	res := r.db.Find(&items)
	return &items, res.Error
}

func (r *Repository) GetUserByID(id int) (*ds.User, error) {
	user := &ds.User{}
	err := r.db.First(user, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) DeleteUser(id int) error {
	err := r.db.Exec("UPDATE users SET status ='deleted' WHERE id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
