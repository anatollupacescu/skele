# Skele

Generates a go test skeleton based on a specification file.
Look into `cmd/skele/TODO.skl` for a reference spec file.

In order to get a sample output comment out the lines that start with `defer` in the `cmd/skele/gen_test` and the run:

```sh
cd cmd/skele
go test
```

## Usage

```sh
go get github.com/anatollupacescu/skele
skele def1.skl def2.skl
```

## Spec

A spec can contain multiple package definitions, each package definition contains an optional `fol` (folder) and `doc` statements.

A `fun` statement defines a function name, followed by zero or more `pre` (preconditions).

A precondition is a stop rule - if the precondition body evaluates to true - it means the flow of the program stops.

If two preconditions are on the same line - it means that the first precondition can fail because of technical reasons so the second one will generate the test for that case.

The last one is a set of one or more `pos` (postconditions) representing side efects.

They too can be two on the same line, for the same reason as the preconditions.

If a `pos` statement contains only one postcondition it means it's a return statement or a side effect that can not fail (an assignment, for ex.)

## TODO

* add support for comments
* add support for generating other types of test (goconvey, js mocha etc.)
* add check for duplicate function names within the same package
