package subset

import (
	"testing"
)

type TestStruct1 struct {
	Field1       int
	Field2       string
	Field3       float32
	Field4       bool
	Field5       interface{}
	Field6       func()
	Field7       map[string]int
	Field8       *int
	NestedField1 NestedStruct1
}

type NestedStruct1 struct {
	Field1       int
	Field2       string
	Field3       float32
	NestedField2 NestedStruct2
}

type NestedStruct2 struct {
	field5 interface{}
	field6 func()
}

type TestStruct2 struct {
	Field1       int
	Field2       string
	Field3       float32
	Field4       bool
	Field5       interface{}
	Field6       func()
	Field7       map[string]int
	Field8       *int
	NestedField1 NestedStruct1
}

type TestStruct3 struct {
	Field7       map[string]int
	Field8       *int
	NestedField1 NestedStruct1
}

type TestStruct4 struct {
	Field1       int
	Field2       string
	Field3       float32
	Field4       bool
	Field5       interface{}
	Field6       func()
	Field7       map[string]int
	Field8       *int
	Field9       *int
	NestedField1 NestedStruct1
}

func TestSubset(t *testing.T) {
	tt := []struct {
		Name    string
		StructA interface{}
		StructB interface{}
		Output  bool
	}{
		{
			Name:    "Structs are equal",
			StructA: TestStruct1{},
			StructB: TestStruct2{},
			Output:  true,
		},
		{
			Name:    "Structs are Different",
			StructA: TestStruct1{},
			StructB: NestedStruct1{},
			Output:  false,
		},
		{
			Name:    "StructA is a subset of StructB",
			StructA: TestStruct3{},
			StructB: TestStruct1{},
			Output:  true,
		},
		{
			Name:    "StructA has more fields than StructB",
			StructA: TestStruct4{},
			StructB: TestStruct1{},
			Output:  false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			res := Subset(tc.StructA, tc.StructB)
			if res != tc.Output {
				t.Fatalf("Expected %v, but got %v", tc.Output, res)
			}
		})
	}
}
