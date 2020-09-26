package specifics

type strset = map[string]struct{}

func setAdd(set *strset, elem string) {
	(*set)[elem] = struct{}{}
}

func makeStringSet(strarr []string) strset {
	set := make(strset)
	for _, s := range strarr {
		setAdd(&set, s)
	}
	return set
}

func setContains(set *strset, elem string) bool {
	// question: should a nil set be considered
	// empty?
	// if set == nil return false
	// or would this lead some hard-to-track-down bugs?
	_, ok := (*set)[elem]
	return ok
}

func setUnion(set1, set2 *strset) *strset {
	newSet := make(strset)
	for k := range *set1 {
		setAdd(&newSet, k)
	}
	for k := range *set2 {
		setAdd(&newSet, k)
	}
	return &newSet
}

// StringSet is a set with strings as elements
type StringSet struct {
	strset
}

// Add adds an element to the set
func (set *StringSet) Add(s string) {
	setAdd(&set.strset, s)
}

// MakeStringSet makes a new string set
func MakeStringSet(slice []string) StringSet {
	embedded := makeStringSet(slice)
	return StringSet{embedded}
}

// NewStringSet creates a new, empty string set
func NewStringSet() StringSet {
	embedded := make(strset)
	return StringSet{embedded}
}

// Contains checks whether the given string exists
// in the set
func (set *StringSet) Contains(s string) bool {
	return setContains(&set.strset, s)
}

// StringSetUnion returns the union of two string
// sets without modifying the originals
func StringSetUnion(set1, set2 StringSet) StringSet {
	newEmbedded := setUnion(&set1.strset, &set2.strset)
	return StringSet{*newEmbedded}
}
