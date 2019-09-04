package base58

import "testing"

func TestEncode(t *testing.T) {
	if EncodeToString([]byte("Hello World")) != "JxF12TrwUP45BMd" {
		t.Error("Enconded string is not valid")
	}
}

func TestDecode(t *testing.T) {
	output, e := DecodeString("JxF12TrwUP45BMd")

	if e != nil {
		t.Error(e)
		return
	}

	if string(output) != "Hello World" {
		t.Error("Decoded string is not valid")
	}
}

func TestDecodeInvalidInput(t *testing.T) {
	if _, e := DecodeString("SGVsbG8gV29ybGQ="); e == nil {
		t.Error("Decoded string is not valid")
	}
}

func TestEncodeDecode(t *testing.T) {
	testInput(t, "12345678")
	testInput(t, "%$#@!")
	testInput(t, "Contrary to popular belief, Lorem Ipsum is not simply random text")
	testInput(t, "Many desktop publishing packages and web page editors")
	testInput(t, "At vero eos et accusamus et iusto odio dignissimos ducimus qui voluptatum")
	testInput(t, "On the other hand, we denounce with righteous who are beguiled")
}

func testInput(t *testing.T, input string) {
	output, e := DecodeString(EncodeToString([]byte(input)))

	if e != nil {
		t.Error(e)
	} else if string(output) != input {
		t.Error("Decoded string is not valid")
	}
}
