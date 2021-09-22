package mapper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapper(t *testing.T) {

	type subSource struct{ D string }
	type subTarget struct{ D string }

	type source struct {
		A  string
		B  int
		C  bool
		E  []float32
		S  subSource
		P  *subSource
		PP string
	}

	type target struct {
		A  string
		B  string
		E  []float32
		S  subTarget
		P  *subTarget
		PP *subTarget
	}

	tests := map[string]struct {
		target     *target
		source     source
		wantTarget *target
		wantErr    bool
	}{
		"when target has the source fields": {
			target:     &target{},
			source:     source{A: "A", E: []float32{1.0, 2.0}},
			wantTarget: &target{A: "A", E: []float32{1.0, 2.0}},
			wantErr:    false,
		},
		"when source has sub-fields": {
			target:     &target{},
			source:     source{A: "A", S: subSource{D: "D"}},
			wantTarget: &target{A: "A", S: subTarget{D: "D"}},
			wantErr:    false,
		},
		"when has fields with same name and different types": {
			target:     &target{},
			source:     source{A: "A", B: 0},
			wantTarget: &target{A: "A", B: ""},
			wantErr:    false,
		},
		"when target is nil": {
			target:     nil,
			source:     source{A: "A"},
			wantTarget: nil,
			wantErr:    true,
		},
		"when has sub-field pointer": {
			target:     &target{P: &subTarget{D: ""}},
			source:     source{P: &subSource{D: "D"}},
			wantTarget: &target{P: &subTarget{D: "D"}},
			wantErr:    false,
		},
		"when has pointer with same name and different types": {
			target:     &target{PP: &subTarget{D: ""}},
			source:     source{PP: ""},
			wantTarget: &target{PP: &subTarget{D: ""}},
			wantErr:    false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assertion := assert.New(t)

			err := Mapper(tt.source, tt.target)
			assertion.Equal((err != nil), tt.wantErr, "Mapper() error: expected %v, got %v", tt.wantErr, err)
			assertion.Equal(tt.target, tt.wantTarget, "Mapper(): expected %#v, got %#v", tt.wantTarget, tt.target)
		})
	}
}

func ExampleMapper() {
	target := struct {
		A string
		B int
		C bool
	}{}
	source := struct {
		A string
		B int
		C bool
	}{A: "A", B: 1, C: true}
	err := Mapper(source, &target)
	fmt.Printf("Target: %v\nError: %v", target, err)
	//Output:
	//Target: {A 1 true}
	//Error: <nil>
}
