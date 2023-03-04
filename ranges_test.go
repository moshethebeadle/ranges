package ranges

import (
	"reflect"
	"testing"
)

// intRange is an INCLUSIVE range that guarantees that start <= end
type intRange struct {
	start int
	end   int
}

func (ir intRange) Start() int {
	return ir.start
}

func (ir intRange) End() int {
	return ir.end
}

func Test_isBetween(t *testing.T) {
	r := intRange{
		start: 0,
		end:   10,
	}
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "below",
			x:    -1,
			want: false,
		},
		{
			name: "left edge",
			x:    0,
			want: true,
		},
		{
			name: "middle",
			x:    5,
			want: true,
		},
		{
			name: "right edge",
			x:    10,
			want: true,
		},
		{
			name: "above",
			x:    11,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBetween(tt.x, r); got != tt.want {
				t.Errorf("isBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindOverlaps(t *testing.T) {
	left := intRange{2, 4}
	right := intRange{6, 8}
	ranges := []intRange{left, right}

	tests := []struct {
		name      string
		candidate intRange
		want      []intRange
	}{
		{
			name:      "no overlaps - before",
			candidate: intRange{0, 1},
			want:      nil,
		},
		{
			name:      "no overlaps - middle",
			candidate: intRange{5, 5},
			want:      nil,
		},
		{
			name:      "overlap - left side",
			candidate: intRange{1, 3},
			want:      []intRange{left},
		},
		{
			name:      "overlap - right side",
			candidate: intRange{3, 5},
			want:      []intRange{left},
		},
		{
			name:      "overlap - around one",
			candidate: intRange{1, 5},
			want:      []intRange{left},
		},
		{
			name:      "overlap - within both",
			candidate: intRange{3, 7},
			want:      []intRange{left, right},
		},
		{
			name:      "overlap - in left, past right",
			candidate: intRange{3, 9},
			want:      []intRange{left, right},
		},
		{
			name:      "overlap - before left, past right",
			candidate: intRange{1, 9},
			want:      []intRange{left, right},
		},
		{
			name:      "overlap - validate edges",
			candidate: intRange{4, 6},
			want:      []intRange{left, right},
		},
		{
			name:      "overlap - validate edges 2",
			candidate: intRange{2, 8},
			want:      []intRange{left, right},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOverlaps(tt.candidate, ranges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
