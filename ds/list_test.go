package ds_test

import (
	ds "go-dsa/ds"
	"testing"
)

// Unit tests for List implementation
var testList = ds.NewIntegerList([]int{1, 3, 4, 5, 7, 8, 2, 1, 1, 4, 3, 6, 9})

func TestCount(t *testing.T) {
	expected := 13
	result := testList.Count()

	if result != expected {
		t.Errorf("List.Count() expected %d, found %d", expected, result)
	}
}

// Contains
func TestContains(t *testing.T) {
	expected := false
	result := testList.Contains(200)

	if result != expected {
		t.Errorf("List.Contains(200) expected %v, found %v", expected, result)
	}

	expected = true
	result = testList.Contains(2)

	if result != expected {
		t.Errorf("List.Contains(2) expected %v, found %v", expected, result)
	}

}

// IndexOf
func TestIndexOf(t *testing.T) {
	expected := 3
	result := testList.IndexOf(5)

	if result != expected {
		t.Errorf("List.IndexOf(5) expected %v, found %v", expected, result)
	}

	expected = 2
	result = testList.IndexOf(4)

	if result != expected {
		t.Errorf("List.IndexOf(4) expected %v, found %v", expected, result)
	}
}

// Add
func TestAdd(t *testing.T) {
	testList.Add(0)

	expectedInt := 13
	resultInt := testList.IndexOf(0)

	if resultInt != expectedInt {
		t.Errorf("Add:List.IndexOf(0) expected %v, found %v", expectedInt, resultInt)
	}

	expectedInt = 14
	resultInt = testList.Count()

	if resultInt != expectedInt {
		t.Errorf("Add:List.Count() expected %v, found %v", expectedInt, resultInt)
	}

	expectedBool := true
	resultBool := testList.Contains(0)

	if resultBool != expectedBool {
		t.Errorf("Add:List.IndexOf(0) expected %v, found %v", expectedBool, resultBool)
	}
}

// Insert
func TestInsert(t *testing.T) {
	testList.Insert(100, 4)

	expectedInt := 4
	resultInt := testList.IndexOf(100)

	if resultInt != expectedInt {
		t.Errorf("Insert:List.IndexOf(100) expected %v, found %v", expectedInt, resultInt)
	}

	expectedInt = 15
	resultInt = testList.Count()

	if resultInt != expectedInt {
		t.Errorf("Insert:List.Count() expected %v, found %v", expectedInt, resultInt)
	}

	expectedBool := true
	resultBool := testList.Contains(100)

	if resultBool != expectedBool {
		t.Errorf("Insert:List.IndexOf(100) expected %v, found %v", expectedBool, resultBool)
	}
}

func TestInsertDoesNotExist(t *testing.T) {
	testList.Insert(400, 250)

	expectedInt := 15
	resultInt := testList.Count()

	if resultInt != expectedInt {
		t.Errorf("Insert:List.Count() expected %v, found %v", expectedInt, resultInt)
	}
}

// Remove
func TestRemove(t *testing.T) {
	testList.Remove(100)

	expectedInt := -1
	resultInt := testList.IndexOf(100)

	if resultInt != expectedInt {
		t.Errorf("Remove:List.IndexOf(100) expected %v, found %v", expectedInt, resultInt)
	}

	expectedInt = 14
	resultInt = testList.Count()

	if resultInt != expectedInt {
		t.Errorf("Remove:List.Count() expected %v, found %v", expectedInt, resultInt)
	}

	expectedBool := false
	resultBool := testList.Contains(100)

	if resultBool != expectedBool {
		t.Errorf("Remove:List.IndexOf(100) expected %v, found %v", expectedBool, resultBool)
	}
}

func TestRemoveDoesNotExist(t *testing.T) {
	testList.Remove(12345)

	expected := 14
	result := testList.Count()

	if result != expected {
		t.Errorf("Remove:List.Count() expected %d, found %d", expected, result)
	}
}

// RemoveAt
func TestRemoveAt(t *testing.T) {
	testList.RemoveAt(2)

	expectedInt := 8
	resultInt := testList.IndexOf(4)

	if resultInt != expectedInt {
		t.Errorf("RemoveAt:List.IndexOf(4) expected %v, found %v", expectedInt, resultInt)
	}

	expectedInt = 13
	resultInt = testList.Count()

	if resultInt != expectedInt {
		t.Errorf("RemoveAt:List.Count() expected %v, found %v", expectedInt, resultInt)
	}

	expectedBool := true
	resultBool := testList.Contains(4)

	if resultBool != expectedBool {
		t.Errorf("RemoveAt:List.IndexOf(100) expected %v, found %v", expectedBool, resultBool)
	}
}

func TestRemoveAtDoesNotExist(t *testing.T) {
	testList.RemoveAt(200)

	expected := 13
	result := testList.Count()

	if result != expected {
		t.Errorf("RemoveAt:List.Count() expected %d, found %d", expected, result)
	}
}

// Clear
func TestClear(t *testing.T) {
	testList.Clear()

	expectedInt := -1
	resultInt := testList.IndexOf(100)

	if resultInt != expectedInt {
		t.Errorf("Clear:List.IndexOf(100) expected %v, found %v", expectedInt, resultInt)
	}

	expectedInt = 0
	resultInt = testList.Count()

	if resultInt != expectedInt {
		t.Errorf("Clear:List.Count() expected %v, found %v", expectedInt, resultInt)
	}

	expectedBool := false
	resultBool := testList.Contains(100)

	if resultBool != expectedBool {
		t.Errorf("Clear:List.IndexOf(100) expected %v, found %v", expectedBool, resultBool)
	}
}
