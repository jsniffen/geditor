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
		{0, pt.table, 0},
	}

	for _, test := range tests {
		entry, entryIndex := pt.findEntry(test.index)

		if entry != test.entry || entryIndex != test.entryIndex {
			t.Errorf("want: %v, %v, got: %v, %v", &test.entry, test.entryIndex, &entry, entryIndex)
		}
	}
}
