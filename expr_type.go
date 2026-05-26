package actionlint

// Types

// ExprType is interface for types of values in expression.
type ExprType interface {
	// String returns string representation of the type.
	String() string
	// Assignable returns if other type can be assignable to the type.
	Assignable(other ExprType) bool
	// Merge merges other type into this type. When other type conflicts with this type, the merged
	// result is any type as fallback.
	Merge(other ExprType) ExprType
	// DeepCopy duplicates itself. All its child types are copied recursively.
	DeepCopy() ExprType
}

// AnyType represents type which can be any type. It also indicates that a value of the type cannot
// be type-checked since it's type cannot be known statically.
type AnyType struct{}

func (ty AnyType) String() string {
	_ = "STUB: not implemented"

	// Assignable returns if other type can be assignable to the type.
	return ""
}

func (ty AnyType) Assignable(_ ExprType) bool {
	_ = "STUB: not implemented"

	// Merge merges other type into this type. When other type conflicts with this type, the merged
	// result is any type as fallback.
	return false
}

func (ty AnyType) Merge(other ExprType) ExprType {
	_ = "STUB: not implemented"

	// DeepCopy duplicates itself. All its child types are copied recursively.
	return *new(ExprType)
}

func (ty AnyType) DeepCopy() ExprType {
	_ = "STUB: not implemented"

	// NullType is type for null value.
	return *new(ExprType)
}

type NullType struct{}

func (ty NullType) String() string {
	_ = "STUB: not implemented"

	// Assignable returns if other type can be assignable to the type.
	return ""
}

func (ty NullType) Assignable(other ExprType) bool { _ = "STUB: not implemented"; return false }

// Merge merges other type into this type. When other type conflicts with this type, the merged
// result is any type as fallback.
func (ty NullType) Merge(other ExprType) ExprType { _ = "STUB: not implemented"; return *new(ExprType) }

// DeepCopy duplicates itself. All its child types are copied recursively.
func (ty NullType) DeepCopy() ExprType {
	_ = "STUB: not implemented"

	// NumberType is type for number values such as integer or float.
	return *new(ExprType)
}

type NumberType struct{}

func (ty NumberType) String() string {
	_ = "STUB: not implemented"

	// Assignable returns if other type can be assignable to the type.
	return ""
}

func (ty NumberType) Assignable(other ExprType) bool {
	_ = "STUB: not implemented"
	// TODO: Is string of numbers corced into number?
	return false
}

// Merge merges other type into this type. When other type conflicts with this type, the merged
// result is any type as fallback.
func (ty NumberType) Merge(other ExprType) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// DeepCopy duplicates itself. All its child types are copied recursively.
func (ty NumberType) DeepCopy() ExprType {
	_ = "STUB: not implemented"

	// BoolType is type for boolean values.
	return *new(ExprType)
}

type BoolType struct{}

func (ty BoolType) String() string {
	_ = "STUB: not implemented"

	// Assignable returns if other type can be assignable to the type.
	return ""
}

func (ty BoolType) Assignable(other ExprType) bool {
	_ = "STUB: not implemented"
	// Any type can be converted into bool..
	// e.g.
	//
	//	if: ${{ steps.foo }}
	return false
}

// Merge merges other type into this type. When other type conflicts with this type, the merged
// result is any type as fallback.
func (ty BoolType) Merge(other ExprType) ExprType { _ = "STUB: not implemented"; return *new(ExprType) }

// DeepCopy duplicates itself. All its child types are copied recursively.
func (ty BoolType) DeepCopy() ExprType {
	_ = "STUB: not implemented"

	// StringType is type for string values.
	return *new(ExprType)
}

type StringType struct{}

func (ty StringType) String() string {
	_ = "STUB: not implemented"

	// Assignable returns if other type can be assignable to the type.
	return ""
}

func (ty StringType) Assignable(other ExprType) bool {
	_ = "STUB: not implemented"
	// Bool and null types also can be coerced into string. But in almost all case, those coercing
	// would be mistakes.
	return false
}

// Merge merges other type into this type. When other type conflicts with this type, the merged
// result is any type as fallback.
func (ty StringType) Merge(other ExprType) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// DeepCopy duplicates itself. All its child types are copied recursively.
func (ty StringType) DeepCopy() ExprType {
	_ = "STUB: not implemented"

	// ObjectType is type for objects, which can hold key-values.
	return *new(ExprType)
}

type ObjectType struct {
	// Props is map from properties name to their type.
	Props map[string]ExprType
	// Mapped is an element type of this object. This means all props have the type. For example,
	// The element type of env context is string.
	// AnyType means its property types can be any type so it shapes a loose object. Setting nil
	// means properties are mapped to no type so it shapes a strict object.
	//
	// Invariant: All types in Props field must be assignable to this type.
	Mapped ExprType
}

// NewEmptyObjectType creates new loose ObjectType instance which allows unknown props. When
// accessing to unknown props, their values will fall back to any.
func NewEmptyObjectType() *ObjectType { _ = "STUB: not implemented"; return nil }

// NewObjectType creates new loose ObjectType instance which allows unknown props with given props.
func NewObjectType(props map[string]ExprType) *ObjectType { _ = "STUB: not implemented"; return nil }

// NewEmptyStrictObjectType creates new ObjectType instance which does not allow unknown props.
func NewEmptyStrictObjectType() *ObjectType { _ = "STUB: not implemented"; return nil }

// NewStrictObjectType creates new ObjectType instance which does not allow unknown props with
// given prop types.
func NewStrictObjectType(props map[string]ExprType) *ObjectType {
	_ = "STUB: not implemented"
	return nil
}

// NewMapObjectType creates new ObjectType which maps keys to a specific type value.
func NewMapObjectType(t ExprType) *ObjectType { _ = "STUB: not implemented"; return nil }

// IsStrict returns if the type is a strict object, which means no unknown prop is allowed.
func (ty *ObjectType) IsStrict() bool { _ = "STUB: not implemented"; return false }

// IsLoose returns if the type is a loose object, which allows any unknown props.
func (ty *ObjectType) IsLoose() bool { _ = "STUB: not implemented"; return false }

// Strict sets the object is strict, which means only known properties are allowed.
func (ty *ObjectType) Strict() {
	_ = "STUB: not implemented"

	// Loose sets the object is loose, which means any properties can be set.
	return
}

func (ty *ObjectType) Loose() { _ = "STUB: not implemented"; return }

func (ty *ObjectType) String() string { _ = "STUB: not implemented"; return "" }

// Assignable returns if other type can be assignable to the type.
// In other words, rhs type is more strict than lhs (receiver) type.
func (ty *ObjectType) Assignable(other ExprType) bool { _ = "STUB: not implemented"; return false }

// ty is strict

// ty and other are strict

// Merge merges two object types into one. When other object has unknown props, they are merged into
// current object. When both have same property, when they are assignable, it remains as-is.
// Otherwise, the property falls back to any type.
func (ty *ObjectType) Merge(other ExprType) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// Shortcuts

// DeepCopy duplicates itself. All its child types are copied recursively.
func (ty *ObjectType) DeepCopy() ExprType { _ = "STUB: not implemented"; return *new(ExprType) }

// ArrayType is type for arrays.
type ArrayType struct {
	// Elem is type of element of the array.
	Elem ExprType
	// Deref is true when this type was derived from object filtering syntax (foo.*).
	Deref bool
}

func (ty *ArrayType) String() string { _ = "STUB: not implemented"; return "" }

// Assignable returns if other type can be assignable to the type.
func (ty *ArrayType) Assignable(other ExprType) bool { _ = "STUB: not implemented"; return false }

// Merge merges two object types into one. When other object has unknown props, they are merged into
// current object. When both have same property, when they are assignable, it remains as-is.
// Otherwise, the property falls back to any type.
func (ty *ArrayType) Merge(other ExprType) ExprType {
	_ = "STUB: not implemented"
	return *new(ExprType)
}

// When fusing array deref type, it means prop deref chain breaks

// DeepCopy duplicates itself. All its child types are copied recursively.
func (ty *ArrayType) DeepCopy() ExprType { _ = "STUB: not implemented"; return *new(ExprType) }

// EqualTypes returns if the two types are equal.
func EqualTypes(l, r ExprType) bool { _ = "STUB: not implemented"; return false }

// typeOfJSONValue returns the type of the given JSON value. The JSON value is an any value decoded by json.Unmarshal.
// https://pkg.go.dev/encoding/json#Unmarshal
//
// To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:
//   - bool, for JSON booleans
//   - float64, for JSON numbers
//   - string, for JSON strings
//   - []interface{}, for JSON arrays
//   - map[string]interface{}, for JSON objects
//   - nil for JSON null
func typeOfJSONValue(v any) ExprType { _ = "STUB: not implemented"; return *new(ExprType) }

// Unreachable
