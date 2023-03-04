package ranges

// Range has a start and an end
// guaranteed that start <= end
type Range interface {
	Start() int
	End() int
}

// FindOverlaps find all ranges that candidate overlaps with
// assumes all ranges are valid, ie Start() <= End()
func FindOverlaps[K Range](candidate K, ranges []K) []K {
	var result []K
	for _, r := range ranges {
		if isBetween(candidate.Start(), r) ||
			isBetween(candidate.End(), r) ||
			isBetween(r.Start(), candidate) ||
			isBetween(r.End(), candidate) {
			result = append(result, r)
		}
	}

	return result
}

func isBetween(x int, r Range) bool {
	return x >= r.Start() && x <= r.End()
}
