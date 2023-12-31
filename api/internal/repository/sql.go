package repository

var sqlCreateCustomer = `
	INSERT INTO customers(first_name, last_name, birth_date) 
	VALUES(lower(:first_name), lower(:last_name), :birth_date) 
	RETURNING
		id,
		first_name, 
		last_name, 
		TO_CHAR(birth_date, 'YYYY-MM-DD') AS birth_date,
		created_at,
		updated_at,
		deleted_at
`
var sqlGetCustomers = `
	SELECT 
		id,
		first_name,
		last_name,
		TO_CHAR(birth_date, 'YYYY-MM-DD') AS birth_date,
		updated_at,
		created_at
	FROM customers
	WHERE deleted_at IS NULL
`

var sqlGetCustomerById = `
	SELECT 
		id,
		first_name,
		last_name,
		TO_CHAR(birth_date, 'YYYY-MM-DD') AS birth_date,
		updated_at,
		created_at
	FROM customers
	WHERE id = $1 AND deleted_at IS NULL
`

var sqlUpdateCustomerById = `
	UPDATE customers
	SET 
		first_name = lower(:first_name),
		last_name = lower(:last_name),
		birth_date = :birth_date,
		updated_at = now()
	WHERE id = :id
	RETURNING
		id,
		first_name, 
		last_name, 
		TO_CHAR(birth_date, 'YYYY-MM-DD') AS birth_date,
		created_at,
		updated_at,
		deleted_at
`

var sqlDeleteCustomerById = `
	DELETE FROM customers WHERE id = $1
`
