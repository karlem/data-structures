package main

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
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

func assertStackPop(st *Stack, expected interface{}) error {
	item, err := st.Pop()
	if err != nil {
		return errors.Errorf("Popped item error occur, err: %s", err)
	}

	if item != expected {
		return errors.Errorf("Popped item was incorrect, got: %v, want: %v", item, expected)
	}

	return nil
}

func assertStackEmpty(st *Stack, expected bool) error {
	if isEmpty := st.IsEpmty(); isEmpty != expected {
		return errors.Errorf("IsEpmty was incorrect, got: %v, want: %v", isEmpty, expected)
	}

	return nil
}

func TestStack(t *testing.T) {
	st := NewStack()

	if err := assertStackEmpty(st, true); err != nil {
		t.Error(err)
	}

	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)

	if err := assertStackEmpty(st, false); err != nil {
		t.Error(err)
	}

	if err := assertStackPop(st, 4); err != nil {
		t.Error(err)
	}
	if err := assertStackPop(st, 3); err != nil {
		t.Error(err)
	}
	if err := assertStackPop(st, 2); err != nil {
		t.Error(err)
	}
	if err := assertStackPop(st, 1); err != nil {
		t.Error(err)
	}

	if err := assertStackEmpty(st, true); err != nil {
		t.Error(err)
	}
}
