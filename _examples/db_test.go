package _examples

import (
	"PastePlus/core/plugin/db"
	"fmt"
	"testing"
)

func TestCreateDBFile(t *testing.T) {
	fmt.Println(db.FindPasteListByGTDate("2023-10-01"))
}
