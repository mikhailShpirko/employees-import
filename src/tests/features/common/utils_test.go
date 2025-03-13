package tests

import (
	common "employees-import/features/common"
	"testing"
	"time"
)

func Test_IsStringEmptyOrWhiteSpace_NonEmptyString_True(t *testing.T) {
	result := common.IsStringEmptyOrWhiteSpace("ABCD")

	if result {
		t.Fatalf(`ABCD is valid non empty string`)
	}
}

func Test_IsStringEmptyOrWhiteSpace_NonEmptyStringWithSpaces_True(t *testing.T) {
	result := common.IsStringEmptyOrWhiteSpace(" ABCD ")

	if result {
		t.Fatalf(`" ABCD " is valid non empty string`)
	}
}

func Test_IsStringEmptyOrWhiteSpace_EmptyString_False(t *testing.T) {
	result := common.IsStringEmptyOrWhiteSpace("")

	if !result {
		t.Fatalf(`Empty string should not be processed as not empty`)
	}
}

func Test_IsStringEmptyOrWhiteSpace_Space_False(t *testing.T) {
	result := common.IsStringEmptyOrWhiteSpace(" ")

	if !result {
		t.Fatalf(`Space string should not be processed as not empty`)
	}
}

func Test_IsStringEmptyOrWhiteSpace_MultipleSpaces_False(t *testing.T) {
	result := common.IsStringEmptyOrWhiteSpace("   ")

	if !result {
		t.Fatalf(`Multiple Spaces string should not be processed as not empty`)
	}
}

func Test_IsTimeBefore_SrcIsBefore_True(t *testing.T) {
	src, _ := time.Parse(time.DateOnly, "2000-01-01")
	future, _ := time.Parse(time.DateOnly, "2024-12-01")

	result := common.IsTimeBefore(src, future)

	if !result {
		t.Fatalf(`%v should not be greater than %v`, src, future)
	}
}

func Test_IsTimeBefore_SrcIsAfter_False(t *testing.T) {
	src, _ := time.Parse(time.DateOnly, "2034-05-20")
	future, _ := time.Parse(time.DateOnly, "2024-12-01")

	result := common.IsTimeBefore(src, future)

	if result {
		t.Fatalf(`%v should be greater than %v`, src, future)
	}
}

func Test_IsStringNumeric_NumericString_True(t *testing.T) {
	result := common.IsStringNumeric("998991234567")

	if !result {
		t.Fatalf(`998991234567 is a numeric string`)
	}
}

func Test_IsStringNumeric_CharacterString_False(t *testing.T) {
	result := common.IsStringNumeric("asdfghjkll")

	if result {
		t.Fatalf(`asdfghjkll is not a numeric string`)
	}
}

func Test_IsStringNumeric_AlphaNumericString_False(t *testing.T) {
	result := common.IsStringNumeric("123asdfg456hjkll789")

	if result {
		t.Fatalf(`123asdfg456hjkll789 is not a numeric string`)
	}
}
