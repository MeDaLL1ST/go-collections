package list

type LinkedListInt []int
type LinkedListStr []string
type LinkedListF []float64

type List interface {
	AddStr(element string, index int) LinkedListStr
	AddInt(element int, index int) LinkedListInt
	AddF(element float64, index int) LinkedListF
	AddLastStr(element string) LinkedListStr
	AddLastInt(element int) LinkedListInt
	AddLastF(element float64) LinkedListF
	AddFirstStr(element string) LinkedListStr
	AddFirstInt(element int) LinkedListInt
	AddFirstF(element float64) LinkedListF
	RemoveInt(index int) LinkedListInt
	RemoveStr(index int) LinkedListStr
	RemoveF(index int) LinkedListF
	RemoveLastInt() LinkedListInt
	RemoveLastStr() LinkedListStr
	RemoveLastF() LinkedListF
	IsEmptyInt() bool
	IsEmptyStr() bool
	IsEmptyF() bool
	SizeInt() int
	SizeStr() int
	SizeF() int
	GetInt(index int) int
	GetStr(index int) string
	GetF(index int) float64
	GetListInt(start int, end int) LinkedListInt
	GetListStr(start int, end int) LinkedListStr
	GetListF(start int, end int) LinkedListF
}

func (l LinkedListInt) AddInt(element int, index int) LinkedListInt {
	as1 := make([]int, index-1)
	copy(as1, l[:index-1])
	as2 := make([]int, len(l)-index+1)
	copy(as2, l[index-1:])
	as1 = append(as1, element)
	as1 = append(as1, as2...)
	return as1
}

func (l LinkedListInt) AddLastInt(element int) LinkedListInt {
	return append(l, element)
}

func (l LinkedListInt) AddFirstInt(element int) LinkedListInt {
	index := 1
	as1 := make([]int, index-1)
	copy(as1, l[:index-1])
	as2 := make([]int, len(l)-index+1)
	copy(as2, l[index-1:])
	as1 = append(as1, element)
	as1 = append(as1, as2...)
	return as1
}

func (l LinkedListInt) RemoveInt(index int) LinkedListInt {
	as1 := make([]int, index-1)
	copy(as1, l[:index-1])
	as2 := make([]int, len(l)-index)
	copy(as2, l[index:])
	as1 = append(as1, as2...)
	return as1
}

func (l LinkedListInt) RemoveLastInt() LinkedListInt {
	len1 := len(l)
	if len1 > 0 {
		return l[:len(l)-1]
	} else {
		return l
	}
}

func (l LinkedListInt) IsEmptyInt() bool {
	if len(l) == 0 {
		return true
	} else {
		return false
	}
}

func (l LinkedListInt) SizeInt() int {
	i := len(l)
	return i
}

func (l LinkedListInt) GetInt(index int) int {
	i := l[index]
	return i
}

func (l LinkedListInt) GetListInt(start int, end int) LinkedListInt {
	return l[start:end]
}

func (l LinkedListStr) AddStr(element string, index int) LinkedListStr {
	as1 := make([]string, index-1)
	copy(as1, l[:index-1])
	as2 := make([]string, len(l)-index+1)
	copy(as2, l[index-1:])
	as1 = append(as1, element)
	as1 = append(as1, as2...)
	return as1
}

func (l LinkedListStr) AddLastStr(element string) LinkedListStr {
	return append(l, element)
}

func (l LinkedListStr) AddFirstStr(element string) LinkedListStr {
	index := 1
	as1 := make([]string, index-1)
	copy(as1, l[:index-1])
	as2 := make([]string, len(l)-index+1)
	copy(as2, l[index-1:])
	as1 = append(as1, element)
	as1 = append(as1, as2...)
	return as1
}

func (l LinkedListStr) RemoveStr(index int) LinkedListStr {
	as1 := make([]string, index-1)
	copy(as1, l[:index-1])
	as2 := make([]string, len(l)-index)
	copy(as2, l[index:])
	as1 = append(as1, as2...)
	return as1
}

func (l LinkedListStr) RemoveLastStr() LinkedListStr {
	len1 := len(l)
	if len1 > 0 {
		return l[:len(l)-1]
	} else {
		return l
	}
}

func (l LinkedListStr) IsEmptyStr() bool {
	if len(l) == 0 {
		return true
	} else {
		return false
	}
}

func (l LinkedListStr) SizeStr() int {
	i := len(l)
	return i
}

func (l LinkedListStr) GetStr(index int) string {
	i := l[index]
	return i
}

func (l LinkedListStr) GetListStr(start int, end int) LinkedListStr {
	return l[start:end]
}

func (l LinkedListF) AddF(element float64, index int) LinkedListF {
	as1 := make([]float64, index-1)
	copy(as1, l[:index-1])
	as2 := make([]float64, len(l)-index+1)
	copy(as2, l[index-1:])
	as1 = append(as1, element)
	as1 = append(as1, as2...)
	return as1
}

func (l LinkedListF) AddLastF(element float64) LinkedListF {
	return append(l, element)
}

func (l LinkedListF) AddFirstF(element float64) LinkedListF {
	index := 1
	as1 := make([]float64, index-1)
	copy(as1, l[:index-1])
	as2 := make([]float64, len(l)-index+1)
	copy(as2, l[index-1:])
	as1 = append(as1, element)
	as1 = append(as1, as2...)
	return as1
}

func (l LinkedListF) RemoveF(index int) LinkedListF {
	as1 := make([]float64, index-1)
	copy(as1, l[:index-1])
	as2 := make([]float64, len(l)-index)
	copy(as2, l[index:])
	as1 = append(as1, as2...)
	return as1
}

func (l LinkedListF) RemoveLastF() LinkedListF {
	len1 := len(l)
	if len1 > 0 {
		return l[:len(l)-1]
	} else {
		return l
	}
}

func (l LinkedListF) IsEmptyF() bool {
	if len(l) == 0 {
		return true
	} else {
		return false
	}
}

func (l LinkedListF) SizeF() int {
	i := len(l)
	return i
}

func (l LinkedListF) GetF(index int) float64 {
	i := l[index]
	return i
}

func (l LinkedListF) GetListF(start int, end int) LinkedListF {
	return l[start:end]
}