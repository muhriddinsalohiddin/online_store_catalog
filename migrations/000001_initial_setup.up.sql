CREATE TABLE authors(
	id UUID NOT NULL PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleated_at timestamp	
);

CREATE TABLE books (
	id UUID NOT NULL PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	author_id UUID NOT NULL REFERENCES authors (id),
	created_at timestamp,
	updated_at timestamp,
	deleated_at timestamp
);

CREATE TABLE categories (
	id UUID NOT NULL PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	parent_id UUID,
	created_at timestamp,
	updated_at timestamp,
	deleated_at timestamp
);

CREATE TABLE books_categories (
	book_id UUID REFERENCES books (id),
	category_id UUID REFERENCES categories (id)
);