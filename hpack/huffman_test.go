package hpack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombineHuffmanNothingWithSomething(t *testing.T) {
	overflow, code := combineHuffman(
		HuffmanCode{},
		HuffmanCode{0x03ffffba, 26},
	)

	assert.Equal(t, code, HuffmanCode{ 0x03ffffba, 26 })
	assert.Equal(t, overflow, "")
}

func TestCombineHuffman32Bits(t *testing.T) {
	overflow, code := combineHuffman(
		HuffmanCode{0x1f, 6},
		HuffmanCode{0x03ffffba, 26},
	)

	assert.Equal(t, code, HuffmanCode{ 0x00000000, 0 })
	assert.Equal(t, overflow, "\x7f\xff\xff\xba")
}

func TestCombineHuffmanWithLessThan32Bits(t *testing.T) {
	overflow, code := combineHuffman(
		HuffmanCode{0x1f, 6},
		HuffmanCode{0x07, 5},
	)

/*
 * 00|01 1111 -- 6 bits
 * 000|0 0111 -- 5 bits
 * 0|011 1110 0111 -- 11 bits
 */
	assert.Equal(t, code.bitLength, uint(11))
	assert.Equal(t, code.bits, uint32(0x000003e7))
	assert.Equal(t, overflow, "")
}

func TestCombineHuffmanWithOverflow(t *testing.T) {
	overflow, code := combineHuffman(
		HuffmanCode{0x03ffffba, 26},
		HuffmanCode{0x03ffffd2, 26},
	)

	assert.Equal(t, code.bitLength, uint(20))
	assert.Equal(t, code.bits, uint32(0x000fffd2))
	assert.Equal(t, overflow, "\xff\xff\xee\xbf", "Encoded octet did not match")
/*
0000 00|11 1111 1111 1111 1111 1011 1010 -- 26 bits
0000 00|11 1111 1111 1111 1111 1101 0010 -- 26 bits

1111 1111 1111 1111 1110 1110 1011 1111 -- overflow
0000 0000 0000 1111 1111 1111 1101 0010 -- remaining code
*/
}

func TestCombineHuffmanWithEOS(t *testing.T) {
	overflow, code := combineHuffman(
		HuffmanTable[0x43],
		HuffmanEOS,
	)

	assert.Equal(t, code, HuffmanCode{ 0x00000000, 1 })
	assert.Equal(t, overflow, "\xee\xff\xff\xee")
}

func TestEncodeSingleCharacterHuffman(t *testing.T) {
	encoded := EncodeHuffman("C")

/*
"C" -> 1110 1110
EOS -> 000|1 1111 1111 1111 1111 1101 1100

1110 1110 1111 1111 1111 1111 1110 1110 | 0

1ffffdc

 */

	assert.Equal(t, len(encoded), 4)
	assert.Equal(t, encoded, "\xee\xff\xff\xee")
}

func TestDecodeSingleCharacterHuffman(t *testing.T) {

}
