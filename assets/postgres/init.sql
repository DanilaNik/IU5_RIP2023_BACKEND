CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    login VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) CHECK (role IN ('creator', 'moderator'))
);

CREATE TABLE IF NOT EXISTS Item (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image_url VARCHAR(500) DEFAULT 'https://raw.githubusercontent.com/DanilaNik/IU5_RIP2023/lab1/resources/Intel-Core-i7-9700K.jpg',
    status INT NOT NULL,
    quantity BIGINT NOT NULL,
    height BIGINT NOT NULL,
    width BIGINT NOT NULL,
    depth BIGINT NOT NULL,
    barcode BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS Request  (
    id SERIAL PRIMARY KEY,
    status INT CHECK (status IN (1, 2, 3, 4, 5)),
    creation_date TIMESTAMP NOT NULL,
    formation_date TIMESTAMP,
    completion_date TIMESTAMP,
    creator_id INT REFERENCES users(id),
    moderator_id INT REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS Request_Item (
    request_id INT REFERENCES requests(id),
    item_id BIGINT REFERENCES items(id),
    PRIMARY KEY (request_id, item_id)
);
