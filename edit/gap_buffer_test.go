package edit

import "testing"

func TestGapBufferInsert(t *testing.T) {
	gb := NewGapBuffer(1024)

	tests := []struct {
		b    byte
		want string
	}{
		{'t', "t"},
		{'e', "te"},
		{'s', "tes"},
		{'t', "test"},
	}

	for i, test := range tests {
		gb.Insert(test.b)
		got := gb.String()
		if test.want != got {
			t.Errorf("test %d: want: %s, got %s", i, test.want, got)
		}
	}
}

func TestGapBufferDelete(t *testing.T) {
	gb := NewGapBuffer(1024)

	gb.Insert('t')
	gb.Insert('e')
	gb.Insert('s')
	gb.Insert('t')

	wants := []string{
		"tes",
		"te",
		"t",
		"",
	}

	for i, want := range wants {
		gb.Delete()
		got := gb.String()
		if want != got {
			t.Errorf("test %d: want: %s, got %s", i, want, got)
		}
	}
}
