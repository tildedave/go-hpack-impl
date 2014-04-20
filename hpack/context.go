package hpack

import (
	"errors"
	"fmt"
)

type EncodingContext struct {
	HeaderTable HeaderTable
	ReferenceSet ReferenceSet
	Update struct {
		ReferenceSetEmptying bool
		MaximumHeaderTableSizeChange int
	}
}

func NewEncodingContext() *EncodingContext {
	context := &EncodingContext{}
	context.HeaderTable.MaxSize = 1024

	return context
}

func (context *EncodingContext) AddHeader(h HeaderField) {
	ref := context.HeaderTable.AddHeader(h)
	refset := &context.ReferenceSet

	refset.Entries = append(refset.Entries, ref)
}

func (context *EncodingContext) EncodeField(h HeaderField) string {
	var idx int

	table := &context.HeaderTable
	idx = table.ContainsHeader(h)
	if idx != 0 {
		a := make([]byte, 1)
		a[0] = byte(idx)
		a[0] |= 0x80

		table.AddHeader(h)
		encodedHeaders := string(a)
		return encodedHeaders
	}

	idx = table.ContainsName(h.Name)
	if idx != 0 {
		a := make([]byte, 2)
		a[0] = byte(idx)
		a[0] |= 0x40
		a[1] = byte(len(h.Value))

		table.AddHeader(h)
		encodedHeaders := string(a) + h.Value
		return encodedHeaders
	}

	// Literal name, literal value
	table.AddHeader(h)
	encodedHeaders := ""
	encodedHeaders += string(0x40)
	encodedHeaders += string(byte(len(h.Name)))
	encodedHeaders += h.Name
	encodedHeaders += string(byte(len(h.Value)))
	encodedHeaders += h.Value

	return string(encodedHeaders)
}

func (context *EncodingContext) Encode(hs HeaderSet) string {
	encoded := ""

	if context.Update.ReferenceSetEmptying {
		context.ReferenceSet = ReferenceSet{}
		context.Update.ReferenceSetEmptying = false
		encoded += "\x30"
	}

	refset := &context.ReferenceSet

	for _, h := range hs.Headers {
		mustEncode := true

		for _, refHeader := range refset.Entries {
			if *refHeader == h {
				mustEncode = false
			}
		}

		if mustEncode {
			encoded += context.EncodeField(h)
			// Not the correct way to do this
			refset.Entries = append(refset.Entries, &HeaderField{h.Name, h.Value})
		}
	}
	return encoded
}

const (
	IndexedMask = 0x80
	LiteralIndexedMask = 0x40
	LiteralNeverIndexedMask = 0x10
	ContextUpdateMask = 0x20
	LiteralNoIndexMask = 0x00
)

func fmtIsNotUnused() {
	fmt.Println("line to not complain about unused fmt import")
}

func unpackLiteral(wireBytes *[]byte) (string) {
	length := int((*wireBytes)[0] & 0x4F)
	str := string((*wireBytes)[1:1 + length])

	*wireBytes = (*wireBytes)[1 + length:]

	return str
}

func decodeLiteralHeader(wireBytes *[]byte, table *HeaderTable) (HeaderField) {
	nameIndex := (*wireBytes)[0] & 0x2F
	*wireBytes = (*wireBytes)[1:]

	if nameIndex == byte(0) {
		name := unpackLiteral(wireBytes)
		value := unpackLiteral(wireBytes)

		return HeaderField{ name, value }
	}

	nameHeader := table.HeaderAt(int(nameIndex))
	value := unpackLiteral(wireBytes)
	return HeaderField{ nameHeader.Name, value }
}

func (context *EncodingContext) Decode(wire string) (hs HeaderSet, err error) {
	headers := []HeaderField{}
	wireBytes := []byte(wire)

	table := &context.HeaderTable
	refset := &context.ReferenceSet

	for ; len(wireBytes) > 0 ; {
		if wireBytes[0] & ContextUpdateMask == ContextUpdateMask {
			if wireBytes[0] & 0x30 == 0x30 {
				// empty reference set
				refset.Entries = []*HeaderField{}
			}
			wireBytes = wireBytes[1: ]
			continue
		}

		if wireBytes[0] & IndexedMask == IndexedMask {
			index := wireBytes[0] & 0x4F
			header := table.HeaderAt(int(index))
			headers = append(headers, header)
			context.AddHeader(header)

			wireBytes = wireBytes[1: ]

			continue
		}

		if wireBytes[0] & LiteralIndexedMask == LiteralIndexedMask {
			header := decodeLiteralHeader(&wireBytes, table)
			headers = append(headers, header)
			context.AddHeader(header)
			continue
		}

		if wireBytes[0] & LiteralNoIndexMask == LiteralNoIndexMask {
			header := decodeLiteralHeader(&wireBytes, table)
			headers = append(headers, header)
			continue
		}

		return HeaderSet{}, errors.New("Could not decode")
	}

	for _, h := range refset.Entries {
		found := false

		for _, emitted := range headers {
			if emitted == *h {
				found = true
			}
		}

		if !found {
			headers = append(headers, *h)
		}
	}

	return HeaderSet{ headers }, nil
}
