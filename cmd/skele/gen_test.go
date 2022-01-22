package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {
	m := new(machine)

	m.read("TODO.skl")
	m.write()

	defer removeFile("list")

	item, err := os.Open("item_test.go")
	if err != nil {
		t.Fatal(err)
	}

	defer removeFile("item_test.go")

	data, err := ioutil.ReadAll(item)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, itemFileContents, string(data))

	list, err := os.Open("list/list_test.go")
	if err != nil {
		t.Fatal(err)
	}

	defer removeFile("list/list_test.go")

	data, err = ioutil.ReadAll(list)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, listFileContents, string(data))

	defer removeFile("list/doc.go")
	defer removeFile("doc.go")
}

func removeFile(path string) {
	if e := os.Remove(path); e != nil {
		log.Fatal(e)
	}
}

var itemFileContents = `package item

import "testing"

func TestAddItemToList(t *testing.T) { 
	t.Run("given text too short", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
	t.Run("given text not unique in list", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
	t.Run("given fail to check for uniqueness", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
	t.Run("given item added", func(t *testing.T) {
		t.Run("assert success", func(t *testing.T) {
		})
	})
	t.Run("given fail to add item", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
}

func TestMarkItemAsDone(t *testing.T) { 
	t.Run("given item marked as done", func(t *testing.T) {
		t.Run("assert success", func(t *testing.T) {
		})
	})
	t.Run("given fail to mark as done", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
}

func TestRemoveItem(t *testing.T) { 
	t.Run("given item is active", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
	t.Run("given fail to check item status", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
	t.Run("given item removed", func(t *testing.T) {
		t.Run("assert success", func(t *testing.T) {
		})
	})
	t.Run("given fail to remove item", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
}
`

var listFileContents = `package list

import "testing"

func TestAddList(t *testing.T) { 
	t.Run("given name invalid", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
	t.Run("given name not unique", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
	t.Run("given fail to check name uniqueness", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
	t.Run("given list added", func(t *testing.T) {
		t.Run("assert success", func(t *testing.T) {
		})
	})
	t.Run("given fail to add list", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
}

func TestRemoveList(t *testing.T) { 
	t.Run("given contains pending items", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
	t.Run("given list removed", func(t *testing.T) {
		t.Run("assert success", func(t *testing.T) {
		})
	})
	t.Run("given fail to remove list", func(t *testing.T) {
		t.Run("assert error", func(t *testing.T) {
		})
	})
}
`
