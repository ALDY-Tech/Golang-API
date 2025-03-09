package utils

const (
	INSERT_CUSTOMER = `INSERT INTO Customers (name, phoneNumber, address) VALUES ($1, $2, $3) RETURNING id`
	SELECT_ALL_CUSTOMER = `SELECT id, name, phoneNumber, address FROM Customers LIMIT $1 OFFSET $2`
	SELECT_CUSTOMER_BY_ID = `SELECT id, name, phoneNumber, address FROM Customers WHERE id = $1`
	UPDATE_CUSTOMER = `UPDATE Customers SET name = :name, phonenumber = :phonenumber, address = :address WHERE id =:`
	DELETE_CUSTOMER = `DELETE FROM Customers WHERE id = $1`
)