-- Create a table for "Customer" domain
CREATE TABLE Customers
(
    Id   SERIAL PRIMARY KEY,
    Name VARCHAR(255) UNIQUE NOT NULL
);

-- Create a table for "Product" domain
CREATE TABLE Products
(
    Id    SERIAL PRIMARY KEY,
    Name  VARCHAR(255) UNIQUE NOT NULL,
    Price NUMERIC(15, 2)
);

-- Create a table for "Stocks" domain
CREATE TABLE Stocks
(
    Id        SERIAL PRIMARY KEY,
    ProductId INT UNIQUE NOT NULL,
    Quantity  INT UNIQUE NOT NULL
);

ALTER TABLE Stocks
    ADD CONSTRAINT fk_productId
        FOREIGN KEY (ProductId) REFERENCES products (Id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;

-- Create a table for "Store" domain
CREATE TABLE Stores
(
    Id   SERIAL PRIMARY KEY,
    Name VARCHAR(255) UNIQUE NOT NULL
);

-- Create a table for "StoreProduct" domain
CREATE TABLE StoreProducts
(
    StoreId   INT,
    ProductId INT,
    Price     DECIMAL(15, 2)
);

ALTER TABLE StoreProducts
    ADD CONSTRAINT pk_storeProducts
        PRIMARY KEY (StoreId, ProductId);

ALTER TABLE StoreProducts
    ADD CONSTRAINT fk_storeProducts_storeId
        FOREIGN KEY (StoreId) REFERENCES Stores (Id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;

ALTER TABLE StoreProducts
    ADD CONSTRAINT fk_storeProducts_productId
        FOREIGN KEY (ProductId) REFERENCES Stocks (Id)
            ON UPDATE CASCADE
            ON DELETE CASCADE;
