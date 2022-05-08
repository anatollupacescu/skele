package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTODO(t *testing.T) {
	m := new(machine)

	m.read("TODO.skl")

	// list
	pkg_list := m.pkgs[0]
	assert.Equal(t, "list", pkg_list.name)
	assert.Equal(t, "list", pkg_list.fol)

	assert.Equal(t, "A list is a named collection of reminders, or \"TODO items\". A set of such lists constitutes the ToDoApp.", m.pkgs[0].doc[0])
	assert.Equal(t, "A list can be added if the name is unique and not empty. A list can be removed if all of its items are 'done'", m.pkgs[0].doc[1])

	assert.Equal(t, "list_test.go", pkg_list.file[0].name)

	// fsm
	assert.Equal(t, "list", pkg_list.fsm[0].name)
	assert.Equal(t, "active", pkg_list.fsm[0].states[0])
	assert.Equal(t, "removed", pkg_list.fsm[0].states[1])

	// add list fun
	fun_add_list := pkg_list.file[0].fun[0]
	assert.Equal(t, "add list", fun_add_list.name)

	assert.Equal(t, "name invalid", fun_add_list.pre[0].succ)
	assert.Equal(t, "", fun_add_list.pre[0].fail)
	assert.Equal(t, "name is too short", fun_add_list.pre[0].tcS[0])
	assert.Equal(t, "name starts with a number, assert bad", fun_add_list.pre[0].tcS[1])

	assert.Equal(t, "name not unique", fun_add_list.pre[1].succ)
	assert.Equal(t, "failure to check name uniqueness", fun_add_list.pre[1].fail)
	assert.Equal(t, "one", fun_add_list.pre[1].tcF[0])
	assert.Equal(t, "two, assert not great", fun_add_list.pre[1].tcF[1])

	assert.Equal(t, "list added", fun_add_list.pos[0].succ)
	assert.Equal(t, "failure to add list", fun_add_list.pos[0].fail)
	assert.Equal(t, "failed to notify", fun_add_list.pos[0].tcF[0])
	assert.Equal(t, "could not persist, assert correct reason, assert error", fun_add_list.pos[0].tcF[1])
	assert.Equal(t, "list", fun_add_list.pos[0].fsm[0].name)
	assert.Equal(t, "active", fun_add_list.pos[0].fsm[0].state)

	fun_remove_list := pkg_list.file[0].fun[1]
	assert.Equal(t, "remove list", fun_remove_list.name)
	assert.Equal(t, "list", fun_remove_list.fsm.name)
	assert.Equal(t, "active", fun_remove_list.fsm.state)

	assert.Equal(t, "has pending items", fun_remove_list.pre[0].succ)
	assert.Equal(t, "", fun_remove_list.pre[0].fail)

	assert.Equal(t, "item", fun_remove_list.pre[0].fsm[0].name)
	assert.Equal(t, "done", fun_remove_list.pre[0].fsm[0].state)

	assert.Equal(t, "list removed", fun_remove_list.pos[0].succ)
	assert.Equal(t, "failure to remove list", fun_remove_list.pos[0].fail)

	assert.Equal(t, "list", fun_remove_list.pos[0].fsm[0].name)
	assert.Equal(t, "removed", fun_remove_list.pos[0].fsm[0].state)

	// item
	pkg_item := m.pkgs[1]
	assert.Equal(t, "item", pkg_item.name)
	assert.Equal(t, "", pkg_item.fol)
	assert.Equal(t, "item_test.go", pkg_item.file[0].name)

	// fsm
	assert.Equal(t, "item", pkg_item.fsm[0].name)
	assert.Equal(t, "active", pkg_item.fsm[0].states[0])
	assert.Equal(t, "done", pkg_item.fsm[0].states[1])
	assert.Equal(t, "removed", pkg_item.fsm[0].states[2])

	assert.Equal(t, "An item represents a TODO reminder comprised of a text and a status.", pkg_item.doc[0])
	assert.Equal(t, "An active item can be marked as 'done' and only after that it can be removed.", pkg_item.doc[1])

	fun_add_item_to_list := pkg_item.file[0].fun[0]
	assert.Equal(t, "add item to list", fun_add_item_to_list.name)
	assert.Equal(t, "text too short", fun_add_item_to_list.pre[0].succ)
	assert.Equal(t, "", fun_add_item_to_list.pre[0].fail)

	assert.Equal(t, "text not unique in list", fun_add_item_to_list.pre[1].succ)
	assert.Equal(t, "failure to check for uniqueness", fun_add_item_to_list.pre[1].fail)

	assert.Equal(t, "item added", fun_add_item_to_list.pos[0].succ)
	assert.Equal(t, "failure to add item", fun_add_item_to_list.pos[0].fail)

	fun_mark_item_as_done := pkg_item.file[0].fun[1]
	assert.Equal(t, "mark item as done", fun_mark_item_as_done.name)

	assert.Equal(t, "item marked as done", fun_mark_item_as_done.pos[0].succ)
	assert.Equal(t, "failure to mark as done", fun_mark_item_as_done.pos[0].fail)

	fun_remove_item := pkg_item.file[0].fun[2]
	assert.Equal(t, "remove item", fun_remove_item.name)

	assert.Equal(t, "item is active", fun_remove_item.pre[0].succ)
	assert.Equal(t, "failure to check item status", fun_remove_item.pre[0].fail)

	assert.Equal(t, "item removed", fun_remove_item.pos[0].succ)
	assert.Equal(t, "failure to remove item", fun_remove_item.pos[0].fail)
}
