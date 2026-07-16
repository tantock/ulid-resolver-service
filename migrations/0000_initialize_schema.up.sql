CREATE TABLE product_code_type(
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    display_name TEXT NOT NULL,
    CONSTRAINT product_code_type_display_name_unique
        UNIQUE (display_name)
);

CREATE TABLE product (
    id TEXT PRIMARY KEY,
    product_code TEXT NOT NULL,
    product_code_type_id int NOT NULL REFERENCES product_code_type(id),
    CONSTRAINT product_id_ulid_check
        CHECK (id ~ '^[0-9A-HJKMNP-TV-Z]{26}$'),
    CONSTRAINT product_product_code_unique
        UNIQUE (product_code)
);

CREATE INDEX idx_product_product_code_type_id
ON product(product_code_type_id);