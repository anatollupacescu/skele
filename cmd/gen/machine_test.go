package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIn(t *testing.T) {
	var obj = new(machine)

	obj.read("TESTGEN.pl")

	assert.Equal(t, "inventory", obj.specs[0].pkg)
	assert.Equal(t, "inventory_test.go", obj.specs[0].filename)

	assert.Equal(t, "create", obj.specs[0].fun[0].name)
	assert.Equal(t, "name", obj.specs[0].fun[0].arg)
	assert.Equal(t, "create item", obj.specs[0].fun[0].comm)

	assert.Equal(t, "nameempty", obj.specs[0].fun[0].pre[0].name)
	assert.Equal(t, "empty name", obj.specs[0].fun[0].pre[0].comm)
	assert.Equal(t, "nameduplicate", obj.specs[0].fun[0].pre[1].name)
	assert.Equal(t, "name present", obj.specs[0].fun[0].pre[1].comm)

	assert.Equal(t, "itemcreated", obj.specs[0].fun[0].pos[0].name)
	assert.Equal(t, "valid item | item is created", obj.specs[0].fun[0].pos[0].comm)
	assert.Equal(t, "itemnotcreated", obj.specs[0].fun[0].pos[1].name)
	assert.Equal(t, "valid item is not created | error", obj.specs[0].fun[0].pos[1].comm)

	assert.Equal(t, "second", obj.specs[0].fun[1].name)
	assert.Equal(t, "test", obj.specs[0].fun[1].arg)
	assert.Equal(t, "second item", obj.specs[0].fun[1].comm)

	assert.Equal(t, "second_empty", obj.specs[0].fun[1].pre[0].name)
	assert.Equal(t, "empty second", obj.specs[0].fun[1].pre[0].comm)
	assert.Equal(t, "second_duplicate", obj.specs[0].fun[1].pre[1].name)
	assert.Equal(t, "second present", obj.specs[0].fun[1].pre[1].comm)

	assert.Equal(t, "second_created", obj.specs[0].fun[1].pos[0].name)
	assert.Equal(t, "valid second | second is created", obj.specs[0].fun[1].pos[0].comm)
	assert.Equal(t, "second_notcreated", obj.specs[0].fun[1].pos[1].name)
	assert.Equal(t, "valid second is not created | error", obj.specs[0].fun[1].pos[1].comm)

	assert.Equal(t, "stock", obj.specs[1].pkg)
	assert.Equal(t, "stock_test.go", obj.specs[1].filename)

	assert.Equal(t, "provision", obj.specs[1].fun[0].name)
	assert.Equal(t, "position", obj.specs[1].fun[0].arg)
	assert.Equal(t, "provision item", obj.specs[1].fun[0].comm)

	assert.Equal(t, "name_empty", obj.specs[1].fun[0].pre[0].name)
	assert.Equal(t, "empty name", obj.specs[1].fun[0].pre[0].comm)
	assert.Equal(t, "name_duplicate", obj.specs[1].fun[0].pre[1].name)
	assert.Equal(t, "name present", obj.specs[1].fun[0].pre[1].comm)

	assert.Equal(t, "item_provisioned", obj.specs[1].fun[0].pos[0].name)
	assert.Equal(t, "valid item | item is provisioned", obj.specs[1].fun[0].pos[0].comm)
	assert.Equal(t, "item_notprovisioned", obj.specs[1].fun[0].pos[1].name)
	assert.Equal(t, "valid item is not provisioned | error", obj.specs[1].fun[0].pos[1].comm)

	assert.Equal(t, "sell", obj.specs[1].fun[1].name)
	assert.Equal(t, "item", obj.specs[1].fun[1].arg)
	assert.Equal(t, "sell item", obj.specs[1].fun[1].comm)

	assert.Equal(t, "third_empty", obj.specs[1].fun[1].pre[0].name)
	assert.Equal(t, "empty third", obj.specs[1].fun[1].pre[0].comm)
	assert.Equal(t, "third_duplicate", obj.specs[1].fun[1].pre[1].name)
	assert.Equal(t, "third present", obj.specs[1].fun[1].pre[1].comm)

	assert.Equal(t, "third_provisioned", obj.specs[1].fun[1].pos[0].name)
	assert.Equal(t, "valid third | third is provisioned", obj.specs[1].fun[1].pos[0].comm)
	assert.Equal(t, "third_notprovisioned", obj.specs[1].fun[1].pos[1].name)
	assert.Equal(t, "valid third is not provisioned | error", obj.specs[1].fun[1].pos[1].comm)
}
