package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeFile(path string) {
	e := os.Remove(path)
	if e != nil {
		log.Fatal(e)
	}
}

func TestGen(t *testing.T) {
	var m = new(machine)

	m.read("TESTGEN.pl")
	m.write()

	file, err := os.Open("inventory/inventory_test.go")
	if err != nil {
		t.Fatal(err)
	}

	defer removeFile("inventory/inventory_test.go")

	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, file1, string(data))

	file, err = os.Open("stock/stock_test.go")
	if err != nil {
		t.Fatal(err)
	}

	defer removeFile("stock/stock_test.go")

	data, err = ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, file2, string(data))
}

var file1 = `package inventory

import "testing"

func TestCreateItem(t *testing.T) { 
	t.Run("given empty name", func(t *testing.T) {

		t.Run("assert error", func(t *testing.T) {

		})
	})
	t.Run("given name present", func(t *testing.T) {

		t.Run("assert error", func(t *testing.T) {

		})
	})
	t.Run("given valid item", func(t *testing.T) {

		t.Run("assert item is created", func(t *testing.T) {

		})
	})
	t.Run("given valid item is not created", func(t *testing.T) {

		t.Run("assert error", func(t *testing.T) {

		})
	})
}

func TestSecondItem(t *testing.T) { 
	t.Run("given empty second", func(t *testing.T) {

		t.Run("assert error", func(t *testing.T) {

		})
	})
	t.Run("given second present", func(t *testing.T) {

		t.Run("assert error", func(t *testing.T) {

		})
	})
	t.Run("given valid second", func(t *testing.T) {

		t.Run("assert second is created", func(t *testing.T) {

		})
	})
	t.Run("given valid second is not created", func(t *testing.T) {

		t.Run("assert error", func(t *testing.T) {

		})
	})
}

`

var file2 = `package stock

import "testing"

func TestProvisionItem(t *testing.T) { 
	t.Run("given empty name", func(t *testing.T) {

		t.Run("assert error", func(t *testing.T) {

		})
	})
	t.Run("given name present", func(t *testing.T) {

		t.Run("assert error", func(t *testing.T) {

		})
	})
	t.Run("given valid item", func(t *testing.T) {

		t.Run("assert item is provisioned", func(t *testing.T) {

		})
	})
	t.Run("given valid item is not provisioned", func(t *testing.T) {

		t.Run("assert error", func(t *testing.T) {

		})
	})
}

func TestSellItem(t *testing.T) { 
	t.Run("given empty third", func(t *testing.T) {

		t.Run("assert error", func(t *testing.T) {

		})
	})
	t.Run("given third present", func(t *testing.T) {

		t.Run("assert error", func(t *testing.T) {

		})
	})
	t.Run("given valid third", func(t *testing.T) {

		t.Run("assert third is provisioned", func(t *testing.T) {

		})
	})
	t.Run("given valid third is not provisioned", func(t *testing.T) {

		t.Run("assert error", func(t *testing.T) {

		})
	})
}

`
