package mylib

import "testing"

func Test_BasicChecks(t *testing.T) {
	t.Parallel()
	t.Run("Go can add", func(t *testing.T) {
		if 1+1 != 2 {
			t.Fail()
		}
	})
	t.Run("Go can concatenate", func(t *testing.T) {
		if "Hello, "+"Go" != "Hello, Go" {
			t.Fail()
		}
	})
}

func TestAddingNumbers(t *testing.T) {
	t.Run("Add 1 and 2 returns 3", func(t *testing.T) {
		r := add(1, 2)
		if r != 3 {
			t.Fail()
		}
	})
}

func BenchmarkAddingNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(2, 45)
	}
}
