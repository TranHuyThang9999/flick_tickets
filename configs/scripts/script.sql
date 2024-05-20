-- Active: 1715291960470@@127.0.0.1@5432@flick_tickets@public

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
    movie_time integer,
    cinema_name varchar(255),
    created_at integer,
    updated_at integer,
    selected_seat varchar(1024),
    quantity integer,
    ticket_id bigint,
    original_number integer,
    price double precision,
    PRIMARY KEY(id)
);
--tức là 1 vé sẽ được chiếu ở nhiều phòng, 1 phòng sẽ có nhiều h chiếu'

CREATE TABLE cinemas(
    id bigint NOT NULL,
    cinema_name varchar(255),
    description varchar(255),
    conscious varchar(255),
    district varchar(255),
    commune varchar(255),
    address_details varchar(255),
    width_container integer,
    height_container integer,
    PRIMARY KEY(id)
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

CREATE TABLE carts(
    id BIGINT PRIMARY key,
    user_name VARCHAR(255),
    show_time_id BIGINT,
    seats_position VARCHAR(1024),
    price FLOAT,
    created_at int,
    updated_at int
);
SELECT *
FROM orders
WHERE cinema_name = 'Phong 01'
  AND movie_name = '8 viên ngọc rồng'
  AND created_at BETWEEN 1714842000 AND 1718384400
  AND status = 9;
 SELECT * FROM "show_times" WHERE cinema_name = 'Phong 01' and movie_time = 1716219576  ORDER BY "show_times"."id" LIMIT 1