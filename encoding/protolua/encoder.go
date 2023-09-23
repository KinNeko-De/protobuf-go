package protolua

import (
	"errors"
	"math/bits"
	"strconv"
	"strings"
	"unicode/utf8"
)

const DefaultIndent = "  "

const KeyAssign = "="
const NullValue = "nil"
const BoolTrue = "true"
const BoolFalse = "false"
const BeginString = "\""
const EndString = "\""
const ArrayOpen = "{"
const ArrayClose = "}"
const TableOpen = "{"
const TableClose = "}"

type kind uint8

const (
	_ kind = (1 << iota) / 2
	key
	scalar
	objectOpen
	objectClose
	arrayOpen
	arrayClose
)

type Encoder struct {
	indent   string
	lastKind kind
	indents  []byte
	out      []byte
}

func NewEncoder(indent string) (*Encoder, error) {
	err := checkForInvalidIndentChars(indent)
	if err != nil {
		return nil, err
	}

	e := createEncoder(indent)
	return e, nil
}

func checkForInvalidIndentChars(indent string) error {
	if strings.Trim(indent, " \t") != "" {
		return errors.New("indent must be space or tab characters")
	}
	return nil
}

func createEncoder(indent string) *Encoder {
	e := &Encoder{}
	e.indent = indent
	return e
}

func (e *Encoder) Bytes() []byte {
	return e.out
}

func (e *Encoder) WriteNull() {
	e.prepareNext(scalar)
	e.out = append(e.out, NullValue...)
}

func (e *Encoder) WriteBool(b bool) {
	e.prepareNext(scalar)
	if b {
		e.out = append(e.out, BoolTrue...)
	} else {
		e.out = append(e.out, BoolFalse...)
	}
}

// WriteString escaped the string according to rules needed for luatex
func (e *Encoder) WriteString(s string) error {
	e.prepareNext(scalar)
	e.StartString()
	defer e.EndString()
	var err error
	if e.out, err = appendString(e.out, s); err != nil {
		return err
	}
	return nil
}

func (e *Encoder) WriteNumber(number string) {
	e.prepareNext(scalar)
	e.out = append(e.out, number...)
}

func (e *Encoder) StartObject() {
	e.prepareNext(objectOpen)
	e.out = append(e.out, TableOpen...)
}

func (e *Encoder) EndObject() {
	e.prepareNext(objectClose)
	e.out = append(e.out, TableClose...)
}

func (e *Encoder) WriteKey(s string) error {
	e.prepareNext(key)
	e.out = append(e.out, s...)
	e.writeKeyAssign()
	return nil
}

func (e *Encoder) StartArray() {
	e.prepareNext(arrayOpen)
	e.out = append(e.out, ArrayOpen...)
}

func (e *Encoder) EndArray() {
	e.prepareNext(arrayClose)
	e.out = append(e.out, ArrayClose...)
}

func (e *Encoder) StartString() {
	e.out = append(e.out, BeginString...)
}

func (e *Encoder) EndString() {
	e.out = append(e.out, EndString...)
}

func (e *Encoder) WriteIndexedList(i int) {
	e.prepareNext(key)
	e.out = append(e.out, "["...)
	e.out = append(e.out, strconv.FormatInt(int64(i), 10)...)
	e.out = append(e.out, "]"...)
	e.writeKeyAssign()
}

func (e *Encoder) writeKeyAssign() {
	if len(e.indent) != 0 {
		e.out = append(e.out, " "...)
	}
	e.out = append(e.out, KeyAssign...)
	if len(e.indent) != 0 {
		e.out = append(e.out, " "...)
	}
}

// prepareNext adds possible comma and indentation for the next value based on the previous type and indent option.
func (e *Encoder) prepareNext(next kind) {
	defer func() {
		e.lastKind = next
	}()

	if len(e.indent) == 0 {
		if e.lastKind&(scalar|objectClose|arrayClose) != 0 &&
			next&(key|scalar|objectOpen|arrayOpen) != 0 {
			e.out = append(e.out, ',')
		}
		return
	}

	switch {
	case e.lastKind&(objectOpen|arrayOpen) != 0:
		if next&(objectClose|arrayClose) == 0 {
			e.indents = append(e.indents, e.indent...)
			e.out = append(e.out, '\n')
			e.out = append(e.out, e.indents...)
		}

	case e.lastKind&(scalar|objectClose|arrayClose) != 0:
		switch {
		case next&(key|scalar|objectOpen|arrayOpen) != 0:
			e.out = append(e.out, ',', '\n')

		case next&(objectClose|arrayClose) != 0:
			e.indents = e.indents[:len(e.indents)-len(e.indent)]
			e.out = append(e.out, '\n')
		}
		e.out = append(e.out, e.indents...)
	}

}

func appendString(out []byte, in string) ([]byte, error) {
	i := indexNeedEscapeInString(in)
	in, out = in[i:], append(out, in[:i]...)
	for len(in) > 0 {
		switch r, n := utf8.DecodeRuneInString(in); {
		case r == utf8.RuneError && n == 1:
			return out, errors.New("the string contains invalid UTF-8 characters")
		case r < ' ' || r == '"' || r == '\\' || r == '%':
			switch r {
			case '%': // need to be escaped twice
				out = append(out, "\\\\"...)
				out = append(out, byte(r))
			case '\\': // need to be escaped once
				out = append(out, "\\"...)
				out = append(out, byte(r))
			case '"': // need to be escaped once
				out = append(out, '\\')
				out = append(out, byte(r))
			case '\b':
				return out, errors.New("not implemented yet")
			case '\f':
				return out, errors.New("not implemented yet")
			case '\n': // replaced with a latex line break that needs to be escaped
				out = append(out, "\\\\\\\\newline"...)
			case '\r':
				// do nothing as \r\n and \n are reduced to line break
			case '\t':
				return out, errors.New("not implemented yet")
			default:
				out = append(out, 'u')
				out = append(out, "0000"[1+(bits.Len32(uint32(r))-1)/4:]...)
				out = strconv.AppendUint(out, uint64(r), 16)
			}
			in = in[n:]
		default:
			i := indexNeedEscapeInString(in[n:])
			in, out = in[n+i:], append(out, in[:n+i]...)
		}
	}
	return out, nil
}

// indexNeedEscapeInString returns the index of the character that needs
// escaping. If no characters need escaping, the input length will be returned.
func indexNeedEscapeInString(s string) int {
	for i, r := range s {
		if r < ' ' || r == '\\' || r == '"' || r == '%' || r == utf8.RuneError {
			return i
		}
	}
	return len(s)
}
