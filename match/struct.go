package match

import (
	"fmt"
	"reflect"

	"github.com/krelinga/go-deep"
)

type Struct map[string]Matcher

func (s Struct) Match(env deep.Env, vals Vals) Result {
	got := vals.Want1()
	if got.Kind() != reflect.Struct {
		panic(fmt.Errorf("%w: value is not a struct: got %s", ErrBadType, got.Kind()))
	}
	for fieldName, matcher := range s {
		t := got.Type()
		field, ok := t.FieldByName(fieldName)
		if !ok {
			panic(fmt.Errorf("%w: struct does not have field %q", ErrBadType, fieldName))
		}
		fieldVal := got.FieldByIndex(field.Index)
		result := matcher.Match(env, Vals{fieldVal})
		if !result.Matched() {
			return NewResult(false, fmt.Sprintf("field %q does not match", fieldName))
		}
	}
	return NewResult(true, "struct matches")
}

type Struct2 struct {
	Default Matcher
	Fields  map[deep.Field]Matcher
}
