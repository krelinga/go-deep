package deep

import (
	"fmt"
	"reflect"
)

type Field interface {
	GetField(t reflect.Type) reflect.StructField
	fieldIsAClosedType()
}

type namedField string

func (nf namedField) GetField(t reflect.Type) reflect.StructField {
	field, ok := t.FieldByName(string(nf))
	if !ok {
		panic(fmt.Errorf("%w: struct %s does not have field %q", ErrWrongType, t, string(nf)))
	}
	return field
}

func (nf namedField) fieldIsAClosedType() {}

func NamedField(name string) Field {
	return namedField(name)
}

func EmbedTypeField(typ reflect.Type) Field {
	return embedField{Typ: typ}
}

func EmbedField[T any]() Field {
	return EmbedTypeField(reflect.TypeFor[T]())
}

type embedField struct {
	Typ reflect.Type
}

func (ef embedField) GetField(t reflect.Type) reflect.StructField {
	for i := range t.NumField() {
		field := t.Field(i)
		if field.Type == ef.Typ && field.Anonymous {
			return field
		}
	}
	panic(fmt.Errorf("%w: struct %s does not have embedded field of type %s", ErrWrongType, t, ef.Typ))
}

func (ef embedField) fieldIsAClosedType() {}
