package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type List struct {
	data []interface{}
	maxData int
}

func New(size int) Queue {
	impl := &List{
		maxData: size,
	}
	return impl
}

func (impl *List) Push(key interface{}) {
	if impl.Len() == impl.maxData {
		impl.Pop()
	}
	impl.data = append(impl.data, key)
}

func (impl *List) Pop() interface{} {
	var value interface{}
	newdata := make([]interface{}, 0)
	if impl.Len() > 0 {
		value = impl.data[0]
		if impl.Len() > 1 {
			for i := 1; i < impl.Len(); i++ {
				newdata = append(newdata, impl.data[i])
			}
		}
	}
	impl.data = newdata
	return value
}

func (impl *List) Contains(key interface{}) bool {
	isContain := false
	for _,data := range impl.data{
		if data == key {
			isContain = true
		}
	}
	return isContain
}

func (impl *List) Len() int {
	return len(impl.data)
}

func (impl *List) Keys() []interface{} {
	return impl.data
}