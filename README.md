# go-collections

LinkedList and Queue realization for golang(Java analog) 

<code>
go get github.com/MeDaLL1ST/go-collections
</code>
  ignore errors
  
  ## Import
  
 ```golang
 import (
	"github.com/MeDaLL1ST/go-collections/list"
	"github.com/MeDaLL1ST/go-collections/queue"
)
```

## Using

There are only 3 data types(int, string, float):
```golang
lls:=list.LinkedListStr{}
lli:=list.LinkedListInt{}
llf:=list.LinkedListF{}

qi:=queue.QueueInt{}
qs:=queue.QueueStr{}
qf:=queue.QueueF{}
```
They have the same methods and differ only in data types

```golang
type List interface {
	AddStr(element string, index int) LinkedListStr //Adds an element by a specific index
	AddInt(element int, index int) LinkedListInt
	AddF(element float64, index int) LinkedListF
	AddLastStr(element string) LinkedListStr //Adds an element by a len() index
	AddLastInt(element int) LinkedListInt
	AddLastF(element float64) LinkedListF
	AddFirstStr(element string) LinkedListStr //Adds an element by a 1 index
	AddFirstInt(element int) LinkedListInt
	AddFirstF(element float64) LinkedListF
	RemoveInt(index int) LinkedListInt //Deletes an element by index
	RemoveStr(index int) LinkedListStr
	RemoveF(index int) LinkedListF
	RemoveLastInt() LinkedListInt //Deletes an element by len() index
	RemoveLastStr() LinkedListStr
	RemoveLastF() LinkedListF
	IsEmptyInt() bool //return true if list is empty
	IsEmptyStr() bool
	IsEmptyF() bool
	SizeInt() int //return size of list
	SizeStr() int
	SizeF() int
	GetInt(index int) int //return element by specific index
	GetStr(index int) string
	GetF(index int) float64
	GetListInt(start int, end int) LinkedListInt //return list of elements by specific index
	GetListStr(start int, end int) LinkedListStr
	GetListF(start int, end int) LinkedListF
}

type Queue interface {
	AddStr(element string) QueueStr //Adds an element by a len() index
	AddInt(element int) QueueInt
	AddF(element float64) QueueF
	RemoveInt() QueueInt //Remove an element by a len() index
	RemoveStr() QueueStr
	RemoveF() QueueF
	IsEmptyInt() bool //return true if queue is empty
	IsEmptyStr() bool
	IsEmptyF() bool
	SizeInt() int //return size of queue
	SizeStr() int
	SizeF() int
	PeekInt() int //return last element in a queue
	PeekStr() string
	PeekF() float64
	PollInt() (QueueInt, int) //return and deleting last element in a queue
	PollStr() (QueueStr, string)
	PollF() (QueueF, float64)
}
```

## Example

```golang
lls:=list.LinkedListInt{}
lls=lls.AddFirstInt(123)
lls=lls.AddLastInt(456)
size:=lls.SizeInt()
i:=lls.GetInt(1)
fmt.Println(i,size) // returning 456 2
```
-----------------------------------------------------------------
## Open to forks of this repository
