//! This module defines the types that can be used in the schema at a type-level.
//!
//! Unless you are working with the implementation details of schema encoding, then you
//! should consider this module as something that ensures type safety.
//! This module uses a programming style known as type-level programming where types
//! are defined to build other types.
//! None of the types in this module are expected to be instantiated other than as type-level
//! parameters.

/// The `Type` trait is implemented for all types that can be used in the schema.
pub trait Type: Private {}
trait Private {}

/// The `U8T` type represents an unsigned 8-bit integer.
pub struct U8T;
impl Private for U8T {}
impl Type for U8T {}
impl ListElementType for U8T {}

/// The `U16T` type represents an unsigned 16-bit integer.
pub struct U16T;
impl Private for U16T {}
impl Type for U16T {}
impl ListElementType for U16T {}

/// The `U32` type represents an unsigned 32-bit integer.
pub struct U32T;
impl Private for U32T {}
impl Type for U32T {}
impl ListElementType for U32T {}

/// The `U64T` type represents an unsigned 64-bit integer.
pub struct U64T;
impl Private for U64T {}
impl Type for U64T {}

/// The `UIntNT` type represents an unsigned N-bit integer.
pub struct UIntNT<const N: usize> {}
impl<const N: usize> Private for UIntNT<N> {}
impl<const N: usize> Type for UIntNT<N> {}

/// The `I8T` type represents a signed 8-bit integer.
pub struct I8T;
impl Private for I8T {}
impl Type for I8T {}

/// The `I16T` type represents a signed 16-bit integer.
pub struct I16T;
impl Private for I16T {}
impl Type for I16T {}

/// The `I32T` type represents a signed 32-bit integer.
pub struct I32T;
impl Private for I32T {}
impl Type for I32T {}

/// The `I64T` type represents a signed 64-bit integer.
pub struct I64T;
impl Private for I64T {}
impl Type for I64T {}

/// The `IntNT` type represents a signed integer represented by N bytes (not bits).
pub struct IntNT<const N: usize>;
impl<const N: usize> Private for IntNT<N> {}
impl<const N: usize> Type for IntNT<N> {}

/// The `Bool` type represents a boolean.
pub struct Bool;
impl Private for Bool {}
impl Type for Bool {}

/// The `StrT` type represents a string.
pub struct StrT;
impl Private for StrT {}
impl Type for StrT {}

/// The `AddressT` type represents an address.
pub struct AddressT;
impl Private for AddressT {}
impl Type for AddressT {}

/// The `TimeT` type represents a time.
pub struct TimeT;
impl Private for TimeT {}
impl Type for TimeT {}

/// The `DurationT` type represents a duration.
pub struct DurationT;
impl Private for DurationT {}
impl Type for DurationT {}

/// The `NullableT` type represents a nullable type.
pub struct NullableT<T> {
    _phantom: std::marker::PhantomData<T>,
}
impl <T> Private for NullableT<T> {}
impl <T> Type for NullableT<T> {}

/// The `ListT` type represents a list type.
pub struct ListT<T: ListElementType> {
    _phantom: std::marker::PhantomData<T>,
}
impl <T:ListElementType> Private for ListT<T> {}
impl <T:ListElementType> Type for ListT<T> {}

/// The `StructT` type represents a struct type.
pub struct StructT<T> {
    _phantom: std::marker::PhantomData<T>,
}
impl <T> Private for StructT<T> {}
impl <T> Type for StructT<T> {}

/// Represents a type that can be used as an element in a list.
pub trait ListElementType {}