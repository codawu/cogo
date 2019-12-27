package integers

import "testing"

func TestAdd(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, except, sum int) {
		if sum != except {
			t.Errorf(" excepted '%d' but go '%d' ", except, sum)
		}
	}

	t.Run("add two number", func(t *testing.T) {
		sum := Add(2, 2)
		except := 4
		assertCorrectMessage(t, except, sum)
		sum = Add(5, 5)
		except = 10
		assertCorrectMessage(t, except, sum)
	})

}
