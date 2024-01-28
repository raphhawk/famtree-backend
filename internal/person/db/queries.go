package db

// data format: yyyymmdd

var (
	InitQuery = `
	DROP TABLE person;
	CREATE TABLE IF NOT EXISTS person (
		p_id BIGSERIAL primary key,
		name TEXT,
		dob DATE,
		email TEXT,
		c_at timestamp,
		u_at  timestamp
	);
	INSERT INTO person ( name, dob, email) VALUES (
		'test user M',
		'11-11-1111',
		'test@test.com'
	);`

	GetAll = `
	SELECT p_id, name, dob, email
	FROM person;`

	GetByIdQuery = `
	SELECT p_id, name, dob, email
	FROM person WHERE p_id=$1;`

	GetByEmailQuery = `
	SELECT p_id, name, dob, email
	FROM person WHERE email=%v;`

	CreateQuery = `
	INSERT INTO person (name, dob, email, c_at, u_at) VALUES (
		$1, $2, $3, $4, $5
	);`

	UpdateName   = `UPDATE person SET name=$1, u_at = $2 WHERE p_id = $3;`
	UpdateEmail  = `UPDATE person SET email = $1, u_at = $2 WHERE p_id = $3;`
	UpdateDob    = `UPDATE person SET dob = $1, u_at = $2 WHERE p_id = $3;`
	UpdateGender = `UPDATE person SET gender = $1, u_at = $2 WHERE p_id = $3;`
	DeleteById   = `DELETE FROM person WHERE p_id = $1;`
)
