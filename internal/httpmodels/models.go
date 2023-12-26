package httpmodels

type Item struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	Status   string `json:"status"`
	Quantity uint64 `json:"quantity"`
	Height   uint64 `json:"height"`
	Width    uint64 `json:"width"`
	Depth    uint64 `json:"depth"`
	Barcode  uint64 `json:"barcode"`
}
