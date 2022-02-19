package main

import "fmt"

type Vector interface {
	Len() int
	Cap() int
	LastIndex() int
	IsFull() bool
	NewVector(a ...interface{}) *vector
	Append(a interface{})
	Appends(a ...interface{})
	PushBack(a interface{})
	InsertAt(index int, a interface{})
	Pop() interface{}
	PopFront() interface{}
	PushFront(a interface{})
	Contains(a interface{}) bool
	Find(a interface{}) int
	FindAll(a interface{}) []int
	Delete(a interface{})
	DeleteAll(a interface{})
	Uniques() *vector
	Reverse()
	ReverseReturn() *vector
	View()
	ViewList()
}

// Base, generic vector type
type vector []interface{}

// Return length (size) of vector.
func (v vector) Len() int { return len(v) }

// Returns capacity of vector.
func (v vector) Cap() int { return cap(v) }

// Returns last populated index of vector.
func (v vector) LastIndex() int {
	if v.Len() > 0 {
		return v.Len() - 1
	}
	return 0
}

// Quick check of space within vector instance.
// Returns true if capacity equals length, false otherwise.
func (v vector) IsFull() bool { return (v.Cap() - v.Len()) <= 0 }

// Create and return new vector instance with items
func NewVector(a ...interface{}) *vector {
	var vec = new(vector)
	for _, el := range a {
		vec.PushBack(el)
	}
	return vec
}

// Append an item to a vector instance.
func (v *vector) Append(a interface{}) {
	*v = append(*v, a)
}

// Append more than one item.
func (v *vector) Appends(a ...interface{}) {
	for _, el := range a {
		*v = append(*v, el)
	}
}

// Alias for Append()
func (v *vector) PushBack(a interface{}) {
	*v = append(*v, a)
}

// Insert item at given index.
func (v *vector) InsertAt(index int, a interface{}) {
	var vv *vector = &vector{a}
	*v = append((*v)[:index], append(*vv, (*v)[index:]...)...)
}

// Pop last item from vector.
func (v *vector) Pop() interface{} {
	var vout interface{}
	vout, *v = (*v)[v.LastIndex()], (*v)[:v.LastIndex()]
	return vout
}

// Pop item from front of vector.
func (v *vector) PopFront() interface{} {
	var out interface{}
	out, *v = (*v)[0], (*v)[1:]
	return out
}

// Push item to front of vector
func (v *vector) PushFront(a interface{}) {
	var vv *vector = &vector{a}
	*v = append(*vv, *v...)
}

// Determine if item in vector
func (v *vector) Contains(a interface{}) bool {
	return v.Find(a) > -1
}

// Looks for item within vector instance.
// Returns index of first successful result, -1 otherwise.
func (v *vector) Find(a interface{}) int {
	for i, el := range *v {
		if el == a {
			return i
		}
	}
	return -1
}

// Looks for item within vector instance.
// Returns slice of all matching indices within vector.
func (v *vector) FindAll(a interface{}) []int {
	var outs []int = make([]int, 0, v.Len())
	for i, el := range *v {
		if el == a {
			outs = append(outs, i)
		}
	}
	return outs
}

// Delete first matching item from vector
func (v *vector) Delete(a interface{}) {
	idx := v.Find(a)
	*v = append((*v)[:idx], (*v)[idx+1:])
}

// Delete all arguments from vector
func (v *vector) DeleteAll(a interface{}) {
	idxs := v.FindAll(a)
	for _, idx := range idxs {
		*v = append((*v)[:idx], (*v)[idx+1:])
	}
}

// Unique values of vector instance.
// Returned object is unsorted.
func (v *vector) Uniques() *vector {
	var visitMap = make(map[interface{}]bool)
	var vout *vector = new(vector)
	for _, el := range *v {
		if _, ok := visitMap[el]; !ok {
			visitMap[el] = true
			*vout = append(*vout, el)
		}
	}
	return vout
}

// func (v *vector) InsertAt(index int, a interface{}) {
// 	*v = append((*v)[:index], append([]a, (*v)[index:]...)...)
// }

// Reverse the order of a vector.
// This mentod works inplace, so the original vector will be modified.
func (v *vector) Reverse() {
	for left, right := 0, v.Len()-1; left < right; left, right = left+1, right-1 {
		(*v)[left], (*v)[right] = (*v)[right], (*v)[left]
	}
}

func (v vector) ReverseReturn() *vector {
	var vout = new(vector)
	*vout = append(*vout, v...)
	for left, right := 0, vout.Len()-1; left < right; left, right = left+1, right-1 {
		(*vout)[left], (*vout)[right] = (*vout)[right], (*vout)[left]
	}
	return vout
}

// --------------------//
// --- Debug/Check --- //
// --------------------//

// View items as space=separated output.
func (v vector) View() {
	var t string = ""
	for _, el := range v {
		t += fmt.Sprintf("%v ", el)
	}
	fmt.Printf("%s\n", t)
}

// View items as list of items.
func (v vector) ViewList() {
	for _, el := range v {
		fmt.Printf("%v\n", el)
	}
}

// ---------------------//
// --- Feature Demo --- //
// ---------------------//
func Demo() {
	var vec = NewVector(1, 2, 3, 3, 6, 6, 1, 4, 3, 9)
	fmt.Printf("Type: %T\n", vec)

	// --- View unique items.
	res := vec.Uniques()
	res.View()
	res.ViewList()
	// fmt.Println("")

	// -- Reversing the vector
	rvec := vec.ReverseReturn()
	fmt.Printf("Original: %v\n", vec)
	fmt.Printf("Xtra crispy: %v\n", rvec)

	// -- Inserting a value
	vec.InsertAt(1, 5)
	fmt.Printf("Post-insert: %v\n", vec)
}
