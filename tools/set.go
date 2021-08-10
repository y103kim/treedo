package tools

type IntSet struct {
	data map[int]bool
}

func CreateIntSet() *IntSet {
	int_set := &IntSet{nil}
	int_set.data = make(map[int]bool)
	return int_set
}

func (s *IntSet) Add(arr ...int) {
	for _, v := range arr {
		s.data[v] = true
	}
}

func (s *IntSet) Erase(arr ...int) {
	for _, v := range arr {
		delete(s.data, v)
	}
}

func (s *IntSet) Check(v int) bool {
	return s.data[v] == true
}

func (s *IntSet) List() []int {
	ret := make([]int, 0, len(s.data))
	for k := range s.data {
		ret = append(ret, k)
	}
	return ret
}

type IntSets struct {
	data map[int]*IntSet
}

func CreateIntSets() *IntSets {
	int_sets := &IntSets{nil}
	int_sets.data = make(map[int]*IntSet)
	return int_sets
}

func (s *IntSets) Add(key int, arr ...int) {
	for _, v := range arr {
		if _, ok := s.data[key]; !ok {
			s.data[key] = CreateIntSet()
		}
		s.data[key].Add(v)
	}
}

func (s *IntSets) Erase(key int, arr ...int) {
	for _, v := range arr {
		if set, ok := s.data[key]; ok {
			set.Erase(v)
		}
	}
}

func (s *IntSets) Check(key, v int) bool {
	if set, ok := s.data[key]; ok {
		return set.Check(v)
	}
	return false
}

func (s *IntSets) List(key int) []int {
	if set, ok := s.data[key]; ok {
		return set.List()
	}
	return []int{}
}
