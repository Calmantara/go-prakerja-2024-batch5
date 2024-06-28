// golang convention: untuk file test ditambahkan _test.go di file yang akan kita test
package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("add positif number", func(t *testing.T) {
		res := Add(1, 2)
		if res != 3 {
			t.Errorf("wrong result, expected 3 got %v", res)
			t.Fail()
			return
		}
		t.Log("ok")
	})

	t.Run("add negative number", func(t *testing.T) {
		res := Add(-10, -20)
		if res != -30 {
			t.Errorf("wrong result, expected -30 got %v", res)
			t.Fail()
			return
		}
		t.Log("ok")
	})
}
