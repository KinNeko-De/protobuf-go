package protolua

import (
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	protoluainternal "github.com/kinneko-de/protobuf-go/internal/encoding/protolua"
	v2proto "github.com/kinneko-de/protobuf-go/internal/encoding/testing/proto/v2"
	v3proto "github.com/kinneko-de/protobuf-go/internal/encoding/testing/proto/v3"
	googleprotobufproto "github.com/kinneko-de/protobuf-go/internal/encoding/testing/proto/v3/google/protobuf"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestMarshal_Positive(t *testing.T) {
	tests := []struct {
		desc     string
		input    proto.Message
		option   LuaMarshalOption
		expected string
	}{
		{
			desc: "Message with one string",
			input: &v3proto.SingleString{
				Fut: "Max Mustermann",
			},
			expected: "SingleString={fut=\"Max Mustermann\"}",
		},
		{
			desc: "Message with multiple strings",
			input: &v3proto.MultipleStrings{
				Name:   "Max Mustermann",
				Street: "Musterstraße 17",
			},
			expected: "MultipleStrings={name=\"Max Mustermann\",street=\"Musterstraße 17\"}",
		},
		{
			desc: "Message with one int32",
			input: &v3proto.SingleInt32{
				Fut: 42,
			},
			expected: "SingleInt32={fut=42}",
		},
		{
			desc: "Message with multiple int32",
			input: &v3proto.MultipleInt32{
				One: 42,
				Two: 84,
			},
			expected: "MultipleInt32={one=42,two=84}",
		},
		{
			desc: "Message with one int64 positive",
			input: &v3proto.SingleInt64{
				Fut: 42000000001,
			},
			expected: "SingleInt64={fut=42000000001}",
		},
		{
			desc: "Message with multiple int64 positive",
			input: &v3proto.MultipleInt64{
				One: 420000000001,
				Two: 840000000001,
			},
			expected: "MultipleInt64={one=420000000001,two=840000000001}",
		},
		{
			desc: "Message with one int64 negative",
			input: &v3proto.SingleInt64{
				Fut: -42000000001,
			},
			expected: "SingleInt64={fut=-42000000001}",
		},
		{
			desc: "Message with multiple int64 negative",
			input: &v3proto.MultipleInt64{
				One: -420000000001,
				Two: -840000000001,
			},
			expected: "MultipleInt64={one=-420000000001,two=-840000000001}",
		},
		{
			desc: "Message with one uint64",
			input: &v3proto.SingleUInt64{
				Fut: 42000000001,
			},
			expected: "SingleUInt64={fut=42000000001}",
		},
		{
			desc: "Message with multiple uint64",
			input: &v3proto.MultipleUInt64{
				One: 420000000001,
				Two: 840000000001,
			},
			expected: "MultipleUInt64={one=420000000001,two=840000000001}",
		},
		{
			desc: "Message with one bool true",
			input: &v3proto.SingleBool{
				Fut: true,
			},
			expected: "SingleBool={fut=true}",
		},
		{
			desc: "Message with one bool false",
			input: &v3proto.SingleBool{
				Fut: false,
			},
			expected: "SingleBool={fut=false}",
		},
		{
			desc: "Message with multiple bool true-false",
			input: &v3proto.MultipleBool{
				One: true,
				Two: false,
			},
			expected: "MultipleBool={one=true,two=false}",
		},
		{
			desc: "Message with multiple bool false-true",
			input: &v3proto.MultipleBool{
				One: false,
				Two: true,
			},
			expected: "MultipleBool={one=false,two=true}",
		},
		{
			desc: "Message with one list",
			input: &v3proto.SingleListInt{
				Fut: []int32{42, 12},
			},
			expected: "SingleListInt={fut={[1]=42,[2]=12}}",
		},
		{
			desc: "Message with one empty list",
			input: &v3proto.SingleListInt{
				Fut: make([]int32, 0),
			},
			expected: "SingleListInt={fut={}}",
		},
		{
			desc: "Message with multiple list",
			input: &v3proto.MultipleList{
				One: []string{"42"},
				Two: []int32{42, 12},
			},
			expected: "MultipleList={one={[1]=\"42\"},two={[1]=42,[2]=12}}",
		},
		{
			desc: "Message with sub message",
			input: &v3proto.SingleMessage{
				Fut: &v3proto.SubMessage{
					Sone: "42",
					Stwo: &v3proto.SubSubMessage{
						Ssfut: "48",
					},
				},
			},
			expected: "SingleMessage={fut={sone=\"42\",stwo={ssfut=\"48\"}}}",
		},
		{
			desc: "Message with multiple sub message",
			input: &v3proto.MultipleMessage{
				One: &v3proto.SubMessage{
					Sone: "42",
					Stwo: &v3proto.SubSubMessage{
						Ssfut: "48",
					},
				},
				Two: []*v3proto.SubSubMessage{
					{
						Ssfut: "42",
					},
					{
						Ssfut: "48",
					},
				},
			},
			expected: "MultipleMessage={one={sone=\"42\",stwo={ssfut=\"48\"}},two={[1]={ssfut=\"42\"},[2]={ssfut=\"48\"}}}",
		},
		{
			desc: "Message with nil sub message",
			input: &v3proto.SingleMessage{
				Fut: nil,
			},
			expected: "SingleMessage={fut=nil}",
		},
		{
			desc: "Message with nil sub sub message",
			input: &v3proto.MultipleMessage{
				One: nil,
				Two: []*v3proto.SubSubMessage{
					nil,
					{
						Ssfut: "42",
					},
				},
			},
			expected: "MultipleMessage={one=nil,two={[1]=nil,[2]={ssfut=\"42\"}}}",
		},
		{
			desc: "Message with multiple strings (Multiline)",
			input: &v3proto.MultipleStrings{
				Name:   "Max Mustermann",
				Street: "Musterstraße 17",
			},
			option:   LuaMarshalOption{Format: struct{ Multiline bool }{Multiline: true}},
			expected: "MultipleStrings = {\n  name = \"Max Mustermann\",\n  street = \"Musterstraße 17\"\n}",
		},
		{
			desc: "Message with special character %",
			input: &v3proto.SingleString{
				Fut: "%",
			},
			expected: "SingleString={fut=\"\\\\%\"}",
		},
		{
			desc: "Message with special character carriage return (Mac OS 9 and lower are not supported)",
			input: &v3proto.SingleString{
				Fut: "\r",
			},
			expected: "SingleString={fut=\"\"}",
		},
		{
			desc: "Message with special character line feed",
			input: &v3proto.SingleString{
				Fut: "\n",
			},
			expected: "SingleString={fut=\" \\\\newline \"}",
		},
		{
			desc: "Message with special character carriage return and line feed",
			input: &v3proto.SingleString{
				Fut: "\r\n",
			},
			expected: "SingleString={fut=\" \\\\newline \"}",
		},
		{
			desc: "Message with special character \\",
			input: &v3proto.SingleString{
				Fut: "t\\t",
			},
			expected: "SingleString={fut=\"t\\\\t\"}",
		},
		{
			desc: "Message with special character \"",
			input: &v3proto.SingleString{
				Fut: "t\"t",
			},
			expected: "SingleString={fut=\"t\\\"t\"}",
		},
		{
			desc: "Message with empty",
			input: &v3proto.SingleString{
				Fut: "",
			},
			expected: "SingleString={fut=\"\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual, err := tt.option.Marshal(tt.input)
			AssertLuaString(actual, err, tt.expected, t)
		})
	}
}

func TestMarshal_Error(t *testing.T) {
	var dummy string = "test"

	tests := []struct {
		desc   string
		input  proto.Message
		option LuaMarshalOption
	}{
		{
			desc:  "nil message throws error",
			input: nil,
		},
		{
			desc: "v2 not supported",
			input: &v2proto.SingleString{
				Fut: &dummy,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			_, err := tt.option.Marshal(tt.input)
			if err == nil {
				t.Errorf("Marshal() returned no error: ")
			}
		})
	}
}

func TestMarshal_DefaultOption(t *testing.T) {
	tests := []struct {
		desc     string
		input    proto.Message
		expected string
	}{
		{
			desc: "Message with one string",
			input: &v3proto.SingleString{
				Fut: "Max Mustermann",
			},
			expected: "SingleString={fut=\"Max Mustermann\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual, err := Marshal(tt.input)
			AssertLuaString(actual, err, tt.expected, t)
		})
	}
}

func TestMarshal_GoogleProtobuf(t *testing.T) {
	tests := []struct {
		desc     string
		input    proto.Message
		option   LuaMarshalOption
		expected string
	}{
		{
			desc: "Message with one timestamp",
			input: &googleprotobufproto.SingleTimestamp{
				Fut: timestamppb.New(time.Date(2021, 8, 15, 0, 0, 0, 0, time.UTC)),
			},
			expected: "SingleTimestamp={fut={seconds=1628985600,nanos=0}}",
		},
		{
			desc: "Message with multiple timestamp",
			input: &googleprotobufproto.MultipleTimestamp{
				One: timestamppb.New(time.Date(2022, 2, 17, 0, 0, 0, 0, time.UTC)),
				Two: timestamppb.New(time.Date(2023, 10, 19, 0, 0, 0, 0, time.UTC)),
			},
			expected: "MultipleTimestamp={one={seconds=1645056000,nanos=0},two={seconds=1697673600,nanos=0}}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual, err := tt.option.Marshal(tt.input)
			AssertLuaString(actual, err, tt.expected, t)
		})
	}
}

func TestMarshal_GoogleProtobuf_CanBeOverridden(t *testing.T) {
	message := &googleprotobufproto.SingleTimestamp{
		Fut: timestamppb.New(time.Date(2021, 8, 15, 0, 0, 0, 0, time.UTC)),
	}
	expected := "SingleTimestamp={fut=\"FortyTwo\"}"

	actual, err := LuaMarshalOption{AdditionalMarshalers: []interface {
		Handle(fullName protoreflect.FullName) (MarshalFunc, error)
	}{&TestOverrider{}}}.marshal(message)

	AssertLuaString(actual, err, expected, t)
}

func AssertLuaString(actual []byte, err error, expected string, t *testing.T) {
	if err != nil {
		t.Errorf("Marshal() returned error: %v\n", err)
	}
	actualString := string(actual)
	if actualString != expected {
		t.Errorf("Marshal()\n<actual>\n%v\n<expected>\n%v\n", actualString, expected)
		if diff := cmp.Diff(expected, actualString); diff != "" {
			t.Errorf("Marshal() diff -expected +actual\n%v\n", diff)
		}
	}
}

type TestOverrider struct {
}

func (TestOverrider) Handle(fullName protoreflect.FullName) (MarshalFunc, error) {

	convertTimestamp := func(encodingRun EncodingRun, m protoreflect.Message) error {
		err := encodingRun.Encoder.WriteString("FortyTwo")
		if err != nil {
			return err
		}
		return nil
	}

	if fullName.Parent() == protoluainternal.GoogleProtobufParentPackage {
		switch fullName.Name() {
		case protoluainternal.GoogleProtobufTimestamp:
			return convertTimestamp, nil
		default:
			return nil, errors.New(string(fullName) + " is not supported yet")
		}
	}
	return nil, nil
}
