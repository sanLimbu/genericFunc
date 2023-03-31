package safemap

import "testing"

func TestSafeMapInsertGet(t *testing.T) {
	m := New[int, int]()
	for i := 0; i < 500000; i++ {
		go func(i int) {
			m.Insert(i, i*3)
			value, err := m.Get(i)
			if err != nil {
				t.Error(err)
			}
			if value != i*3 {
				t.Errorf("%d should be %d", i, i*6)
			}
		}(i)
	}
}
