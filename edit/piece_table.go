package edit

import "fmt"

type Entry struct {
	original bool
	idx      int
	len      int
	next     *Entry
	prev     *Entry
}

type PieceTable struct {
	add      []byte
	original []byte
	table    *Entry
}

func NewPieceTable(contents string) (*PieceTable, error) {
	e := Entry{
		original: true,
		idx: 0,
		len: len(contents),
		next: nil,
		prev: nil,
	}

	pt := PieceTable{
		add: []byte(""),
		original: []byte(contents),
		table: &e,
	}
	return &pt, nil
}

func (pt *PieceTable) Insert(b byte, i int) {
	currEntry, n := pt.findEntry(i)

	if currEntry == nil {
		return
	}

	pt.add = append(pt.add, b)

	newEntry := &Entry{
		original: false,
		idx: len(pt.add)-1,
		len: 1,
	}

	if n == 0 {
		// inserting at the beginning of entry
		newEntry.prev = currEntry.prev
		newEntry.next = currEntry
		currEntry.prev = newEntry

		if pt.table == currEntry {
			pt.table = newEntry
		}
	} else if n == currEntry.len - 1 {
		// inserting at the end of entry

		// if we're adding to the end of an add buffer
		// entry and it just so happens that this entry
		// ends at the end of the add buffer, we can just extend
		// this entry's length by one and call it a day.
		if !currEntry.original && currEntry.idx + currEntry.len == len(pt.add) - 1 {
			currEntry.len += 1
		} else {
			newEntry.next = currEntry.next
			newEntry.prev = currEntry
			currEntry.next = newEntry
		}
	} else {
		// split the entry into 2
		endEntry := &Entry{
			original: currEntry.original,
			idx: currEntry.idx+n,
			len: currEntry.len-n,
			next: currEntry.next,
			prev: newEntry,
		}

		currEntry.len -= n
		currEntry.next = newEntry

		newEntry.prev = currEntry
		newEntry.next = endEntry
	}
}

func (pt *PieceTable) Delete(i int) {
	currEntry, n := pt.findEntry(i)

	if currEntry == nil {
		return
	}

	if n == 0 {
		// deleting at beginning of entry
		currEntry.idx += 1
		currEntry.len -= 1
	} else if n == currEntry.len - 1 {
		// deleting at end of entry
		currEntry.len -= 1
	} else {
		// deleting in the middle of entry
		// so split in 2
	}
}

func (pt *PieceTable) String() string {
	s := ""
	for entry := pt.table; entry != nil; entry = entry.next {
		b := pt.add
		if entry.original {
			b = pt.original
		}
		s += string(b[entry.idx:entry.idx+entry.len])
	}
	return s
}

func (pt *PieceTable) findEntry(i int) (*Entry, int) {
	if i < 0 {
		return nil, 0
	}

	len := 0
	for entry := pt.table; entry != nil; entry = entry.next {
		if i < len + entry.len {
			return entry, i - len
		}

		len += entry.len
	}
	return nil, 0
}

func (pt *PieceTable) printDebug() {
	for entry := pt.table; entry != nil; entry = entry.next {
		fmt.Println(entry)
	}
}

func (pt *PieceTable) Len() int {
	len := 0
	for entry := pt.table; entry != nil; entry = entry.next {
		len += 1
	}
	return len
}
