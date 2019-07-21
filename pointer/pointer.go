package pointer

func Float64(n int) *float64 {
	x := float64(n)
	return &x
}

func Int32(n int) *int32 {
	x := int32(n)
	return &x
}

func True() *bool {
	x := true
	return &x
}

func False() *bool {
	x := false
	return &x
}
