INSERT INTO users (username, password, role) VALUES
                                                 (1, 'user1', 'login1', 'password1', 'user1', 'creator'),
                                                 ('user2', 'password2', 'creator'),
                                                 ('moderator1', 'password3', 'moderator'),
                                                 ('moderator2', 'password4', 'moderator'),
                                                 ('user3', 'password5', 'creator');

-- type User struct {
-- 	gorm.Model
-- 	ID       uint64 `json:"id" gorm:"primary_key"`
-- 	UserName string `json:"userName" gorm:"type:varchar(100);not null"`
-- 	Login    string `json:"login" gorm:"type:text;not null"`
-- 	Password string `json:"password" gorm:"type:varchar(100);not null"`
-- 	Email    string `json:"email" gorm:"unique;type:varchar(100);not null"`
-- 	Role     string `json:"role" gorm:"type:varchar(20);check:role IN ('Admin', 'Moderator');not null"`
-- 	ImageURL string `json:"image_url" gorm:"type:varchar(500);default:'https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/default-avatar.png'"`
-- }


INSERT INTO items (id, name, image_url, status, quantity, height, width, depth, barcode) VALUES
                                                                                             (1, 'Процессор Intel Core i7-9700K', 'https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Intel-Core-i7-9700K.jpg', 'enabled', 1, 43, 37, 37, 1234567890123),
                                                                                             (2, 'Видеокарта NVIDIA GeForce RTX 3080', 'https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Corsair-Vengeance-RGB-Pro-16GB-DDR4.jpg', 'enabled', 2, 53, 133, 302, 2345678901234),
                                                                                             (3, 'Оперативная память Corsair Vengeance RGB Pro 16GB DDR4', 'https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Corsair-Vengeance-RGB-Pro-16GB-DDR4.jpg', 'enabled', 3, 51, 133, 7, 3456789012345
                                                                                             (4, 'Материнская плата MSI B450 TOMAHAWK MAX', 'https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/MSI-B450-TOMAHAWK-MAX.png', 'enabled', 4, 30, 244, 305, 4567890123456),
                                                                                             (5, 'Жесткий диск Western Digital Blue 2TB', 'https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Western-Digital-Blue-2TB.jpg', 'enabled', 5, 26, 101, 147, 5678901234567),
                                                                                             (6, 'Блок питания Corsair RM750X, 750W', 'https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Corsair-RM750X-750W.jpg', 'enabled', 6, 86, 150 , 160, 6789012345678),
			                                                                                 (7, 'SSD накопитель Samsung 970 EVO Plus 1TB', 'https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Samsung-970-EVO-Plus-1TB.jpg', 'enabled', 23, 80, 228, 7890123456789),
                                                                                             (8, 'Корпус Phanteks Eclipse P400A', 'https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Phanteks-Eclipse-P400A.jpg', 'enabled', 210, 210, 265, 8901234567890);

INSERT INTO requests (status, creation_date, formation_date, completion_date, creator_id, moderator_id) VALUES
                                                                                                            ('draft', '2022-01-01 00:00:00', NULL, NULL, 1, 3),
                                                                                                            ('formed', '2022-01-02 00:00:00', '2022-01-03 00:00:00', NULL, 2, 3),
                                                                                                            ('completed', '2022-01-03 00:00:00', '2022-01-04 00:00:00', '2022-01-05 00:00:00', 1, 4),
                                                                                                            ('rejected', '2022-01-04 00:00:00', '2022-01-05 00:00:00', '2022-01-06 00:00:00', 2, 4),
                                                                                                            ('draft', '2022-01-05 00:00:00', NULL, NULL, 1, 3);

INSERT INTO request_item (request_id, item_id) VALUES
                                                   (1, 1),
                                                   (1, 2),
                                                   (2, 2),
                                                   (2, 3),
                                                   (3, 3),
                                                   (3, 4),
                                                   (4, 4),
                                                   (4, 5),
                                                   (5, 5),
                                                   (5, 1);
