package protolua

import "google.golang.org/protobuf/reflect/protoreflect"

// default for KeyName, will use the jsonName for keys. Except for the root messages where the json name can not be extracted.
type JsonName struct {
}

func (j JsonName) keyName(fieldDescriptor protoreflect.FieldDescriptor) string {
	return fieldDescriptor.JSONName()
}

// will use protobufName for keys
type ProtobufName struct {
}

func (p ProtobufName) keyName(fieldDescriptor protoreflect.FieldDescriptor) string {
	return fieldDescriptor.TextName()
}
