package repository

import (
	"time"

	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
)

// func (r *Repository) GetRequests() (*[]ds.Request, error) {
// 	var requests []ds.Request
// 	res := r.db.Where("status != ?", "deleted").Find(&requests)
// 	return &requests, res.Error
// }

// func (r *Repository) DeleteRequest(id int) error {
// 	err := r.db.Exec("UPDATE requests SET status ='deleted' WHERE id = ?", id).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (r *Repository) UpdateRequestStatus(id int, status string) error {
	err := r.db.Exec("UPDATE requests SET status = $1 WHERE id = $2", status, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetRequestByIDAndStatus(id int, status string) (*ds.Request, error) {
	request := &ds.Request{}
	err := r.db.Table("requests").Where("deleted_at IS NULL").Where("creator_id = ? AND status = ?", id, status).First(&request).Error
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

func (r *Repository) GetRequestsForAdminWithFilters(minData time.Time, maxData time.Time, status string) ([]*ds.Request, error) {
	var requests []*ds.Request
	if status == "all" {
		res := r.db.Where("deleted_at IS NULL").Where("status != 'draft' AND status != 'deleted'").Where("formation_date <= ?", maxData).Where("formation_date >= ?", minData).Find(&requests)
		return requests, res.Error
	}

	res := r.db.Where("deleted_at IS NULL").Where("status != 'deleted' AND status = ?", status).Where("formation_date <= ?", maxData).Where("formation_date >= ?", minData).Find(&requests)
	return requests, res.Error
}

func (r *Repository) GetRequestsForAdminWithFiltersAndUser(minData time.Time, maxData time.Time, status string, id uint64) ([]*ds.Request, error) {
	var requests []*ds.Request
	if status == "all" {
		res := r.db.Where("deleted_at IS NULL").Where("creator_id = ?", id).Where("status != 'draft' AND status != 'deleted'").Where("formation_date <= ?", maxData).Where("formation_date >= ?", minData).Find(&requests)
		return requests, res.Error
	}

	res := r.db.Where("deleted_at IS NULL").Where("creator_id = ?", id).Where("status != 'deleted' AND status = ?", status).Where("formation_date <= ?", maxData).Where("formation_date >= ?", minData).Find(&requests)
	return requests, res.Error
}

func (r *Repository) GetRequests(id int64) ([]*ds.Request, error) {
	var requests []*ds.Request

	res := r.db.Where("deleted_at IS NULL").Where("creator_id = ?", id).Where("status != 'draft' AND status != 'deleted'").Find(&requests)
	return requests, res.Error
}

func (r *Repository) GetRequestItems(id int64) ([]*ds.Item, error) {
	var items []*ds.Item

	res := r.db.Table("items_requests").Select("items.*").Joins("JOIN items ON items.id = items_requests.item_id").Where("items_requests.deleted_at IS NULL").Where("items_requests.request_id = ?", id).Scan(&items)
	return items, res.Error
}

func (r *Repository) GetRequestByID(id int64) (*ds.Request, error) {
	request := &ds.Request{}
	err := r.db.First(request, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (r *Repository) PutRequestStatus(id int64, status string) error {
	var request ds.Request
	if err := r.db.Where("id = ?", id).First(&request).Error; err != nil {
		return err
	}

	request.Status = status
	if request.Status == "completed" {
		location, _ := time.LoadLocation("Europe/Moscow")
		currentDate := time.Now().In(location).Truncate(24 * time.Hour)
		request.CompletionDate = currentDate
	}
	if err := r.db.Save(&request).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) ConfirmRequest(id int64, status string) error {
	var request ds.Request
	if err := r.db.Where("id = ?", id).First(&request).Error; err != nil {
		return err
	}

	request.Status = status

	location, _ := time.LoadLocation("Europe/Moscow")
	currentDate := time.Now().In(location).Truncate(24 * time.Hour)

	request.CreatedAt = time.Now()
	request.CreationDate = currentDate
	request.FormationDate = currentDate

	if err := r.db.Save(&request).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteRequest(id int64) error {
	err := r.db.Model(&ds.Request{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":     "deleted",
		"deleted_at": r.db.NowFunc(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteDraftRequestItem(requestID int64, itemID int64) error {
	var itemRequest ds.ItemsRequest
	if err := r.db.Where("request_id = ? AND item_id = ?", requestID, itemID).Order("created_at desc").First(&itemRequest).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&itemRequest).Error; err != nil {
		return err
	}

	return nil
}
