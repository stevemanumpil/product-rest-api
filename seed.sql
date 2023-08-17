CREATE TABLE product(  
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    price INT,
    description TEXT,
    quantity INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;

INSERT INTO product(name, price, description, quantity, created_at) VALUES
('Samsung S21', 20000000, 'Flagship mobile phone from samsung', 10, FROM_UNIXTIME( UNIX_TIMESTAMP('2010-04-30 14:53:27') + FLOOR(0 + (RAND() * 63072000)) )),
('Iphone 14', 25000000, 'High class phone type from apple', 15, FROM_UNIXTIME( UNIX_TIMESTAMP('2010-04-30 14:53:27') + FLOOR(0 + (RAND() * 63072000)) )),
('MacBook Pro', 34500000, 'MacBook Pro 2021 with mini-LED display may launch between September, November', 83, FROM_UNIXTIME( UNIX_TIMESTAMP('2010-04-30 14:53:27') + FLOOR(0 + (RAND() * 63072000)) )),
('Microsoft Surface Laptop 4', 14990000, 'Style and speed. Stand out on HD video calls backed by Studio Mics. Capture ideas on the vibrant touchscreen.', 32, FROM_UNIXTIME( UNIX_TIMESTAMP('2010-04-30 14:53:27') + FLOOR(0 + (RAND() * 63072000)) )),
('Skintific sunscreen', 120000, 'Sunscreen with spf 50 and PA ++++', 7, FROM_UNIXTIME( UNIX_TIMESTAMP('2010-04-30 14:53:27') + FLOOR(0 + (RAND() * 63072000)) )),
('Whimzees XXS 113pcs', 250000, 'Dog treats with natural ingredients', 5, FROM_UNIXTIME( UNIX_TIMESTAMP('2010-04-30 14:53:27') + FLOOR(0 + (RAND() * 63072000)) )),
('Perfume Oil', 75000, 'Mega Discount, Impression of Acqua Di Gio by GiorgioArmani concentrated attar perfume Oil', 17, FROM_UNIXTIME( UNIX_TIMESTAMP('2010-04-30 14:53:27') + FLOOR(0 + (RAND() * 63072000)) )),
('Plant Hanger For Home', 41000, 'Boho Decor Plant Hanger For Home Wall Decoration Macrame Wall Hanging Shelf', 27, FROM_UNIXTIME( UNIX_TIMESTAMP('2010-04-30 14:53:27') + FLOOR(0 + (RAND() * 63072000)) )),
('3D Embellishment Art Lamp', 200000, '3D led lamp sticker Wall sticker 3d wall art light on/off button  cell operated (included)', 54, FROM_UNIXTIME( UNIX_TIMESTAMP('2010-04-30 14:53:27') + FLOOR(0 + (RAND() * 63072000)) )),
('Key Holder', 30000, 'Attractive DesignMetallic materialFour key hooksReliable & DurablePremium Quality', 33, FROM_UNIXTIME( UNIX_TIMESTAMP('2010-04-30 14:53:27') + FLOOR(0 + (RAND() * 63072000)) ))