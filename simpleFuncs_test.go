package main

import "testing"

func TestPutThenGet(t *testing.T) {
	mapper := make(map[string]int)

	put("KTA", 27, mapper)
	val, found := get("KTA", mapper)
	if !found {
		t.Fatalf("expected key KTA to be found")
	}
	if val != 27 {
		t.Errorf("get(KTA) = %d, want 27", val)
	}
}

func TestGetMissingKey(t *testing.T) {
	mapper := make(map[string]int)
	put("KTA", 27, mapper)

	_, found := get("STA", mapper)
	if found {
		t.Fatalf("Key doesn't exist but was found")
	}

}

func TestDelete(t *testing.T) {
	mapper := make(map[string]int)
	put("KTA", 27, mapper)

	test := del("KTA", mapper)
	if !test {
		t.Fatalf("Key exists and should return true instead returned false")
	}
	_, found := get("KTA", mapper)
	if found {
		t.Fatalf("Key shouldn't be found as it was just deleted")
	}

	doubleFail := del("KTA", mapper)
	if doubleFail {
		t.Fatalf("Should've failed as the key was already deleted")
	}
}
