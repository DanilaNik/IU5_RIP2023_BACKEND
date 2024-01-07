package repository

import (
	"strconv"
	"time"

	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
)

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

func (r *Repository) GetUserRequests(id int) (*[]ds.Request, error) {

	return nil, nil
}

func (r *Repository) DeleteUser(id int) error {
	err := r.db.Exec("UPDATE users SET status ='deleted' WHERE id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CreateUser(user ds.User) error {
	err := r.db.Table("users").Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserByLogin(login string) (ds.User, error) {
	user := ds.User{}
	err := r.db.Table("users").Where(`"login" = ?`, login).Find(&user).Error
	if err != nil {
		return ds.User{}, err
	}
	return user, nil
}

func (r *Repository) SaveJWTToken(id uint, token string) error {
	expiration := time.Hour * 1

	idStr := strconv.FormatUint(uint64(id), 10)

	err := r.rd.Set(idStr, token, expiration).Err()
	r.rd.Del()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteJWTToken(id string) error {
	err := r.rd.Del(id).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetJWTToken(id string) error {
	_, err := r.rd.Get(id).Result()
	if err != nil {
		return err
	}

	return nil
}
