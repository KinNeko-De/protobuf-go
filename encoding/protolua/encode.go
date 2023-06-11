package protolua

import (
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Marshal converts the given proto.Message into a lua table using default options.
// Supports only proto3 messages currently.
func Marshal(m proto.Message) ([]byte, error) {
	return LuaMarshalOption{}.Marshal(m)
}

type (
	EncodingRun struct {
		*Encoder
		options LuaMarshalOption
	}

	MarshalFunc func(EncodingRun, protoreflect.Message) error

	LuaMarshalOption struct {
		Format struct {
			// If set to false the output is formated as one line (default)
			// If set to true the output is formated in multiple line with indent (better for humans)
			Multiline bool
		}

		// Defines how the name of message are crated
		// If set to nil, [protolua.JsonName] will be used
		KeyName interface {
			keyName(protoreflect.FieldDescriptor) string
		}

		// Additional Marshalers for non standard proto messages
		// see [protolua.GoogleWellKnownTypesMarshaler] for an example
		AdditionalMarshalers []interface {
			Handle(fullName protoreflect.FullName) (MarshalFunc, error)
		}
	}
)

// Marshal convert the given proto.Message into a lua table using the given options.
func (o LuaMarshalOption) Marshal(m proto.Message) ([]byte, error) {
	return o.marshal(m)
}

func (option LuaMarshalOption) marshal(m proto.Message) ([]byte, error) {
	if m == nil {
		return nil, errors.New("message can not be nil")
	}

	setDefaults(&option)
	indent := ""
	if option.Format.Multiline {
		indent = DefaultIndent
	}

	encoder, err := NewEncoder(indent)
	if err != nil {
		return nil, err
	}

	encodingRun := EncodingRun{encoder, option}

	bytes, err2 := marshalRootMessage(m.ProtoReflect(), encodingRun)
	if err2 != nil {
		return bytes, err2
	}

	return encodingRun.Encoder.Bytes(), nil
}

func setDefaults(option *LuaMarshalOption) {
	if option.KeyName == nil {
		option.KeyName = JsonName{}
	}
}

func marshalRootMessage(m protoreflect.Message, encodingRun EncodingRun) ([]byte, error) {
	// The json name is not populated, so the Protobuf name is used hereX
	encodingRun.Encoder.WriteKey(string(m.Descriptor().Name()))

	if err := encodingRun.marshalMessage(m); err != nil {
		return nil, err
	}
	return nil, nil
}

// marshalMessage marshals the message and fields in the given protoreflect.Message.
func (e EncodingRun) marshalMessage(m protoreflect.Message) error {
	if m.Descriptor().Syntax() != protoreflect.Proto3 {
		return errors.New("only proto3 syntax is supported")
	}

	for _, marshaler := range e.options.AdditionalMarshalers {
		marshalFunc, unsupportedTypeError := marshaler.Handle(m.Descriptor().FullName())
		shouldReturn, returnValue := isHandledByOtherMarshaler(unsupportedTypeError, marshalFunc, e, m)
		if shouldReturn {
			return returnValue
		}
	}

	marshalFunc, unsupportedTypeError := GoogleWellKnownTypesMarshaler{}.Handle(m.Descriptor().FullName())
	shouldReturn, returnValue := isHandledByOtherMarshaler(unsupportedTypeError, marshalFunc, e, m)
	if shouldReturn {
		return returnValue
	}

	if !m.IsValid() {
		e.Encoder.WriteNull()
		return nil
	}

	var err error
	e.Encoder.StartObject()
	defer e.Encoder.EndObject()

	fields := m.Descriptor().Fields()
	upper := fields.Len()
	for i := 0; i < upper; i++ {
		currentField := fields.Get(i)
		name := e.options.KeyName.keyName(currentField)

		if err = e.Encoder.WriteKey(name); err != nil {
			return err
		}
		if err = e.marshalValue(m.Get(currentField), currentField); err != nil {
			return err
		}
	}

	return err
}

func isHandledByOtherMarshaler(unsupportedTypeError error, marshalFunc MarshalFunc, e EncodingRun, m protoreflect.Message) (bool, error) {
	if unsupportedTypeError != nil {
		return true, unsupportedTypeError
	}
	if marshalFunc != nil {
		return true, marshalFunc(e, m)
	}
	return false, nil
}

func (e EncodingRun) marshalValue(val protoreflect.Value, fd protoreflect.FieldDescriptor) error {
	switch {
	case fd.IsList():
		return e.marshalList(val.List(), fd)
	case fd.IsMap():
		return e.marshalMap(val.Map(), fd)
	default:
		return e.marshalSingular(val, fd)
	}
}

func (e EncodingRun) marshalSingular(val protoreflect.Value, fd protoreflect.FieldDescriptor) error {
	if !val.IsValid() {
		return nil
	}

	switch kind := fd.Kind(); kind {
	case protoreflect.BoolKind:
		e.WriteBool(val.Bool())

	case protoreflect.StringKind:
		if err := e.Encoder.WriteString(val.String()); err != nil {
			return err
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind,
		protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind:
		e.Encoder.WriteNumber(val.String())

	case protoreflect.FloatKind:
		return errors.New("float is not supported yet")

	case protoreflect.DoubleKind:
		return errors.New("float is not supported yet")

	case protoreflect.BytesKind:
		return errors.New("byte is not supported yet")

	case protoreflect.EnumKind:
		return errors.New("enum is not supported yet")

	case protoreflect.MessageKind, protoreflect.GroupKind:
		if err := e.marshalMessage(val.Message()); err != nil {
			return err
		}

	default:
		panic(fmt.Sprintf("%v has unknown kind: %v", fd.FullName(), kind))
	}
	return nil
}

func (e EncodingRun) marshalList(list protoreflect.List, fd protoreflect.FieldDescriptor) error {

	e.Encoder.StartArray()
	defer e.Encoder.EndArray()

	for i := 0; i < list.Len(); i++ {
		item := list.Get(i)
		e.Encoder.WriteIndexedList(i + 1)
		if err := e.marshalSingular(item, fd); err != nil {
			return err
		}
	}
	return nil
}

func (e EncodingRun) marshalMap(mmap protoreflect.Map, fd protoreflect.FieldDescriptor) error {
	return errors.New("maps are not supported yet")
}

