# Skele

Generates a go test skeleton based on a specification file.
See `EXAMPLE.skl` or `cmd/skele/TODO.skl` for a reference spec file.

In order to get a sample output comment out the lines that start with `defer` in the `cmd/skele/gen_test` and then run:

```sh
cd cmd/skele
go test
```

## Usage

```sh
go get github.com/anatollupacescu/skele/cmd/skele
skele def1.skl def2.skl
```

## Spec

A spec can contain multiple package definitions, each package definition contains an optional `fol` (folder) and `doc` statements.

A `fun` statement defines a function name, followed by zero or more `pre` (preconditions).

A precondition is a stop rule - if the precondition body evaluates to true - it means the flow of the program stops.

Two preconditions can be defined the same line. The first one represents a real life invariant and the second one represents a failure that occurred for a technical reason.

The last one is a set of one or more `pos` (postconditions) representing side efects.

Postconditions can also be two on the same line, for the same reason as the preconditions.

If a `pos` statement contains only one postcondition it means it's a return statement or a side effect that can not fail (an assignment for ex.)

## Comments

### State machine

To declare a state machine use `fsm` keyward, place before any function declarations.
In the comments for the function we can declare the start state for a specific machine.
In the comments for the preconditions we can declare the dependencies.
In the comments for the postconditions we can declare the final states for the transition

...WIP...

### Test cases and assertions

An assertion `arn{}` is ment to override the default `assert success` or `assert fail` constructs.
By default it will override the success assertion, however we can explicitly add the `<`, `>` and `<>` qualifiers to indicate that we want to override the success, fail or both constructs.

A test case `tcs{}` declares a subtest. By default it will be declared for the success path, but we can explicitly add the `<` and `>` qualifiers to indicate that we want to override the success or fail constructs.
A test cases can contain zero or more assertions...

## TODO

* add support for generating other types of testing frameworks (goconvey, js mocha etc.)
