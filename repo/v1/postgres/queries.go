package postgres

var createHellosTableQuery = `
CREATE TABLE if not exists hellos (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	hello_word VARCHAR
);`

var helloInsertQuery = `
INSERT INTO hellos (
	hello_word
) VALUES ($1) RETURNING id;`
