package entity

import (
	"fmt"
	"errors"
	"testing"
)

func TestDistinguishGrade(t *testing.T) {
	type test struct {
		id string
		name string
		expect grade
		expectError error
	}

	testValues := []test{
		{
			"21T325",
			"hoge",
			B1,
			nil,
		},
		{
			"20T000",
			"fuho",
			B2,
			nil,
		},
		{
			"19T325",
			"higuruchi",
			B3,
			nil,
		},
		{
			"19A000",
			"fuga",
			B3,
			nil,
		},
		{
			"00T000",
			"fugahoge",
			Error,
			ErrInvalidYear,
		},
		{
			"AAT000",
			"fuho",
			Error,
			ErrInvalidYear,
		},
		{
			"AAAT000",
			"fuga",
			Error,
			ErrInvalidYear,
		},
		{
			"AAAAT000",
			"fuga",
			Error,
			ErrInvalidYear,
		},
	}
	
	for _, testValue := range testValues {
		entity, _:= NewParticipantEntity(testValue.id, testValue.name)
		result, err := entity.DistinguishGrade()

		fmt.Printf("%v", ErrInvalidYear)

		if result != testValue.expect {
			t.Errorf("expect: %v result: %v", testValue.expect, result)
		}

		if result == Error {
			if !errors.Is(ErrInvalidYear, err) {
				t.Errorf("expect: %v result: %v", testValue.expectError, err)
			}

		}
	}

	t.Log("finish test of entity.DistinguishGrade module")
}