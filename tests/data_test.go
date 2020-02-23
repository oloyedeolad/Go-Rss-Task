package tests

import (
	"rssfeed/datapack"
	"testing"
)

func TestConnectDB(t *testing.T) {
	collection := datapack.ConnectDB()

	if collection == nil {
		t.Error("collection not available")
	}
}
