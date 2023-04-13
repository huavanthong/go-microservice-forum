package common

import (
	"reflect"
	"sort"
)

type ValueObject struct{}

func EqualOperator(left, right *ValueObject) bool {
	if (left == nil) != (right == nil) {
		return false
	}
	if left == nil {
		return true
	}
	return left.Equals(right)
}

func NotEqualOperator(left, right *ValueObject) bool {
	return !EqualOperator(left, right)
}

func (vo *ValueObject) GetEqualityComponents() []interface{} {
	// Override this method in subclasses to return the components that define equality
	return []interface{}{}
}

func (vo *ValueObject) Equals(obj interface{}) bool {

	if obj == nil || reflect.TypeOf(obj) != reflect.TypeOf(vo) {
		return false
	}

	other := obj.(*ValueObject)
	thisComponents := vo.GetEqualityComponents()
	otherComponents := other.GetEqualityComponents()
	if len(thisComponents) != len(otherComponents) {
		return false
	}
	for i := range thisComponents {
		if thisComponents[i] != otherComponents[i] {
			return false
		}
	}
	return true
}

const hash31 int = 31

func (vo ValueObject) HashCode() int {
	components := vo.GetEqualityComponents()
	hashes := make([]int, len(components))
	for i, component := range components {
		if component != nil {
			hashes[i] = component.(interface{ HashCode() int }).HashCode()
		}
	}
	sort.Ints(hashes)
	hash := 17
	for _, h := range hashes {
		hash = hash31 + h
	}
	return hash
}
