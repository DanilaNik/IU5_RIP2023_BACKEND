package storage

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

type ItemsData struct {
	Items      []Item
	Filter     string
	SearchText string
	Status     string
}

func GetItems() *ItemsData {
	return &ItemsData{
		Items: []Item{
			{
				ID:       0,
				Name:     "Процессор Intel Core i7-9700K",
				ImageURL: "https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Intel-Core-i7-9700K.jpg",
				Status:   "storage",
				Quantity: 1,
				Height:   43,
				Width:    37,
				Depth:    37,
				Barcode:  1234567890123,
			},
			{
				ID:       1,
				Name:     "Видеокарта NVIDIA GeForce RTX 3080",
				ImageURL: "https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/NVIDIA-GeForce-RTX-3080.jpg",
				Status:   "shipment",
				Quantity: 2,
				Height:   53,
				Width:    133,
				Depth:    302,
				Barcode:  2345678901234,
			},
			{
				ID:       2,
				Name:     "Оперативная память Corsair Vengeance RGB Pro 16GB DDR4",
				ImageURL: "https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Corsair-Vengeance-RGB-Pro-16GB-DDR4.jpg",
				Status:   "recieve",
				Quantity: 3,
				Height:   51,
				Width:    133,
				Depth:    7,
				Barcode:  3456789012345,
			},
			{
				ID:       3,
				Name:     "Материнская плата MSI B450 TOMAHAWK MAX",
				ImageURL: "https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/MSI-B450-TOMAHAWK-MAX.png",
				Status:   "storage",
				Quantity: 4,
				Height:   30,
				Width:    244,
				Depth:    305,
				Barcode:  4567890123456,
			},
			{
				ID:       4,
				Name:     "Жесткий диск Western Digital Blue 2TB",
				ImageURL: "https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Western-Digital-Blue-2TB.jpg",
				Status:   "shipment",
				Quantity: 5,
				Height:   26,
				Width:    101,
				Depth:    147,
				Barcode:  5678901234567,
			},
			{
				ID:       5,
				Name:     "Блок питания Corsair RM750X, 750W",
				ImageURL: "https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Corsair-RM750X-750W.jpg",
				Status:   "recieve",
				Quantity: 6,
				Height:   86,
				Width:    150,
				Depth:    160,
				Barcode:  6789012345678,
			},
			{
				ID:       6,
				Name:     "SSD накопитель Samsung 970 EVO Plus 1TB",
				ImageURL: "https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Samsung-970-EVO-Plus-1TB.jpg",
				Status:   "storage",
				Quantity: 7,
				Height:   23,
				Width:    80,
				Depth:    228,
				Barcode:  7890123456789,
			},
			{
				ID:       7,
				Name:     "Корпус Phanteks Eclipse P400A",
				ImageURL: "https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Phanteks-Eclipse-P400A.jpg",
				Status:   "shipment",
				Quantity: 8,
				Height:   210,
				Width:    210,
				Depth:    465,
				Barcode:  8901234567890,
			},
			// {
			// 	ID:       8,
			// 	Name:     "Мышь Logitech G502 HERO",
			// 	ImageURL: "https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Logitech-G502-HERO.jpg",
			// 	Status:   "recieve",
			// 	Quantity: 9,
			// 	Height:   43,
			// 	Width:    75,
			// 	Depth:    132,
			// 	Barcode:  123456789012,
			// },
			// {
			// 	ID:       9,
			// 	Name:     "Клавиатура Logitech G915 Wireless",
			// 	ImageURL: "https://c.dns-shop.ru/thumb/st1/fit/500/500/cce3145868175caea49d904091598383/c66e11037852f4dc5bf8f1863eca01ca6c3974d483a9aece246ce340180082bb.jpg",
			// 	Status:   "recieve",
			// 	Quantity: 10,
			// 	Height:   22,
			// 	Width:    475,
			// 	Depth:    150,
			// 	Barcode:  9012345678901,
			// },
		},
	}
}
