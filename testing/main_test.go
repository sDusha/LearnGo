package printer

import "testing"

type Test struct {
	in  []byte
	out int
	err error
}

var tests = []Test{
	{[]byte("Hello World!"), 12, nil},
	{[]byte(""), 0, nil},
	{[]byte{}, 0, nil},
	{[]byte("asd"), 3, nil},
	{[]byte("фывыф"), 5, nil},
	{[]byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8},
}

func TestMultiply(t *testing.T) {
	for i, test := range tests {
		multi, err := GetUTFLength(test.in)
		if multi != test.out || err != test.err {
			t.Errorf("#%d: DeleteVowels(%s)=%d; want %d", i, test.in, multi, test.out)
		}
	}
}
