CREATE TABLE authors(
	id VARCHAR(64) NOT NULL PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleated_at timestamp	
);

CREATE TABLE books (
	id VARCHAR(64) NOT NULL PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	author_id VARCHAR(64) NOT NULL REFERENCES authors (id),
	created_at timestamp,
	updated_at timestamp,
	deleated_at timestamp
);

CREATE TABLE categories (
	id VARCHAR(64) NOT NULL PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	parent_id VARCHAR(64),
	created_at timestamp,
	updated_at timestamp,
	deleated_at timestamp
);

CREATE TABLE books_categories (
	book_id VARCHAR(64) REFERENCES books (id),
	category_id VARCHAR(64) REFERENCES categories (id),
	created_at timestamp,
	updated_at timestamp,
	deleated_at timestamp
);