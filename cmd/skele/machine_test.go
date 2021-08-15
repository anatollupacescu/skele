package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTODO(t *testing.T) {
	var m = new(machine)

	m.read("TODO.skl")

	// list
	assert.Equal(t, "list", m.pkgs[0].name)
	assert.Equal(t, "list", m.pkgs[0].fol)

	assert.Equal(t, "A list is a named collection of reminders (TODO items). A set of such lists constitutes the ToDoApp.", m.pkgs[0].doc[0])
	assert.Equal(t, "A list can be added if the name is unique and not empty. A list can be removed if all of its items are 'done'", m.pkgs[0].doc[1])

	assert.Equal(t, "list_test.go", m.pkgs[0].file[0].name)

	assert.Equal(t, "add list", m.pkgs[0].file[0].fun[0].name)

	assert.Equal(t, "name too short", m.pkgs[0].file[0].fun[0].pre[0].domain)
	assert.Equal(t, "", m.pkgs[0].file[0].fun[0].pre[0].impl)

	assert.Equal(t, "name not unique", m.pkgs[0].file[0].fun[0].pre[1].domain)
	assert.Equal(t, "fail to check name uniqueness", m.pkgs[0].file[0].fun[0].pre[1].impl)

	assert.Equal(t, "list added", m.pkgs[0].file[0].fun[0].pos[0].domain)
	assert.Equal(t, "fail to add list", m.pkgs[0].file[0].fun[0].pos[0].impl)

	assert.Equal(t, "remove list", m.pkgs[0].file[0].fun[1].name)

	assert.Equal(t, "contains pending items", m.pkgs[0].file[0].fun[1].pre[0].domain)
	assert.Equal(t, "", m.pkgs[0].file[0].fun[1].pre[0].impl)

	assert.Equal(t, "list removed", m.pkgs[0].file[0].fun[1].pos[0].domain)
	assert.Equal(t, "fail to remove list", m.pkgs[0].file[0].fun[1].pos[0].impl)

	// item
	assert.Equal(t, "item", m.pkgs[1].name)
	assert.Equal(t, "", m.pkgs[1].fol)
	assert.Equal(t, "item_test.go", m.pkgs[1].file[0].name)

	assert.Equal(t, "An item represents a TODO reminder comprised of a text and a status.", m.pkgs[1].doc[0])
	assert.Equal(t, "An item can be removed only if it's in 'done' status.", m.pkgs[1].doc[1])

	assert.Equal(t, "add item to list", m.pkgs[1].file[0].fun[0].name)
	assert.Equal(t, "text too short", m.pkgs[1].file[0].fun[0].pre[0].domain)
	assert.Equal(t, "", m.pkgs[1].file[0].fun[0].pre[0].impl)

	assert.Equal(t, "text not unique in list", m.pkgs[1].file[0].fun[0].pre[1].domain)
	assert.Equal(t, "fail to check for uniqueness", m.pkgs[1].file[0].fun[0].pre[1].impl)

	assert.Equal(t, "item added", m.pkgs[1].file[0].fun[0].pos[0].domain)
	assert.Equal(t, "fail to add item", m.pkgs[1].file[0].fun[0].pos[0].impl)

	assert.Equal(t, "mark item as done", m.pkgs[1].file[0].fun[1].name)

	assert.Equal(t, "item marked as done", m.pkgs[1].file[0].fun[1].pos[0].domain)
	assert.Equal(t, "fail to mark as done", m.pkgs[1].file[0].fun[1].pos[0].impl)

	assert.Equal(t, "remove item", m.pkgs[1].file[0].fun[2].name)

	assert.Equal(t, "item is active", m.pkgs[1].file[0].fun[2].pre[0].domain)
	assert.Equal(t, "fail to check item status", m.pkgs[1].file[0].fun[2].pre[0].impl)

	assert.Equal(t, "item removed", m.pkgs[1].file[0].fun[2].pos[0].domain)
	assert.Equal(t, "fail to remove item", m.pkgs[1].file[0].fun[2].pos[0].impl)
}
