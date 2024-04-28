-- Active: 1714154952322@@127.0.0.1@5432@flick_tickets@public

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
    use_name VARCHAR(255),
    address VARCHAR(255),
    age INT ,
    email VARCHAR(255),
    phone_number VARCHAR(20),
    OTP BIGINT,
    created_at int,
    updated_at int
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
CREATE TABLE show_times(
    id bigint NOT NULL,
    ticket_id bigint,
    movie_time int,
    cinema_name varchar(255),
    created_at integer,
    updated_at integer,
    PRIMARY KEY(id)
);
--tức là 1 vé sẽ được chiếu ở nhiều phòng, 1 phòng sẽ có nhiều h chiếu'

CREATE Table cinemas(
    id BIGINT PRIMARY KEY,
    cinema_name VARCHAR(255)
);

CREATE TABLE movie_theaters (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  address VARCHAR(255),
  city VARCHAR(255),
  state VARCHAR(255),
  country VARCHAR(255)
);


create Table movie_types (
    id BIGINT PRIMARY key,
    movie_type_name VARCHAR(1024)
)

select DISTINCT name from cities;

SELECT DISTINCT district from cities where name in('Thành phố Hà Nội','Tỉnh Hưng Yên');
