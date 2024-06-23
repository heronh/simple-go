CREATE TABLE public.todos (
	id SERIAL PRIMARY KEY,
	description VARCHAR(255) NOT NULL,
	status boolean default false,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
