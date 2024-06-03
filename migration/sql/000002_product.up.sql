CREATE TABLE product(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    category VARCHAR NOT NULL,
    price NUMERIC NOT NULL,
    stock INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO product (name, category, price, stock) VALUES
('Indomie Mi Goreng', 'Instant Noodles', 3000, 500),
('Luwak White Coffee', 'Beverage', 10000, 300),
('SilverQueen Chocolate', 'Snacks', 15000, 200),
('Sari Roti', 'Bakery', 8000, 400),
('Amidis Bottled Water', 'Beverage', 5000, 1000),
('Teh Botol Sosro', 'Beverage', 6000, 600),
('Roma Marie Biscuits', 'Snacks', 12000, 250),
('BonCabe Level 15', 'Condiments', 6000, 220),
('Ultra Milk', 'Dairy', 9000, 450),
('Pocari Sweat', 'Beverage', 8000, 350),
('Belibis Sambal Asli', 'Condiments', 7000, 350),
('Chitato Potato Chips', 'Snacks', 15000, 250),
('Torabika Coffee', 'Beverage', 12000, 400),
('Nutrisari Orange Drink', 'Beverage', 5000, 500),
('Tolak Angin Herbal Drink', 'Health', 13000, 200),
('Rinso Detergent', 'Household', 18000, 150),
('Saori Saus Tiram', 'Condiments', 12000, 250),
('Garnier Facial Wash', 'Personal Care', 25000, 180),
('Wardah Lip Balm', 'Cosmetics', 28000, 160),
('Sedaap Mie Goreng', 'Instant Noodles', 2500, 550),
('Yakult Probiotic Drink', 'Health', 9000, 600),
('Good Day Coffee', 'Beverage', 7000, 450),
('Nabati Wafer', 'Snacks', 10000, 400),
('Ciptadent Toothpaste', 'Personal Care', 12000, 300),
('Shinzui Soap', 'Personal Care', 7000, 350),
('Gery Saluut Biscuit', 'Snacks', 6000, 500),
('Laurier Sanitary Pad', 'Personal Care', 18000, 200),
('Milo Chocolate Drink', 'Beverage', 17000, 280),
('Sido Muncul Tolak Linu', 'Health', 12000, 230),
('Hemaviton Stamina Plus', 'Health', 14000, 210);