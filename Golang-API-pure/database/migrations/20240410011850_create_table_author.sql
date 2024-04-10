-- +goose Up
CREATE TABLE authors ( 
		id SERIAL PRIMARY KEY, 
		Age Int NOT NULL, 
		FirstName VARCHAR(255) NOT NULL, 
		LastName VARCHAR(255) NOT NULL,
		Country VARCHAR(255) NOT NULL, 
		Description VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
	);
	
-- +goose Down
DROP TABLE authors;
