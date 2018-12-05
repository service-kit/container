package iterator

type Iterator interface {
	Next() Iterator
	Pre() Iterator
	SetNext(Iterator)
	SetPre(Iterator)
	Value() interface{}
	SetValue(interface{})
}
