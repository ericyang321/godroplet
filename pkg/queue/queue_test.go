package queue

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	t.Run("returns true when queue has no items", func(t *testing.T) {
		q := new(Queue)
		if q.Empty() != true {
			t.Errorf("Queue should return empty true when no items have been inserted, but instead found")
		}
	})

	t.Run("returns false when queue has items", func(t *testing.T) {
		q := new(Queue)
		q.Push(0)
		if q.Empty() != false {
			t.Errorf("Queue should return empty false when an item has been inserted")
		}
	})
}

func TestPush(t *testing.T) {
	t.Run("Push adds item into queue", func(t *testing.T) {
		q := new(Queue)
		item := 0
		q.Push(item)
		wasPushed := q.Pop()

		if wasPushed != item {
			t.Errorf("Queue should be pushing item %d into queue", item)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("Pop removes the least recently popped item", func(t *testing.T) {
		q := new(Queue)
		pushed := []Any{0, 1, 2, 3, 4, 5}

		expectedPopped := []Any{0, 1, 2, 3, 4, 5}
		actuallyPopped := []Any{}

		for _, num := range pushed {
			q.Push(num)
		}
		for q.Empty() == false {
			actuallyPopped = append(actuallyPopped, q.Pop())
		}
		for i := range actuallyPopped {
			expected := expectedPopped[i]
			actual := actuallyPopped[i]
			if expected != actual {
				t.Errorf("Queue should be popping items in reverse order. Expected %d, but was given %d", expected, actual)
			}
		}
	})
}
