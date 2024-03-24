-- Active: 1701521287160@@127.0.0.1@5432@flick_tickets

CREATE Table users (
    id BIGINT PRIMARY key,
    user_name VARCHAR(255),
    password VARCHAR(512),
    age INT ,
    avatar_url VARCHAR(255),
    address VARCHAR(255),
    role int,
    is_active INT ,
    expired_time int,
    created_at int,
    updated_at int
);
CREATE Table customers (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(20),
    OTP BIGINT
);

CREATE TABLE orders (
    id BIGINT PRIMARY KEY,
    ticket_id BIGINT,
    showtime int,
    release_date int,
    email VARCHAR(255),
    otp BIGINT,
    decription VARCHAR(2014),
    status int,
    price DECIMAL(10, 2),
    file VARCHAR(1024),
    created_at int,
    updated_at int
);

CREATE TABLE tickets (
    id BIGINT PRIMARY KEY,
    price DECIMAL(10, 2),
    quantity INT,
    decription VARCHAR(2014),
    sale int,
    showtime int,
    release_date int
    created_at int,
    updated_at int
);

COMMENT ON COLUMN tickets.showtime IS 'Thời gian chiếu phim (integer)';

-- Thêm comment cho cột release_date
COMMENT ON COLUMN tickets.release_date IS 'Ngày công chiếu (integer)';

CREATE Table file_storages(
    id BIGINT PRIMARY KEY,
    ticket_id  BIGINT,
    url VARCHAR(255),
    created_at int
);

CREATE TABLE payments(
    id BIGINT PRIMARY KEY,
    amount DECIMAL(10, 2),
    order_id BIGINT,
    transaction_id BIGINT,
);

CREATE Table transactions (
    id BIGINT PRIMARY KEY,
    transaction_date INT,
    status int
);

INSERT INTO "orders" ("ticket_id","showtime","release_date","email","otp","description","status","price","file","seats","sale","created_at","updated_at","id") VALUES 
(3744462,121213,0,'thuynguyen151387@gmail.com',0,'ves vipo',0,120000,'',10,0,1711287384,1711287654,3246993) RETURNING "id"

