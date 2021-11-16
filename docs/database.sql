CREATE TABLE users (
    id CHAR(6) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    name VARCHAR(20) NOT NULL,
    mac_address BIT(48) NOT NULL
);
CREATE TABLE packet_logs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    transit_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    mac_address BIT(48) NOT NULL
);
INSERT INTO users (id, name, mac_address) VALUES ("19T325", "higuruchi", b'000111100000001100011000001000011001111110101000');
INSERT INTO packet_logs (mac_address) VALUES (b'000111100000001100011000001000011001111110101000');
INSERT INTO packet_logs (mac_address) VALUES (b'000111100000001100011000001000011001111110101001');