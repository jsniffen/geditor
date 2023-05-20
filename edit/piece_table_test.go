package edit

import "testing"

func TestNewPieceTable(t *testing.T) {
	pt, _ := NewPieceTable("test")
	got := pt.String()
	want := "test"
	if got != want {
		t.Errorf("Got: %s, want: %s", got, want)
	}
}

func TestFindEntry(t *testing.T) {
	pt, _ := NewPieceTable("test")

	tests := []struct {
		index int 
		entry *Entry
		entryIndex int
	}{
		{-1, nil, 0},
		{0, pt.table, 0},
		{3, pt.table, 3},
		{4, nil, 0},
	}

	for _, test := range tests {
		entry, entryIndex := pt.findEntry(test.index)

		if test.entry == nil && entry == nil {
			continue
		}

		if entry != test.entry || entryIndex != test.entryIndex {
			t.Errorf("want: %v, %v, got: %v, %v", &test.entry, test.entryIndex, &entry, entryIndex)
		}
	}
}

func TestInsert(t *testing.T) {
	pt, _ := NewPieceTable("test")

	tests := []struct {
		index int 
		b byte
		want string
		len int
	}{
		{-1, 'x', "test", 1},
		// {5, 'x', "testx", 2},
		// {0, 'x', "xtest", 2},
		// {3, 'x', "xtexst", 4},
		// {6, 'x', "xtexstx", 2},
	}

	for i, test := range tests {
		pt.Insert(test.b, test.index)
		got := pt.String()
		if test.want != got {
			t.Errorf("test %d: want: %s, got %s", i, test.want, got)
		}

		if test.len != pt.Len() {
			t.Errorf("test %d: want len: %d, got %d", i, test.len, pt.Len())
		}
	}
}

func TestDelete(t *testing.T) {
	pt, _ := NewPieceTable("tests")

	tests := []struct {
		index int 
		want string
		len int
	}{
		{-1, "tests", 1},
		{5, "tests", 1},
		{0, "ests", 1},
		{3, "est", 1},
	}

	for i, test := range tests {
		pt.Delete(test.index)
		got := pt.String()
		if test.want != got {
			t.Errorf("test %d: want: %s, got %s", i, test.want, got)
		}

		if test.len != pt.Len() {
			t.Errorf("test %d: want len: %d, got %d", i, test.len, pt.Len())
		}
	}
}
