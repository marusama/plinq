package main

import (
	"reflect"
)

type IEnumerable interface {
	Select(f interface{}) IEnumerable
	Where(f interface{}) IEnumerable
	Take(n int) IEnumerable
	ToSlice(r interface{}) error

	SelectX(f func(interface{}) interface{}) IEnumerable
	WhereX(f func(interface{}) bool) IEnumerable
}

type list struct {
	err     error
	hasNext func() bool
	next    func() interface{}
}

func FromIntSlice(s []int) IEnumerable {
	i := -1

	return &list{
		hasNext: func() bool {
			return i+1 < len(s)
		},
		next: func() interface{} {
			i++
			return s[i]
		},
	}
}

func FromSlice(s interface{}) IEnumerable {

	sl := reflect.ValueOf(s)
	len := sl.Len()
	i := -1

	return &list{
		hasNext: func() bool {
			return i+1 < len
		},
		next: func() interface{} {
			i++
			r := sl.Index(i).Interface()
			return r
		},
	}
}

func (l *list) Select(f interface{}) IEnumerable {

	refFunc := reflect.ValueOf(f)
	refFuncArgs := make([]reflect.Value, 1)
	mapper := func(x interface{}) interface{} {
		refFuncArgs[0] = reflect.ValueOf(x)
		res := refFunc.Call(refFuncArgs)
		return res[0].Interface()
	}

	return &list{
		hasNext: func() bool {
			return l.hasNext()
		},
		next: func() interface{} {
			x := l.next()
			r := mapper(x)
			return r
		},
	}
}

func (l *list) Where(f interface{}) IEnumerable {

	refFunc := reflect.ValueOf(f)
	refFuncArgs := make([]reflect.Value, 1)
	pred := func(x interface{}) bool {
		refFuncArgs[0] = reflect.ValueOf(x)
		res := refFunc.Call(refFuncArgs)
		return res[0].Bool()
	}

	var nxt interface{}

	return &list{
		hasNext: func() bool {
			for l.hasNext() {
				x := l.next()
				if pred(x) {
					nxt = x
					return true
				}
			}
			return false
		},
		next: func() interface{} {
			return nxt
		},
	}
}

func (l *list) Take(n int) IEnumerable {
	i := -1

	return &list{
		hasNext: func() bool {
			if l.hasNext() {
				i++
				return i < n
			}
			return false
		},
		next: func() interface{} {
			return l.next()
		},
	}
}

func (l *list) ToSlice(r interface{}) error {

	sl := reflect.ValueOf(r).Elem()
	for l.hasNext() {
		x := l.next()
		newItem := reflect.ValueOf(x)
		sl.Set(reflect.Append(sl, newItem))
	}
	return nil
}

func (l *list) SelectX(f func(interface{}) interface{}) IEnumerable {

	return &list{
		hasNext: func() bool {
			return l.hasNext()
		},
		next: func() interface{} {
			x := l.next()
			r := f(x)
			return r
		},
	}
}

func (l *list) WhereX(f func(interface{}) bool) IEnumerable {

	var nxt interface{}

	return &list{
		hasNext: func() bool {
			for l.hasNext() {
				x := l.next()
				if f(x) {
					nxt = x
					return true
				}
			}
			return false
		},
		next: func() interface{} {
			return nxt
		},
	}
}
