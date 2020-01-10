//
//  Base 58
//
//  Copyright (c) 2020 Wellington Marthas
//
//  Permission is hereby granted, free of charge, to any person obtaining a copy
//  of this software and associated documentation files (the "Software"), to deal
//  in the Software without restriction, including without limitation the rights
//  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the Software is
//  furnished to do so, subject to the following conditions:
//
//  The above copyright notice and this permission notice shall be included in
//  all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//  THE SOFTWARE.
//

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
