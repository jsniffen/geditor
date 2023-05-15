package edit

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

func (pt *PieceTable) Insert(i int) {
}

func (pt *PieceTable) Delete(idx int) {
}

func (pt *PieceTable) String() string {
	s := ""
	for entry := pt.table; entry != nil; entry = entry.next {
		b := pt.add
		if entry.original {
			b = pt.original
		}
		s += string(b[entry.idx:entry.len])
	}
	return s
}

func (pt *PieceTable) findEntry(idx int) (*Entry, int) {
	len := 0
	for entry := pt.table; entry != nil; entry = entry.next {
		if idx <= len + entry.len {
			return entry, idx - len
		}

		len += entry.len
	}
	return nil, 0
}
