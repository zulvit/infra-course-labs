SET search_path TO shop, public;

COPY shop.manufacturers(manufacturer_name)
FROM '/var/lib/postgresql/import/manufacturers.csv'
WITH (
    FORMAT CSV,
    HEADER TRUE,
    DELIMITER ',',
    ENCODING 'UTF8',
    NULL ''
);

COPY shop.categories(category_name)
FROM '/var/lib/postgresql/import/categories.csv'
WITH (FORMAT CSV, HEADER TRUE, DELIMITER ',', ENCODING 'UTF8', NULL '');

COPY shop.stores(store_name)
FROM '/var/lib/postgresql/import/stores.csv'
WITH (FORMAT CSV, HEADER TRUE, DELIMITER ',', ENCODING 'UTF8', NULL '');

COPY shop.customers(customer_fname, customer_lname)
FROM '/var/lib/postgresql/import/customers.csv'
WITH (FORMAT CSV, HEADER TRUE, DELIMITER ',', ENCODING 'UTF8', NULL '');

ANALYZE shop.manufacturers;
ANALYZE shop.categories;
ANALYZE shop.stores;
ANALYZE shop.customers;
