package tools

// From stackoverflow:
// https://stackoverflow.com/questions/19374219/how-to-find-the-difference-between-two-slices-of-strings
// Returns the elements in slice "a" that aren't in slice "b".
func Difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
