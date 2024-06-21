-- Migration 0001_create_users_table.up.sql
CREATE TABLE public.usuarios (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Migration 0001_create_users_table.down.sql 
DROP TABLE public.usuarios;

