package phoenix

// #include "go-clang.h"
import "C"

// The type of an element in the abstract syntax tree.
type Type struct {
	c C.CXType
}

// Pretty-print the underlying type using the rules of the language of the translation unit from which it came. If the type is invalid, an empty string is returned.
func (t Type) Spelling() string {
	o := cxstring{C.clang_getTypeSpelling(t.c)}
	defer o.Dispose()

	return o.String()
}

// Determine whether two CXTypes represent the same type. \returns non-zero if the CXTypes represent the same type and zero otherwise.
func EqualTypes(t1, t2 Type) bool {
	o := C.clang_equalTypes(t1.c, t2.c)

	return o != C.uint(0)
}

// Return the canonical type for a CXType. Clang's type system explicitly models typedefs and all the ways a specific type can be represented. The canonical type is the underlying type with all the "sugar" removed. For example, if 'T' is a typedef for 'int', the canonical type for 'T' would be 'int'.
func (t Type) CanonicalType() Type {
	return Type{C.clang_getCanonicalType(t.c)}
}

// Determine whether a CXType has the "const" qualifier set, without looking through typedefs that may have added "const" at a different level.
func (t Type) IsConstQualifiedType() bool {
	o := C.clang_isConstQualifiedType(t.c)

	return o != C.uint(0)
}

// Determine whether a CXType has the "volatile" qualifier set, without looking through typedefs that may have added "volatile" at a different level.
func (t Type) IsVolatileQualifiedType() bool {
	o := C.clang_isVolatileQualifiedType(t.c)

	return o != C.uint(0)
}

// Determine whether a CXType has the "restrict" qualifier set, without looking through typedefs that may have added "restrict" at a different level.
func (t Type) IsRestrictQualifiedType() bool {
	o := C.clang_isRestrictQualifiedType(t.c)

	return o != C.uint(0)
}

// For pointer types, returns the type of the pointee.
func (t Type) PointeeType() Type {
	return Type{C.clang_getPointeeType(t.c)}
}

// Return the cursor for the declaration of the given type.
func (t Type) Declaration() Cursor {
	return Cursor{C.clang_getTypeDeclaration(t.c)}
}

// Retrieve the calling convention associated with a function type. If a non-function type is passed in, CXCallingConv_Invalid is returned.
func (t Type) FunctionTypeCallingConv() CallingConv {
	return CallingConv(C.clang_getFunctionTypeCallingConv(t.c))
}

// Retrieve the result type associated with a function type. If a non-function type is passed in, an invalid type is returned.
func (t Type) ResultType() Type {
	return Type{C.clang_getResultType(t.c)}
}

// Retrieve the number of non-variadic arguments associated with a function type. If a non-function type is passed in, -1 is returned.
func (t Type) NumArgTypes() uint16 {
	return uint16(C.clang_getNumArgTypes(t.c))
}

// Return 1 if the CXType is a variadic function type, and 0 otherwise.
func (t Type) IsFunctionTypeVariadic() bool {
	o := C.clang_isFunctionTypeVariadic(t.c)

	return o != C.uint(0)
}

// Return 1 if the CXType is a POD (plain old data) type, and 0 otherwise.
func (t Type) IsPODType() bool {
	o := C.clang_isPODType(t.c)

	return o != C.uint(0)
}

// Return the element type of an array, complex, or vector type. If a type is passed in that is not an array, complex, or vector type, an invalid type is returned.
func (t Type) ElementType() Type {
	return Type{C.clang_getElementType(t.c)}
}

// Return the element type of an array type. If a non-array type is passed in, an invalid type is returned.
func (t Type) ArrayElementType() Type {
	return Type{C.clang_getArrayElementType(t.c)}
}

// Return the class type of an member pointer type. If a non-member-pointer type is passed in, an invalid type is returned.
func (t Type) ClassType() Type {
	return Type{C.clang_Type_getClassType(t.c)}
}

// Retrieve the ref-qualifier kind of a function or method. The ref-qualifier is returned for C++ functions or methods. For other types or non-C++ declarations, CXRefQualifier_None is returned.
func (t Type) RefQualifier() RefQualifierKind {
	return RefQualifierKind(C.clang_Type_getCXXRefQualifier(t.c))
}
