--
-- PostgreSQL database dump
--

\restrict 38FU7ywF3SmIfaeOonvwamoForRSSmgIsKZeXeSBZm4ukR5L9zfl2F6X2PtTnnx

-- Dumped from database version 17.9
-- Dumped by pg_dump version 17.9

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: shop; Type: SCHEMA; Schema: -; Owner: shop_admin
--

CREATE SCHEMA shop;


ALTER SCHEMA shop OWNER TO shop_admin;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: shop; Owner: shop_admin
--

CREATE TABLE shop.categories (
    category_id bigint NOT NULL,
    category_name character varying(100) NOT NULL
);


ALTER TABLE shop.categories OWNER TO shop_admin;

--
-- Name: categories_category_id_seq; Type: SEQUENCE; Schema: shop; Owner: shop_admin
--

ALTER TABLE shop.categories ALTER COLUMN category_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME shop.categories_category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: customers; Type: TABLE; Schema: shop; Owner: shop_admin
--

CREATE TABLE shop.customers (
    customer_id bigint NOT NULL,
    customer_fname character varying(100) NOT NULL,
    customer_lname character varying(100) NOT NULL
);


ALTER TABLE shop.customers OWNER TO shop_admin;

--
-- Name: customers_customer_id_seq; Type: SEQUENCE; Schema: shop; Owner: shop_admin
--

ALTER TABLE shop.customers ALTER COLUMN customer_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME shop.customers_customer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: deliveries; Type: TABLE; Schema: shop; Owner: shop_admin
--

CREATE TABLE shop.deliveries (
    store_id bigint NOT NULL,
    product_id bigint NOT NULL,
    delivery_date date NOT NULL,
    product_count integer NOT NULL,
    CONSTRAINT deliveries_product_count_check CHECK ((product_count > 0))
);


ALTER TABLE shop.deliveries OWNER TO shop_admin;

--
-- Name: manufacturers; Type: TABLE; Schema: shop; Owner: shop_admin
--

CREATE TABLE shop.manufacturers (
    manufacturer_id bigint NOT NULL,
    manufacturer_name character varying(100) NOT NULL
);


ALTER TABLE shop.manufacturers OWNER TO shop_admin;

--
-- Name: manufacturers_manufacturer_id_seq; Type: SEQUENCE; Schema: shop; Owner: shop_admin
--

ALTER TABLE shop.manufacturers ALTER COLUMN manufacturer_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME shop.manufacturers_manufacturer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: price_change; Type: TABLE; Schema: shop; Owner: shop_admin
--

CREATE TABLE shop.price_change (
    product_id bigint NOT NULL,
    price_change_ts timestamp with time zone DEFAULT now() NOT NULL,
    new_price numeric(18,2) NOT NULL,
    CONSTRAINT price_change_new_price_check CHECK ((new_price >= (0)::numeric))
);


ALTER TABLE shop.price_change OWNER TO shop_admin;

--
-- Name: products; Type: TABLE; Schema: shop; Owner: shop_admin
--

CREATE TABLE shop.products (
    product_id bigint NOT NULL,
    category_id bigint NOT NULL,
    manufacturer_id bigint NOT NULL,
    product_name character varying(255) NOT NULL
);


ALTER TABLE shop.products OWNER TO shop_admin;

--
-- Name: products_product_id_seq; Type: SEQUENCE; Schema: shop; Owner: shop_admin
--

ALTER TABLE shop.products ALTER COLUMN product_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME shop.products_product_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: purchase_items; Type: TABLE; Schema: shop; Owner: shop_admin
--

CREATE TABLE shop.purchase_items (
    purchase_id bigint NOT NULL,
    product_id bigint NOT NULL,
    product_count bigint NOT NULL,
    product_price numeric(18,2) NOT NULL,
    CONSTRAINT purchase_items_product_count_check CHECK ((product_count > 0)),
    CONSTRAINT purchase_items_product_price_check CHECK ((product_price >= (0)::numeric))
);


ALTER TABLE shop.purchase_items OWNER TO shop_admin;

--
-- Name: purchases; Type: TABLE; Schema: shop; Owner: shop_admin
--

CREATE TABLE shop.purchases (
    purchase_id bigint NOT NULL,
    store_id bigint NOT NULL,
    customer_id bigint NOT NULL,
    purchase_date timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE shop.purchases OWNER TO shop_admin;

--
-- Name: purchases_purchase_id_seq; Type: SEQUENCE; Schema: shop; Owner: shop_admin
--

ALTER TABLE shop.purchases ALTER COLUMN purchase_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME shop.purchases_purchase_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: stores; Type: TABLE; Schema: shop; Owner: shop_admin
--

CREATE TABLE shop.stores (
    store_id bigint NOT NULL,
    store_name character varying(255) NOT NULL
);


ALTER TABLE shop.stores OWNER TO shop_admin;

--
-- Name: stores_store_id_seq; Type: SEQUENCE; Schema: shop; Owner: shop_admin
--

ALTER TABLE shop.stores ALTER COLUMN store_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME shop.stores_store_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: shop; Owner: shop_admin
--

COPY shop.categories (category_id, category_name) FROM stdin;
1	Смартфоны
2	Ноутбуки
3	Телевизоры
4	Планшеты
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: shop; Owner: shop_admin
--

COPY shop.customers (customer_id, customer_fname, customer_lname) FROM stdin;
1	Иван	Петров
2	Мария	Сидорова
3	Алексей	Козлов
\.


--
-- Data for Name: deliveries; Type: TABLE DATA; Schema: shop; Owner: shop_admin
--

COPY shop.deliveries (store_id, product_id, delivery_date, product_count) FROM stdin;
\.


--
-- Data for Name: manufacturers; Type: TABLE DATA; Schema: shop; Owner: shop_admin
--

COPY shop.manufacturers (manufacturer_id, manufacturer_name) FROM stdin;
1	Apple Inc.
2	Samsung Electronics
3	Sony Group Corporation
4	Lenovo Group
\.


--
-- Data for Name: price_change; Type: TABLE DATA; Schema: shop; Owner: shop_admin
--

COPY shop.price_change (product_id, price_change_ts, new_price) FROM stdin;
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: shop; Owner: shop_admin
--

COPY shop.products (product_id, category_id, manufacturer_id, product_name) FROM stdin;
\.


--
-- Data for Name: purchase_items; Type: TABLE DATA; Schema: shop; Owner: shop_admin
--

COPY shop.purchase_items (purchase_id, product_id, product_count, product_price) FROM stdin;
\.


--
-- Data for Name: purchases; Type: TABLE DATA; Schema: shop; Owner: shop_admin
--

COPY shop.purchases (purchase_id, store_id, customer_id, purchase_date) FROM stdin;
\.


--
-- Data for Name: stores; Type: TABLE DATA; Schema: shop; Owner: shop_admin
--

COPY shop.stores (store_id, store_name) FROM stdin;
1	Москва - Арбат
2	Санкт-Петербург - Невский
3	Казань - Баумана
\.


--
-- Name: categories_category_id_seq; Type: SEQUENCE SET; Schema: shop; Owner: shop_admin
--

SELECT pg_catalog.setval('shop.categories_category_id_seq', 4, true);


--
-- Name: customers_customer_id_seq; Type: SEQUENCE SET; Schema: shop; Owner: shop_admin
--

SELECT pg_catalog.setval('shop.customers_customer_id_seq', 3, true);


--
-- Name: manufacturers_manufacturer_id_seq; Type: SEQUENCE SET; Schema: shop; Owner: shop_admin
--

SELECT pg_catalog.setval('shop.manufacturers_manufacturer_id_seq', 4, true);


--
-- Name: products_product_id_seq; Type: SEQUENCE SET; Schema: shop; Owner: shop_admin
--

SELECT pg_catalog.setval('shop.products_product_id_seq', 1, false);


--
-- Name: purchases_purchase_id_seq; Type: SEQUENCE SET; Schema: shop; Owner: shop_admin
--

SELECT pg_catalog.setval('shop.purchases_purchase_id_seq', 1, false);


--
-- Name: stores_store_id_seq; Type: SEQUENCE SET; Schema: shop; Owner: shop_admin
--

SELECT pg_catalog.setval('shop.stores_store_id_seq', 3, true);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (category_id);


--
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (customer_id);


--
-- Name: manufacturers manufacturers_pkey; Type: CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.manufacturers
    ADD CONSTRAINT manufacturers_pkey PRIMARY KEY (manufacturer_id);


--
-- Name: deliveries pk_deliveries; Type: CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.deliveries
    ADD CONSTRAINT pk_deliveries PRIMARY KEY (store_id, product_id, delivery_date);


--
-- Name: price_change pk_price_change; Type: CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.price_change
    ADD CONSTRAINT pk_price_change PRIMARY KEY (product_id, price_change_ts);


--
-- Name: purchase_items pk_purchase_items; Type: CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.purchase_items
    ADD CONSTRAINT pk_purchase_items PRIMARY KEY (purchase_id, product_id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (product_id);


--
-- Name: purchases purchases_pkey; Type: CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.purchases
    ADD CONSTRAINT purchases_pkey PRIMARY KEY (purchase_id);


--
-- Name: stores stores_pkey; Type: CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.stores
    ADD CONSTRAINT stores_pkey PRIMARY KEY (store_id);


--
-- Name: categories uq_category_name; Type: CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.categories
    ADD CONSTRAINT uq_category_name UNIQUE (category_name);


--
-- Name: manufacturers uq_manufacturer_name; Type: CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.manufacturers
    ADD CONSTRAINT uq_manufacturer_name UNIQUE (manufacturer_name);


--
-- Name: stores uq_store_name; Type: CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.stores
    ADD CONSTRAINT uq_store_name UNIQUE (store_name);


--
-- Name: idx_price_change_product_ts; Type: INDEX; Schema: shop; Owner: shop_admin
--

CREATE INDEX idx_price_change_product_ts ON shop.price_change USING btree (product_id, price_change_ts DESC);


--
-- Name: idx_products_category; Type: INDEX; Schema: shop; Owner: shop_admin
--

CREATE INDEX idx_products_category ON shop.products USING btree (category_id);


--
-- Name: idx_products_manufacturer; Type: INDEX; Schema: shop; Owner: shop_admin
--

CREATE INDEX idx_products_manufacturer ON shop.products USING btree (manufacturer_id);


--
-- Name: idx_purchases_customer; Type: INDEX; Schema: shop; Owner: shop_admin
--

CREATE INDEX idx_purchases_customer ON shop.purchases USING btree (customer_id);


--
-- Name: idx_purchases_date; Type: INDEX; Schema: shop; Owner: shop_admin
--

CREATE INDEX idx_purchases_date ON shop.purchases USING btree (purchase_date DESC);


--
-- Name: deliveries fk_delivery_product; Type: FK CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.deliveries
    ADD CONSTRAINT fk_delivery_product FOREIGN KEY (product_id) REFERENCES shop.products(product_id) ON DELETE RESTRICT;


--
-- Name: deliveries fk_delivery_store; Type: FK CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.deliveries
    ADD CONSTRAINT fk_delivery_store FOREIGN KEY (store_id) REFERENCES shop.stores(store_id) ON DELETE RESTRICT;


--
-- Name: purchase_items fk_item_product; Type: FK CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.purchase_items
    ADD CONSTRAINT fk_item_product FOREIGN KEY (product_id) REFERENCES shop.products(product_id) ON DELETE RESTRICT;


--
-- Name: purchase_items fk_item_purchase; Type: FK CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.purchase_items
    ADD CONSTRAINT fk_item_purchase FOREIGN KEY (purchase_id) REFERENCES shop.purchases(purchase_id) ON DELETE CASCADE;


--
-- Name: price_change fk_price_change_product; Type: FK CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.price_change
    ADD CONSTRAINT fk_price_change_product FOREIGN KEY (product_id) REFERENCES shop.products(product_id) ON DELETE CASCADE;


--
-- Name: products fk_product_category; Type: FK CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.products
    ADD CONSTRAINT fk_product_category FOREIGN KEY (category_id) REFERENCES shop.categories(category_id) ON DELETE RESTRICT;


--
-- Name: products fk_product_manufacturer; Type: FK CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.products
    ADD CONSTRAINT fk_product_manufacturer FOREIGN KEY (manufacturer_id) REFERENCES shop.manufacturers(manufacturer_id) ON DELETE RESTRICT;


--
-- Name: purchases fk_purchase_customer; Type: FK CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.purchases
    ADD CONSTRAINT fk_purchase_customer FOREIGN KEY (customer_id) REFERENCES shop.customers(customer_id) ON DELETE RESTRICT;


--
-- Name: purchases fk_purchase_store; Type: FK CONSTRAINT; Schema: shop; Owner: shop_admin
--

ALTER TABLE ONLY shop.purchases
    ADD CONSTRAINT fk_purchase_store FOREIGN KEY (store_id) REFERENCES shop.stores(store_id) ON DELETE RESTRICT;


--
-- PostgreSQL database dump complete
--

\unrestrict 38FU7ywF3SmIfaeOonvwamoForRSSmgIsKZeXeSBZm4ukR5L9zfl2F6X2PtTnnx

