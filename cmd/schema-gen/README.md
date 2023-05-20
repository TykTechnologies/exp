# schema-gen

The `schema-gen` tool analyzes a single go source code package and dumps
the data model into a json file. The ultimate goal of the tool is to
extract a data model from any package enough to:

- create a new package with only the data model types,
- enable generating documentation based on source code fields,
- enable further rich outputs outputs like jsonschema.

To install:

`go install github.com/TykTechnologies/exp/cmd/schema-gen@main`

Invoke `schema-gen` without arguments to print usage help.

Usage:

- `schema-gen`
- `schema-gen extract help`
- `schema-gen extract -i model/ -o model/model.json`
- ...

Example:

See the `example/` subfolder.

## Random facts

- we exclude `_` fields,
- we exclude unexposed fields,
- run `task` to run everything

## People to talk to

Slack #team-ext-manage-squad, Tit Petric
