package main

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	lil := NewLinkedList()

	if ln := lil.Length(); ln != 0 {
		t.Errorf("Length was incorrect, got: %d, want: 0", ln)
	}

	lil.AddAtFront(1)
	lil.AddAtFront(2)
	lil.AddAtFront(4)
	lil.AddAtEnd(5)

	arr, err := lil.ToArray()
	if err != nil {
		t.Errorf("Position in array were incorrect, err: %s", err)
	}

	arrExpected := []interface{}{4, 2, 1, 5}
	if reflect.DeepEqual(arr, arrExpected) == false {
		t.Errorf("Position in array were incorrect, got: %v, want: %v", arr, arrExpected)
	}

	if ln := lil.Length(); ln != 4 {
		t.Errorf("Length was incorrect, got: %d, want: 4", ln)
	}

	lil.Delete(2)
	arr, err = lil.ToArray()
	if err != nil {
		t.Errorf("Position in array were incorrect, err: %s", err)
	}

	arrExpected = []interface{}{4, 1, 5}
	if reflect.DeepEqual(arr, arrExpected) == false {
		t.Errorf("Position in array were incorrect, got: %v, want: %v", arr, arrExpected)
	}

	lil.Delete(1)
	lil.Delete(5)
	lil.AddAtFront(3)
	lil.AddAtEnd(5)

	arr, err = lil.ToArray()
	if err != nil {
		t.Errorf("Position in array were incorrect, err: %s", err)
	}

	arrExpected = []interface{}{3, 4, 5}
	if reflect.DeepEqual(arr, arrExpected) == false {
		t.Errorf("Position in array were incorrect, got: %v, want: %v", arr, arrExpected)
	}
}
