package protolua

import (
	"errors"

	internal "github.com/kinneko-de/protobuf-go/internal/encoding/protolua"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type GoogleWellKnownTypesMarshaler struct {
}

// GoogleWellKnownTypesMarshaler returns a marshal function for message of the package "google.protobuf"
// it returns an error in case of types of this packages are not supported yet.
// it returns nil if the package is different.
func (GoogleWellKnownTypesMarshaler) Handle(fullName protoreflect.FullName) (MarshalFunc, error) {
	if fullName.Parent() == internal.GoogleProtobufParentPackage {
		switch fullName.Name() {
		case internal.GoogleProtobufTimestamp:
			// Timestamp can be converted to normal table because it can be converted to os.date and os.time with build in function in lua
			return nil, nil
		default:
			return nil, errors.New(string(fullName) + " is not supported yet")
		}
	}
	return nil, nil
}
