package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTODO(t *testing.T) {
	var m = new(machine)

	m.read("TODO.skl")

	// list
	assert.Equal(t, "list", m.specs[0].pkg)
	assert.Equal(t, "list", m.specs[0].fol)

	assert.Equal(t, "A list is a bag of TODO items. A collection of lists constitutes the ToDoApp.", m.specs[0].doc[0])
	assert.Equal(t, "A list can be added if the name is unique and not empty.", m.specs[0].doc[1])
	assert.Equal(t, "A list can be removed if all of its items are 'done'", m.specs[0].doc[2])

	assert.Equal(t, "list_test.go", m.specs[0].file[0].name)

	assert.Equal(t, "add list", m.specs[0].file[0].fun[0].name)

	assert.Equal(t, "name too short", m.specs[0].file[0].fun[0].pre[0].domain)
	assert.Equal(t, "", m.specs[0].file[0].fun[0].pre[0].impl)

	assert.Equal(t, "name not unique", m.specs[0].file[0].fun[0].pre[1].domain)
	assert.Equal(t, "fail to check name uniqueness", m.specs[0].file[0].fun[0].pre[1].impl)

	assert.Equal(t, "list added", m.specs[0].file[0].fun[0].pos[0].domain)
	assert.Equal(t, "fail to add list", m.specs[0].file[0].fun[0].pos[0].impl)

	assert.Equal(t, "remove list", m.specs[0].file[0].fun[1].name)

	assert.Equal(t, "contains pending items", m.specs[0].file[0].fun[1].pre[0].domain)
	assert.Equal(t, "", m.specs[0].file[0].fun[1].pre[0].impl)

	assert.Equal(t, "list removed", m.specs[0].file[0].fun[1].pos[0].domain)
	assert.Equal(t, "fail to remove list", m.specs[0].file[0].fun[1].pos[0].impl)

	// item
	assert.Equal(t, "item", m.specs[1].pkg)
	assert.Equal(t, "", m.specs[1].fol)
	assert.Equal(t, "item_test.go", m.specs[1].file[0].name)

	assert.Equal(t, "An item represents a TODO reminder. Contains the text and a 'done' toggle.", m.specs[1].doc[0])
	assert.Equal(t, "An item can be removed only if it's in 'done' state.", m.specs[1].doc[1])

	assert.Equal(t, "add item to list", m.specs[1].file[0].fun[0].name)
	assert.Equal(t, "text too short", m.specs[1].file[0].fun[0].pre[0].domain)
	assert.Equal(t, "", m.specs[1].file[0].fun[0].pre[0].impl)

	assert.Equal(t, "text not unique in list", m.specs[1].file[0].fun[0].pre[1].domain)
	assert.Equal(t, "fail to check for uniqueness", m.specs[1].file[0].fun[0].pre[1].impl)

	assert.Equal(t, "item added", m.specs[1].file[0].fun[0].pos[0].domain)
	assert.Equal(t, "fail to add item", m.specs[1].file[0].fun[0].pos[0].impl)

	assert.Equal(t, "mark item as done", m.specs[1].file[0].fun[1].name)

	assert.Equal(t, "item marked as done", m.specs[1].file[0].fun[1].pos[0].domain)
	assert.Equal(t, "fail to mark as done", m.specs[1].file[0].fun[1].pos[0].impl)

	assert.Equal(t, "remove item", m.specs[1].file[0].fun[2].name)

	assert.Equal(t, "item is active", m.specs[1].file[0].fun[2].pre[0].domain)
	assert.Equal(t, "fail to check item status", m.specs[1].file[0].fun[2].pre[0].impl)

	assert.Equal(t, "item removed", m.specs[1].file[0].fun[2].pos[0].domain)
	assert.Equal(t, "fail to remove item", m.specs[1].file[0].fun[2].pos[0].impl)
}
