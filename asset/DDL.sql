CREATE DATABASE cofee_shop;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE roles AS ENUM ('admin', 'employee');

CREATE TABLE users(
    user_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role roles NOT NULL
);

CREATE TABLE products(
    product_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);