package domain

import (
	"database/sql"
	"time"

	"github.com/ameydev/banking-app/errs"
	"github.com/ameydev/banking-app/logger"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		logger.Error("Unexpected DB error " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zip, &c.DateofBirth, &c.Status)
		if err != nil {
			logger.Error("Unexpected DB error " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected DB error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryDB(dbURL string) CustomerRepositoryDB {
	client, err := sql.Open("mysql", dbURL)
	if err != nil {
		logger.Error("Error while connecting to SQL server " + err.Error())
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client: client}
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zip, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("No rows found " + err.Error())
			return nil, errs.NewNotFoundError("No customer found")
		} else {
			logger.Error("Unexpected DB error " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected DB error")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDB) ByStatus(status string) ([]Customer, *errs.AppError) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		return nil, errs.NewNotFoundError("Wrong request")
	}

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"

	rows, err := d.client.Query(findAllSql, status)
	if err != nil {
		logger.Error("Unexpected DB error " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zip, &c.DateofBirth, &c.Status)
		if err != nil {
			logger.Error("Unexpected DB error " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected DB error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}
