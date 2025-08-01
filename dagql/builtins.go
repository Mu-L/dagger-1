package dagql

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"

	"github.com/dagger/dagger/dagql/call"
)

func builtinOrTyped(val any) (Typed, error) {
	switch x := val.(type) {
	case Typed:
		return x, nil
	case string:
		return String(x), nil
	case int:
		return Int(x), nil
	case int32:
		return Int(x), nil
	case int64:
		return Int(x), nil
	case float32:
		return Float(x), nil
	case float64:
		return Float(x), nil
	case bool:
		return Boolean(x), nil
	default:
		valT := reflect.TypeOf(val)
		valV := reflect.ValueOf(val)
		if valT == nil {
			return nil, fmt.Errorf("cannot convert %T to a Typed value", val)
		}
		switch valT.Kind() {
		case reflect.Slice:
			elem, err := builtinOrTyped(reflect.New(valT.Elem()).Elem().Interface())
			if err != nil {
				return nil, fmt.Errorf("slice elem: %w", err)
			}
			arr := DynamicArrayOutput{
				Elem: elem,
			}
			for i := range valV.Len() {
				elem, err := builtinOrTyped(valV.Index(i).Interface())
				if err != nil {
					return nil, fmt.Errorf("slice elem %d: %w", i, err)
				}
				arr.Values = append(arr.Values, elem)
			}
			return arr, nil
		case reflect.Ptr:
			elem, err := builtinOrTyped(reflect.New(valT.Elem()).Elem().Interface())
			if err != nil {
				return nil, fmt.Errorf("slice elem: %w", err)
			}
			nul := DynamicNullable{
				Elem: elem,
			}
			if !valV.IsNil() {
				elem, err := builtinOrTyped(valV.Elem().Interface())
				if err != nil {
					return nil, fmt.Errorf("slice elem: %w", err)
				}
				nul.Value = elem
				nul.Valid = true
			}
			return nul, nil
		default:
			return nil, fmt.Errorf("cannot convert %T to a Typed value", val)
		}
	}
}

type DynamicArrayOutput struct {
	Elem   Typed
	Values []Typed
}

var _ Typed = DynamicArrayOutput{}

func (d DynamicArrayOutput) Type() *ast.Type {
	return &ast.Type{
		Elem:    d.Elem.Type(),
		NonNull: true,
	}
}

var _ Enumerable = DynamicArrayOutput{}

func (d DynamicArrayOutput) Element() Typed {
	return d.Elem
}

func (d DynamicArrayOutput) Len() int {
	return len(d.Values)
}

func (d DynamicArrayOutput) Nth(i int) (Typed, error) {
	if i < 1 || i > len(d.Values) {
		return nil, fmt.Errorf("index %d out of bounds", i)
	}
	return d.Values[i-1], nil
}

func (d DynamicArrayOutput) NthValue(i int, enumID *call.ID) (AnyResult, error) {
	t, err := d.Nth(i)
	if err != nil {
		return nil, err
	}
	return Result[Typed]{
		constructor: enumID.SelectNth(i),
		self:        t,
	}, nil
}

func (d DynamicArrayOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Values)
}

func (d DynamicArrayOutput) SetField(val reflect.Value) error {
	if val.Kind() != reflect.Slice {
		return fmt.Errorf("expected slice, got %v", val.Kind())
	}
	val.Set(reflect.MakeSlice(val.Type(), len(d.Values), len(d.Values)))
	for i, elem := range d.Values {
		if err := assign(val.Index(i), elem); err != nil {
			return err
		}
	}
	return nil
}

type DynamicResultArrayOutput struct {
	Elem   Typed
	Values []AnyResult
}

var _ Typed = DynamicResultArrayOutput{}

func (d DynamicResultArrayOutput) Type() *ast.Type {
	return &ast.Type{
		Elem:    d.Elem.Type(),
		NonNull: true,
	}
}

var _ Enumerable = DynamicResultArrayOutput{}

func (d DynamicResultArrayOutput) Element() Typed {
	return d.Elem
}

func (d DynamicResultArrayOutput) Len() int {
	return len(d.Values)
}

func (d DynamicResultArrayOutput) Nth(i int) (Typed, error) {
	val, err := d.NthValue(i, nil)
	if err != nil {
		return nil, err
	}
	return val.Unwrap(), nil
}

func (d DynamicResultArrayOutput) NthValue(i int, _ *call.ID) (AnyResult, error) {
	if i < 1 || i > len(d.Values) {
		return nil, fmt.Errorf("index %d out of bounds", i)
	}
	return d.Values[i-1], nil
}

func (d DynamicResultArrayOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Values)
}

func (d DynamicResultArrayOutput) SetField(val reflect.Value) error {
	if val.Kind() != reflect.Slice {
		return fmt.Errorf("expected slice, got %v", val.Kind())
	}
	val.Set(reflect.MakeSlice(val.Type(), len(d.Values), len(d.Values)))
	for i, elem := range d.Values {
		if err := assign(val.Index(i), elem); err != nil {
			return err
		}
	}
	return nil
}

func builtinOrInput(val any) (Input, error) {
	switch x := val.(type) {
	case Input:
		return x, nil
	case string:
		return String(x), nil
	case int:
		return Int(x), nil
	case int32:
		return Int(x), nil
	case int64:
		return Int(x), nil
	case float32:
		return Float(x), nil
	case float64:
		return Float(x), nil
	case bool:
		return Boolean(x), nil
	default:
		valT := reflect.TypeOf(val)
		if val == nil {
			return nil, fmt.Errorf("cannot convert nil to an Input value")
		}
		reflectVal := reflect.ValueOf(val)
		switch valT.Kind() {
		case reflect.Slice:
			input, err := builtinOrInput(reflect.New(valT.Elem()).Elem().Interface())
			if err != nil {
				return nil, fmt.Errorf("slice elem: %w", err)
			}
			arr := DynamicArrayInput{
				Elem: input,
			}
			for i := range reflectVal.Len() {
				elem, err := builtinOrInput(reflectVal.Index(i).Interface())
				if err != nil {
					return nil, fmt.Errorf("slice elem val %d: %w", i, err)
				}
				arr.Values = append(arr.Values, elem)
			}
			return arr, nil
		case reflect.Ptr:
			input, err := builtinOrInput(reflect.New(valT.Elem()).Elem().Interface())
			if err != nil {
				return nil, fmt.Errorf("pointer elem: %w", err)
			}
			dynOpt := DynamicOptional{
				Elem: input,
			}
			if !reflectVal.IsNil() {
				dynOpt.Value, err = dynOpt.Elem.Decoder().DecodeInput(reflectVal.Elem().Interface())
				if err != nil {
					return nil, fmt.Errorf("pointer elem val: %w", err)
				}
				dynOpt.Valid = true
			}
			return dynOpt, nil
		default:
			return nil, fmt.Errorf("cannot convert %T to an Input value", val)
		}
	}
}

type DynamicArrayInput struct {
	Elem   Input
	Values []Input
}

var _ InputDecoder = DynamicArrayInput{}

func (d DynamicArrayInput) DecodeInput(val any) (Input, error) {
	switch x := val.(type) {
	case []any:
		arr := DynamicArrayInput{
			Elem: d.Elem,
		}
		decoder := d.Elem.Decoder()
		for _, elem := range x {
			decoded, err := decoder.DecodeInput(elem)
			if err != nil {
				return nil, err
			}
			arr.Values = append(arr.Values, decoded)
		}
		return arr, nil
	case string: // default
		var vals []any
		dec := json.NewDecoder(strings.NewReader(x))
		dec.UseNumber()
		if err := dec.Decode(&vals); err != nil {
			return nil, fmt.Errorf("decode %q: %w", x, err)
		}
		return d.DecodeInput(vals)
	default:
		return nil, fmt.Errorf("expected array, got %T", val)
	}
}

var _ Input = DynamicArrayInput{}

func (d DynamicArrayInput) ToLiteral() call.Literal {
	literals := make([]call.Literal, 0, len(d.Values))
	for _, elem := range d.Values {
		literals = append(literals, elem.ToLiteral())
	}
	return call.NewLiteralList(literals...)
}

func (d DynamicArrayInput) Type() *ast.Type {
	return &ast.Type{
		Elem:    d.Elem.Type(),
		NonNull: true,
	}
}

func (d DynamicArrayInput) Decoder() InputDecoder {
	return DynamicArrayInput{
		Elem: d.Elem,
	}
}

var _ Setter = DynamicArrayInput{}

func (d DynamicArrayInput) SetField(val reflect.Value) error {
	if val.Kind() != reflect.Slice {
		return fmt.Errorf("expected slice, got %v", val.Kind())
	}
	val.Set(reflect.MakeSlice(val.Type(), len(d.Values), len(d.Values)))
	for i, elem := range d.Values {
		if err := assign(val.Index(i), elem); err != nil {
			return err
		}
	}
	return nil
}

var _ Enumerable = DynamicArrayInput{}

func (d DynamicArrayInput) Element() Typed {
	return d.Elem
}

func (d DynamicArrayInput) Len() int {
	return len(d.Values)
}

func (d DynamicArrayInput) Nth(i int) (Typed, error) {
	if i < 1 || i > len(d.Values) {
		return nil, fmt.Errorf("index %d out of bounds", i)
	}
	return d.Values[i-1], nil
}

func (d DynamicArrayInput) NthValue(i int, enumID *call.ID) (AnyResult, error) {
	t, err := d.Nth(i)
	if err != nil {
		return nil, err
	}
	return Result[Typed]{
		constructor: enumID.SelectNth(i),
		self:        t,
	}, nil
}
