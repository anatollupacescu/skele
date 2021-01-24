file inventory/inventory_test.go

fun create(name) # create item
pre
    \ nameempty # empty name
    \ nameduplicate # name present
pos
    \ itemcreated # valid item | item is created
    \ itemnotcreated # valid item is not created | error

fun second(test) # second item
pre
    \ second_empty # empty second
    \ second_duplicate # second present
pos
    \ second_created # valid second | second is created
    \ second_notcreated # valid second is not created | error

file stock/stock_test.go

fun provision(position) # provision item
pre
    \ name_empty # empty name
    \ name_duplicate # name present
pos
    \ item_provisioned # valid item | item is provisioned
    \ item_notprovisioned # valid item is not provisioned | error

fun sell(item) # sell item
pre
    \ third_empty # empty third
    \ third_duplicate # third present
pos
    \ third_provisioned # valid third | third is provisioned
    \ third_notprovisioned # valid third is not provisioned | error
