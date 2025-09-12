package types

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

// Type represents a database data type
type Type interface {
	// Size returns the size in bytes for fixed-size types, or -1 for variable-size
	Size() int
	
	// Encode converts a value to bytes
	Encode(value Value) []byte
	
	// Decode converts bytes to a value
	Decode(data []byte) Value
	
	// Compare compares two values of this type
	Compare(a, b Value) int
	
	// Validate validates a value against this type
	Validate(value Value) error
	
	// String returns the string representation of this type
	String() string
}

// Value represents a database value
type Value interface {
	// Type returns the type of this value
	Type() Type
	
	// String returns the string representation
	String() string
	
	// IsNull returns true if this is a null value
	IsNull() bool
}

// NullValue represents a null value
type NullValue struct{}

func (n NullValue) Type() Type {
	return &NullType{}
}

func (n NullValue) String() string {
	return "NULL"
}

func (n NullValue) IsNull() bool {
	return true
}

// Concrete value types
type IntValue int64
type StringValue string
type FloatValue float64
type BoolValue bool
type DateTimeValue time.Time
type BytesValue []byte

// IntValue methods
func (v IntValue) Type() Type {
	return &IntegerType{}
}

func (v IntValue) String() string {
	return strconv.FormatInt(int64(v), 10)
}

func (v IntValue) IsNull() bool {
	return false
}

// StringValue methods
func (v StringValue) Type() Type {
	return &StringType{}
}

func (v StringValue) String() string {
	return string(v)
}

func (v StringValue) IsNull() bool {
	return false
}

// FloatValue methods
func (v FloatValue) Type() Type {
	return &FloatType{}
}

func (v FloatValue) String() string {
	return strconv.FormatFloat(float64(v), 'f', -1, 64)
}

func (v FloatValue) IsNull() bool {
	return false
}

// BoolValue methods
func (v BoolValue) Type() Type {
	return &BoolType{}
}

func (v BoolValue) String() string {
	return strconv.FormatBool(bool(v))
}

func (v BoolValue) IsNull() bool {
	return false
}

// DateTimeValue methods
func (v DateTimeValue) Type() Type {
	return &DateTimeType{}
}

func (v DateTimeValue) String() string {
	return time.Time(v).Format(time.RFC3339)
}

func (v DateTimeValue) IsNull() bool {
	return false
}

// BytesValue methods
func (v BytesValue) Type() Type {
	return &BytesType{}
}

func (v BytesValue) String() string {
	return fmt.Sprintf("0x%x", []byte(v))
}

func (v BytesValue) IsNull() bool {
	return false
}

// Type implementations

// NullType represents a null type
type NullType struct{}

func (t *NullType) Size() int {
	return 0
}

func (t *NullType) Encode(value Value) []byte {
	return []byte{}
}

func (t *NullType) Decode(data []byte) Value {
	return NullValue{}
}

func (t *NullType) Compare(a, b Value) int {
	// Null values are equal to each other
	return 0
}

func (t *NullType) Validate(value Value) error {
	if !value.IsNull() {
		return fmt.Errorf("expected null value")
	}
	return nil
}

func (t *NullType) String() string {
	return "NULL"
}

// IntegerType represents an integer type
type IntegerType struct {
	ByteSize int // 1, 2, 4, or 8 bytes
}

func (t *IntegerType) Size() int {
	return t.ByteSize
}

func (t *IntegerType) Encode(value Value) []byte {
	if value.IsNull() {
		return make([]byte, t.ByteSize)
	}
	
	intVal := int64(value.(IntValue))
	data := make([]byte, t.ByteSize)
	
	switch t.ByteSize {
	case 1:
		data[0] = byte(intVal)
	case 2:
		binary.BigEndian.PutUint16(data, uint16(intVal))
	case 4:
		binary.BigEndian.PutUint32(data, uint32(intVal))
	case 8:
		binary.BigEndian.PutUint64(data, uint64(intVal))
	default:
		panic(fmt.Sprintf("unsupported integer size: %d", t.ByteSize))
	}
	
	return data
}

func (t *IntegerType) Decode(data []byte) Value {
	if len(data) == 0 {
		return NullValue{}
	}
	
	var intVal int64
	
	switch t.ByteSize {
	case 1:
		intVal = int64(data[0])
	case 2:
		intVal = int64(binary.BigEndian.Uint16(data))
	case 4:
		intVal = int64(binary.BigEndian.Uint32(data))
	case 8:
		intVal = int64(binary.BigEndian.Uint64(data))
	default:
		panic(fmt.Sprintf("unsupported integer size: %d", t.ByteSize))
	}
	
	return IntValue(intVal)
}

func (t *IntegerType) Compare(a, b Value) int {
	if a.IsNull() && b.IsNull() {
		return 0
	}
	if a.IsNull() {
		return -1
	}
	if b.IsNull() {
		return 1
	}
	
	aVal := int64(a.(IntValue))
	bVal := int64(b.(IntValue))
	
	if aVal < bVal {
		return -1
	} else if aVal > bVal {
		return 1
	}
	return 0
}

func (t *IntegerType) Validate(value Value) error {
	if value.IsNull() {
		return nil
	}
	
	_, ok := value.(IntValue)
	if !ok {
		return fmt.Errorf("expected integer value")
	}
	
	return nil
}

func (t *IntegerType) String() string {
	return fmt.Sprintf("INT%d", t.ByteSize*8)
}

// StringType represents a string type
type StringType struct {
	MaxLength int
	Variable  bool
}

func (t *StringType) Size() int {
	if t.Variable {
		return -1 // Variable length
	}
	return t.MaxLength
}

func (t *StringType) Encode(value Value) []byte {
	if value.IsNull() {
		return []byte{}
	}
	
	strVal := string(value.(StringValue))
	
	if !t.Variable {
		// Fixed length - pad with nulls
		data := make([]byte, t.MaxLength)
		copy(data, strVal)
		return data
	} else {
		// Variable length - prefix with length
		length := len(strVal)
		data := make([]byte, 4+length)
		binary.BigEndian.PutUint32(data, uint32(length))
		copy(data[4:], strVal)
		return data
	}
}

func (t *StringType) Decode(data []byte) Value {
	if len(data) == 0 {
		return NullValue{}
	}
	
	if !t.Variable {
		// Fixed length - trim nulls
		str := string(data)
		// Find first null byte
		for i, b := range data {
			if b == 0 {
				str = string(data[:i])
				break
			}
		}
		return StringValue(str)
	} else {
		// Variable length - read length prefix
		if len(data) < 4 {
			return NullValue{}
		}
		
		length := int(binary.BigEndian.Uint32(data))
		if len(data) < 4+length {
			return NullValue{}
		}
		
		str := string(data[4 : 4+length])
		return StringValue(str)
	}
}

func (t *StringType) Compare(a, b Value) int {
	if a.IsNull() && b.IsNull() {
		return 0
	}
	if a.IsNull() {
		return -1
	}
	if b.IsNull() {
		return 1
	}
	
	aVal := string(a.(StringValue))
	bVal := string(b.(StringValue))
	
	if aVal < bVal {
		return -1
	} else if aVal > bVal {
		return 1
	}
	return 0
}

func (t *StringType) Validate(value Value) error {
	if value.IsNull() {
		return nil
	}
	
	strVal, ok := value.(StringValue)
	if !ok {
		return fmt.Errorf("expected string value")
	}
	
	if len(string(strVal)) > t.MaxLength {
		return fmt.Errorf("string too long: %d > %d", len(string(strVal)), t.MaxLength)
	}
	
	return nil
}

func (t *StringType) String() string {
	if t.Variable {
		return fmt.Sprintf("VARCHAR(%d)", t.MaxLength)
	}
	return fmt.Sprintf("CHAR(%d)", t.MaxLength)
}

// FloatType represents a float type
type FloatType struct {
	Precision int
	Scale     int
}

func (t *FloatType) Size() int {
	return 8 // Always 8 bytes for float64
}

func (t *FloatType) Encode(value Value) []byte {
	if value.IsNull() {
		return make([]byte, 8)
	}
	
	floatVal := float64(value.(FloatValue))
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(floatVal))
	return data
}

func (t *FloatType) Decode(data []byte) Value {
	if len(data) == 0 {
		return NullValue{}
	}
	
	floatVal := float64(binary.BigEndian.Uint64(data))
	return FloatValue(floatVal)
}

func (t *FloatType) Compare(a, b Value) int {
	if a.IsNull() && b.IsNull() {
		return 0
	}
	if a.IsNull() {
		return -1
	}
	if b.IsNull() {
		return 1
	}
	
	aVal := float64(a.(FloatValue))
	bVal := float64(b.(FloatValue))
	
	if aVal < bVal {
		return -1
	} else if aVal > bVal {
		return 1
	}
	return 0
}

func (t *FloatType) Validate(value Value) error {
	if value.IsNull() {
		return nil
	}
	
	_, ok := value.(FloatValue)
	if !ok {
		return fmt.Errorf("expected float value")
	}
	
	return nil
}

func (t *FloatType) String() string {
	return fmt.Sprintf("DECIMAL(%d,%d)", t.Precision, t.Scale)
}

// BoolType represents a boolean type
type BoolType struct{}

func (t *BoolType) Size() int {
	return 1
}

func (t *BoolType) Encode(value Value) []byte {
	if value.IsNull() {
		return []byte{0}
	}
	
	boolVal := bool(value.(BoolValue))
	if boolVal {
		return []byte{1}
	}
	return []byte{0}
}

func (t *BoolType) Decode(data []byte) Value {
	if len(data) == 0 {
		return NullValue{}
	}
	
	return BoolValue(data[0] != 0)
}

func (t *BoolType) Compare(a, b Value) int {
	if a.IsNull() && b.IsNull() {
		return 0
	}
	if a.IsNull() {
		return -1
	}
	if b.IsNull() {
		return 1
	}
	
	aVal := bool(a.(BoolValue))
	bVal := bool(b.(BoolValue))
	
	if aVal == bVal {
		return 0
	} else if !aVal && bVal {
		return -1
	}
	return 1
}

func (t *BoolType) Validate(value Value) error {
	if value.IsNull() {
		return nil
	}
	
	_, ok := value.(BoolValue)
	if !ok {
		return fmt.Errorf("expected boolean value")
	}
	
	return nil
}

func (t *BoolType) String() string {
	return "BOOLEAN"
}

// DateTimeType represents a datetime type
type DateTimeType struct {
	Precision int // seconds, milliseconds, microseconds
}

func (t *DateTimeType) Size() int {
	return 8 // Always 8 bytes for int64 timestamp
}

func (t *DateTimeType) Encode(value Value) []byte {
	if value.IsNull() {
		return make([]byte, 8)
	}
	
	dtVal := time.Time(value.(DateTimeValue))
	timestamp := dtVal.UnixNano()
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(timestamp))
	return data
}

func (t *DateTimeType) Decode(data []byte) Value {
	if len(data) == 0 {
		return NullValue{}
	}
	
	timestamp := int64(binary.BigEndian.Uint64(data))
	dt := time.Unix(0, timestamp)
	return DateTimeValue(dt)
}

func (t *DateTimeType) Compare(a, b Value) int {
	if a.IsNull() && b.IsNull() {
		return 0
	}
	if a.IsNull() {
		return -1
	}
	if b.IsNull() {
		return 1
	}
	
	aVal := time.Time(a.(DateTimeValue))
	bVal := time.Time(b.(DateTimeValue))
	
	if aVal.Before(bVal) {
		return -1
	} else if aVal.After(bVal) {
		return 1
	}
	return 0
}

func (t *DateTimeType) Validate(value Value) error {
	if value.IsNull() {
		return nil
	}
	
	_, ok := value.(DateTimeValue)
	if !ok {
		return fmt.Errorf("expected datetime value")
	}
	
	return nil
}

func (t *DateTimeType) String() string {
	return "TIMESTAMP"
}

// BytesType represents a bytes type
type BytesType struct {
	MaxLength int
}

func (t *BytesType) Size() int {
	return -1 // Variable length
}

func (t *BytesType) Encode(value Value) []byte {
	if value.IsNull() {
		return []byte{}
	}
	
	bytesVal := []byte(value.(BytesValue))
	
	// Prefix with length
	length := len(bytesVal)
	data := make([]byte, 4+length)
	binary.BigEndian.PutUint32(data, uint32(length))
	copy(data[4:], bytesVal)
	return data
}

func (t *BytesType) Decode(data []byte) Value {
	if len(data) == 0 {
		return NullValue{}
	}
	
	if len(data) < 4 {
		return NullValue{}
	}
	
	length := int(binary.BigEndian.Uint32(data))
	if len(data) < 4+length {
		return NullValue{}
	}
	
	bytes := make([]byte, length)
	copy(bytes, data[4:4+length])
	return BytesValue(bytes)
}

func (t *BytesType) Compare(a, b Value) int {
	if a.IsNull() && b.IsNull() {
		return 0
	}
	if a.IsNull() {
		return -1
	}
	if b.IsNull() {
		return 1
	}
	
	aVal := []byte(a.(BytesValue))
	bVal := []byte(b.(BytesValue))
	
	if len(aVal) < len(bVal) {
		return -1
	} else if len(aVal) > len(bVal) {
		return 1
	}
	
	for i := 0; i < len(aVal); i++ {
		if aVal[i] < bVal[i] {
			return -1
		} else if aVal[i] > bVal[i] {
			return 1
		}
	}
	
	return 0
}

func (t *BytesType) Validate(value Value) error {
	if value.IsNull() {
		return nil
	}
	
	bytesVal, ok := value.(BytesValue)
	if !ok {
		return fmt.Errorf("expected bytes value")
	}
	
	if len([]byte(bytesVal)) > t.MaxLength {
		return fmt.Errorf("bytes too long: %d > %d", len([]byte(bytesVal)), t.MaxLength)
	}
	
	return nil
}

func (t *BytesType) String() string {
	return fmt.Sprintf("BYTES(%d)", t.MaxLength)
}

// Type registry for common types
var (
	Int8Type    = &IntegerType{ByteSize: 1}
	Int16Type   = &IntegerType{ByteSize: 2}
	Int32Type   = &IntegerType{ByteSize: 4}
	Int64Type   = &IntegerType{ByteSize: 8}
	
	VarcharType = &StringType{MaxLength: 255, Variable: true}
	CharType    = &StringType{MaxLength: 255, Variable: false}
	
	FloatTypeVar   = &FloatType{Precision: 10, Scale: 2}
	DoubleTypeVar  = &FloatType{Precision: 15, Scale: 6}
	
	BoolTypeVar    = &BoolType{}
	
	TimestampTypeVar = &DateTimeType{Precision: 6}
	
	BytesTypeVar   = &BytesType{MaxLength: 65535}
)

// TypeFromString creates a type from a string representation
func TypeFromString(typeStr string) (Type, error) {
	switch typeStr {
	case "INT8", "TINYINT":
		return Int8Type, nil
	case "INT16", "SMALLINT":
		return Int16Type, nil
	case "INT32", "INT", "INTEGER":
		return Int32Type, nil
	case "INT64", "BIGINT":
		return Int64Type, nil
	case "VARCHAR":
		return VarcharType, nil
	case "CHAR":
		return CharType, nil
	case "FLOAT", "REAL":
		return FloatTypeVar, nil
	case "DOUBLE", "DOUBLE PRECISION":
		return DoubleTypeVar, nil
	case "BOOLEAN", "BOOL":
		return BoolTypeVar, nil
	case "TIMESTAMP":
		return TimestampTypeVar, nil
	case "BYTES", "BLOB":
		return BytesTypeVar, nil
	default:
		return nil, fmt.Errorf("unknown type: %s", typeStr)
	}
}
