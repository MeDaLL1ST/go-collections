package queue

type QueueInt []int
type QueueStr []string
type QueueF []float64

type Queue interface {
	AddStr(element string) QueueStr
	AddInt(element int) QueueInt
	AddF(element float64) QueueF
	RemoveInt() QueueInt
	RemoveStr() QueueStr
	RemoveF() QueueF
	IsEmptyInt() bool
	IsEmptyStr() bool
	IsEmptyF() bool
	SizeInt() int
	SizeStr() int
	SizeF() int
	PeekInt() int
	PeekStr() string
	PeekF() float64
	PollInt() (QueueInt, int)
	PollStr() (QueueStr, string)
	PollF() (QueueF, float64)
}

func (q2 QueueStr) AddStr(element string) QueueStr {
	return append(q2, element)
}

func (q2 QueueInt) AddInt(element int) QueueInt {
	return append(q2, element)
}

func (q2 QueueF) AddF(element float64) QueueF {
	return append(q2, element)
}

func (q2 QueueInt) RemoveInt() QueueInt {
	len1 := len(q2)
	if len1 > 0 {
		return q2[:len(q2)-1]
	} else {
		return q2
	}
}

func (q2 QueueStr) RemoveStr() QueueStr {
	len1 := len(q2)
	if len1 > 0 {
		return q2[:len(q2)-1]
	} else {
		return q2
	}
}

func (q2 QueueF) RemoveF() QueueF {
	len1 := len(q2)
	if len1 > 0 {
		return q2[:len(q2)-1]
	} else {
		return q2
	}
}

func (q2 QueueInt) IsEmptyInt() bool {
	if len(q2) == 0 {
		return true
	} else {
		return false
	}
}
func (q2 QueueStr) IsEmptyStr() bool {
	if len(q2) == 0 {
		return true
	} else {
		return false
	}
}
func (q2 QueueF) IsEmptyF() bool {
	if len(q2) == 0 {
		return true
	} else {
		return false
	}
}

func (q2 QueueInt) SizeInt() int {
	i := len(q2)
	return i
}
func (q2 QueueStr) SizeStr() int {
	i := len(q2)
	return i
}
func (q2 QueueF) SizeF() int {
	i := len(q2)
	return i
}
func (q2 QueueInt) PeekInt() int {
	i := len(q2)
	if i > 0 {
		i2 := q2[i-1]
		return i2
	} else {
		panic("len of queue is 0")
	}
}
func (q2 QueueStr) PeekStr() string {
	i := len(q2)
	if i > 0 {
		i2 := q2[i-1]
		return i2
	} else {
		panic("len of queue is 0")
	}
}
func (q2 QueueF) PeekF() float64 {
	i := len(q2)
	if i > 0 {
		i2 := q2[i-1]
		return i2
	} else {
		panic("len of queue is 0")
	}
}
func (q2 QueueInt) PollInt() (QueueInt, int) {
	i := len(q2)
	if i > 0 {
		i2 := q2[i-1]
		return q2[:len(q2)-1], i2
	} else {
		panic("len of queue is 0")
	}
}
func (q2 QueueStr) PollStr() (QueueStr, string) {
	i := len(q2)
	if i > 0 {
		i2 := q2[i-1]
		return q2[:len(q2)-1], i2
	} else {
		panic("len of queue is 0")
	}
}
func (q2 QueueF) PollF() (QueueF, float64) {
	i := len(q2)
	if i > 0 {
		i2 := q2[i-1]
		return q2[:len(q2)-1], i2
	} else {
		panic("len of queue is 0")
	}
}