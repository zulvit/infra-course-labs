CREATE SCHEMA IF NOT EXISTS shop;
SET search_path TO shop, public;

BEGIN;

CREATE TABLE IF NOT EXISTS shop.manufacturers (
    manufacturer_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    manufacturer_name VARCHAR(100) NOT NULL,
    CONSTRAINT uq_manufacturer_name UNIQUE (manufacturer_name)
);

CREATE TABLE IF NOT EXISTS shop.categories (
    category_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    category_name VARCHAR(100) NOT NULL,
    CONSTRAINT uq_category_name UNIQUE (category_name)
);

CREATE TABLE IF NOT EXISTS shop.products (
    product_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    category_id BIGINT NOT NULL,
    manufacturer_id BIGINT NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    CONSTRAINT fk_product_category
        FOREIGN KEY (category_id)
        REFERENCES shop.categories(category_id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_product_manufacturer
        FOREIGN KEY (manufacturer_id)
        REFERENCES shop.manufacturers(manufacturer_id)
        ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_products_category
    ON shop.products(category_id);
CREATE INDEX IF NOT EXISTS idx_products_manufacturer
    ON shop.products(manufacturer_id);

CREATE TABLE IF NOT EXISTS shop.price_change (
    product_id BIGINT NOT NULL,
    price_change_ts TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    new_price NUMERIC(18, 2) NOT NULL,
    CHECK (new_price >= 0),
    CONSTRAINT pk_price_change PRIMARY KEY (product_id, price_change_ts),
    CONSTRAINT fk_price_change_product
        FOREIGN KEY (product_id)
        REFERENCES shop.products(product_id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_price_change_product_ts
    ON shop.price_change(product_id, price_change_ts DESC);

CREATE TABLE IF NOT EXISTS shop.stores (
    store_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    store_name VARCHAR(255) NOT NULL,
    CONSTRAINT uq_store_name UNIQUE (store_name)
);

CREATE TABLE IF NOT EXISTS shop.customers (
    customer_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    customer_fname VARCHAR(100) NOT NULL,
    customer_lname VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS shop.deliveries (
    store_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    delivery_date DATE NOT NULL,
    product_count INTEGER NOT NULL,
    CHECK (product_count > 0),
    CONSTRAINT pk_deliveries PRIMARY KEY (store_id, product_id, delivery_date),
    CONSTRAINT fk_delivery_store
        FOREIGN KEY (store_id)
        REFERENCES shop.stores(store_id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_delivery_product
        FOREIGN KEY (product_id)
        REFERENCES shop.products(product_id)
        ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS shop.purchases (
    purchase_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    store_id BIGINT NOT NULL,
    customer_id BIGINT NOT NULL,
    purchase_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_purchase_store
        FOREIGN KEY (store_id)
        REFERENCES shop.stores(store_id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_purchase_customer
        FOREIGN KEY (customer_id)
        REFERENCES shop.customers(customer_id)
        ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_purchases_customer
    ON shop.purchases(customer_id);
CREATE INDEX IF NOT EXISTS idx_purchases_date
    ON shop.purchases(purchase_date DESC);

CREATE TABLE IF NOT EXISTS shop.purchase_items (
    purchase_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    product_count BIGINT NOT NULL,
    product_price NUMERIC(18, 2) NOT NULL,
    CHECK (product_count > 0),
    CHECK (product_price >= 0),
    CONSTRAINT pk_purchase_items PRIMARY KEY (purchase_id, product_id),
    CONSTRAINT fk_item_purchase
        FOREIGN KEY (purchase_id)
        REFERENCES shop.purchases(purchase_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_item_product
        FOREIGN KEY (product_id)
        REFERENCES shop.products(product_id)
        ON DELETE RESTRICT
);

COMMIT;
