package db

import "testing"

func TestDbConn(t *testing.T) {
	if _, err := NewDB(); err != nil {
		t.Errorf("db conn error %v", err)
	}
}
