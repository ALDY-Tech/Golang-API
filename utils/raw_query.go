package utils

type Queries struct {
	Insert  string
	SelectAll string
	SelectByID string
	Search string
	Update  string
	Delete  string
}

var (
	CustomerQueries = Queries{
		Insert:    `INSERT INTO customers (id, name, phonenumber, address) VALUES (:id, :name, :phonenumber, :address);`,
		SelectAll: `SELECT id, name, phonenumber, address FROM customers LIMIT $1 OFFSET $2;`,
		SelectByID: `SELECT id, name, phonenumber, address FROM customers WHERE id = $1;`,
		Update:    `UPDATE customers SET name = :name, phonenumber = :phonenumber, address = :address WHERE id = :id;`,
		Delete:    `DELETE FROM customers WHERE id = $1;`,
	}

	EmployeeQueries = Queries{
		Insert:    `INSERT INTO employees (id, name, phonenumber, address) VALUES (:id, :name, :phonenumber, :address);`,
		SelectAll: `SELECT id, name, phonenumber, address FROM employees LIMIT $1 OFFSET $2;`,
		SelectByID: `SELECT id, name, phonenumber, address FROM employees WHERE id = $1;`,
		Update:    `UPDATE employees SET name = :name, phonenumber = :phonenumber, address = :address WHERE id = :id;`,
		Delete:    `DELETE FROM employees WHERE id = $1;`,
	}

	ProductQueries = Queries{
		Insert:    `INSERT INTO products (id, name, price, unit) VALUES (:id, :name, :price, :unit);`,
		SelectAll: `SELECT id, name, price, unit FROM products LIMIT $1 OFFSET $2;`,
		Search: `WHERE name ILIKE '%' || $1 || '%';`,
		SelectByID: `SELECT id, name, price, unit FROM products WHERE id = $1;`,
		Update:    `UPDATE products SET name = :name, price = :price, unit = :unit WHERE id = :id;`,
		Delete:    `DELETE FROM products WHERE id = $1;`,
	}
)
