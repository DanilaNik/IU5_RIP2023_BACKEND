package repository

import "github.com/DanilaNik/IU5_RIP2023/internal/ds"

func (r *Repository) GetRequests() (*[]ds.Request, error) {
	var requests []ds.Request
	res := r.db.Where("status != ?", "deleted").Find(&requests)
	return &requests, res.Error
}

func (r *Repository) GetRequestByID(id int) (*ds.Request, error) {
	request := &ds.Request{}
	err := r.db.First(request, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (r *Repository) DeleteRequest(id int) error {
	err := r.db.Exec("UPDATE requests SET status ='deleted' WHERE id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateRequestStatus(id int, status string) error {
	err := r.db.Exec("UPDATE requests SET status = $1 WHERE id = $2", status, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetRequestByIDAndStatus(id int, status string) (*ds.Request, error) {
	request := &ds.Request{}
	err := r.db.Table("requests").Where("creator_id = ? AND status = ?", id, status).First(&request).Error
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (r *Repository) PostRequest(request ds.Request) error {
	result := r.db.Create(&request)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) PostRequestItem(requestItem ds.ItemsRequest) error {
	result := r.db.Create(&requestItem)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
