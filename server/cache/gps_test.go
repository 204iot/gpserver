package cache

import "testing"

func TestSet(t *testing.T) {
	if err := Set(&IOTDevice{
		ID:        "1",
		Longitude: "123",
		Latitude:  "456",
	}); err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	if err := Set(&IOTDevice{
		ID:        "1",
		Longitude: "xxx",
		Latitude:  "yyy",
	}); err != nil {
		t.Fatal(err)
	}

	if err := Set(&IOTDevice{
		ID:        "1",
		Longitude: "123",
		Latitude:  "456",
	}); err != nil {
		t.Fatal(err)
	}

	d := Get("1")
	if d == nil {
		t.Fatal()
	}
	if d.Latitude != "456" {
		t.Fatal()
	}
	if d.Longitude != "123" {
		t.Fatal()
	}
}

func TestGetAll(t *testing.T) {
	if err := Set(&IOTDevice{
		ID:        "1",
		Longitude: "123",
		Latitude:  "456",
	}); err != nil {
		t.Fatal(err)
	}

	if err := Set(&IOTDevice{
		ID:        "2",
		Longitude: "123",
		Latitude:  "456",
	}); err != nil {
		t.Fatal(err)
	}

	devices := GetAll()
	if len(devices) != 2 {
		t.Fatal()
	}
}
