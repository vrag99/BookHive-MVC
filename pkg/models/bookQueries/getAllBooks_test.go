package bookQueries

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "bookName", "quantity", "availableQuantity"}).
		AddRow(1, "Book 1", 10, 8).
		AddRow(2, "Book 2", 5, 4)

	mock.ExpectQuery("select \\* from books").
		WillReturnRows(rows)

	_, err = GetAllBooks(db)
	if err != nil {
		t.Errorf("error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
