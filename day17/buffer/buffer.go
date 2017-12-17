package buffer

import "fmt"
import "strings"

// Circular represents a circular buffer
type Circular struct {
	start     *bufferItem
	cursor    *bufferItem
	items     []*bufferItem
	nextIndex int
}

// Insert adds the given value immediately after the current cursor.
// It sets the cursor to the inserted item and returns the index of the
// inserted item.
func (c *Circular) Insert(val int) int {
	// Crate a new buffer item
	newItem := &bufferItem{value: val}
	c.items = append(c.items, newItem)
	if c.cursor == nil {
		// must be inserting first item so make the item self referential
		newItem.next = newItem
		// and set start
		c.start = newItem
	} else {
		// make new item point at the item currently pointed at by the cursor
		newItem.next = c.cursor.next
		// then set the cursor to point at the new item
		c.cursor.next = newItem
	}
	// move cursor to new item
	c.cursor = newItem

	// return the index of the item just inserted
	return len(c.items) - 1
}

// MoveTo changes the cursor to the given item index and returns the buffer to allow chaining.
func (c *Circular) MoveTo(itemIndex int) *Circular {
	c.cursor = c.items[itemIndex]
	return c
}

// Move steps the cursor along the given number of places and returns the buffer to allow chaining.
func (c *Circular) Move(places int) *Circular {
	if c.cursor == nil {
		panic("buffer is empty")
	}
	for p := 0; p < places; p++ {
		c.cursor = c.cursor.next
	}
	return c
}

// Current returns the value at the cursor
func (c *Circular) Current() int {
	if c.cursor == nil {
		panic("no item at cursor")
	}

	return c.cursor.value
}

// String implements Stringer for Circular
func (c *Circular) String() string {
	if len(c.items) == 0 {
		return "[]"
	}

	out := []string{}
	current := c.start
	for ok := true; ok; ok = current != c.start {
		thisVal := ""
		if current == c.cursor {
			thisVal += "("
		}
		thisVal += fmt.Sprintf("%d", current.value)
		if current == c.cursor {
			thisVal += ")"
		}
		out = append(out, thisVal)
		current = current.next
	}
	return strings.Join(out, " ")
}

type bufferItem struct {
	value int
	next  *bufferItem
}
