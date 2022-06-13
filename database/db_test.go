package database

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	dns := "root@tcp(localhost:3306)/country"
	_, err := ConnectDB(dns)
	if err != nil {
		t.Fatalf("error '%s' when attempting to connect to local database server", err)
	}
}
