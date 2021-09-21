package mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapper(t *testing.T) {

	type subSource struct {
		D string
	}

	type subTarget struct {
		D string
	}

	type source struct {
		A string
		B int
		C bool
		S subSource
	}

	type target struct {
		A string
		B string
		S subTarget
	}

	tests := map[string]struct {
		target     target
		source     source
		wantTarget target
		wantErr    bool
	}{
		"when target has the source fields": {
			target:     target{},
			source:     source{A: "A"},
			wantTarget: target{A: "A"},
			wantErr:    false,
		},
		"when source has sub-fields": {
			target:     target{},
			source:     source{A: "A", S: subSource{D: "D"}},
			wantTarget: target{A: "A", S: subTarget{D: "D"}},
			wantErr:    false,
		},
		"when has fields with same name and different types": {
			target:     target{},
			source:     source{A: "A", B: 0},
			wantTarget: target{A: "A", B: ""},
			wantErr:    false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assertion := assert.New(t)

			err := Mapper(tt.source, &tt.target)
			assertion.Equal((err != nil), tt.wantErr, "Mapper() error: expected %v, got %v", tt.wantErr, err)
			assertion.Equal(tt.target, tt.wantTarget, "Mapper(): expected %#v, got %#v", tt.wantTarget, tt.target)
		})
	}
}
