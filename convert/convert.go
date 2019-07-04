package convert

import (
	"encoding/json"

	"github.com/gogo/protobuf/types"
)

func Int32Value(x *types.Int32Value) *int32 {
	if x == nil {
		return nil
	}
	value := x.Value
	return &value
}

func Int32ValueDefault(x *types.Int32Value, defaultValue int32) int32 {
	if x == nil {
		return defaultValue
	}
	return x.Value
}

func Float64Value(x *types.DoubleValue) *float64 {
	if x == nil {
		return nil
	}
	value := x.Value
	return &value
}

func BoolValue(x *types.BoolValue) *bool {
	if x == nil {
		return nil
	}
	value := x.Value
	return &value
}

func BoolValueDefault(x *types.BoolValue, defaultValue bool) bool {
	if x == nil {
		return defaultValue
	}
	return x.Value
}

func NumberToFloat64(num json.Number) float64 {
	value, err := num.Float64()
	if err != nil {
		return 0
	}
	return value
}

func StringValue(x *types.StringValue) *string {
	if x == nil {
		return nil
	}
	value := x.Value
	return &value
}
