package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occured while creating a stub db connection: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "bookName", "quantity", "availableQuantity"}).
		AddRow(1, "book1", 10, 6).
		AddRow(2, "book2", 5, 4)
	
	mock.ExpectQuery("select * from books where quantity >= 1").WillReturnRows(rows)
}
