# schema-gen

The `schema-gen` tool analyzes a single go source code package and dumps
the data model into a json file. The ultimate goal of the tool is to
extract a data model from any package enough to:

- create a new package with only the data model types,
- enable generating documentation based on source code fields,
- enable further rich outputs outputs like jsonschema.

To install:

`go install github.com/TykTechnologies/exp/cmd/schema-gen@main`

Arguments:

- `-i` - input folder of go package, defaults to `.`, needs trailing `/`,
- `-o` - output file, defaults to `schema.json`.

Usage:

- `schema`
- `schema-gen -i structs/ -o structs.json`

The structs.json example is commited to the repo.

## Random facts

- we exclude `_` fields,
- we exclude unexposed fields,
- run `task` to run everything

## People to talk to

Slack #team-ext-manage-squad, Tit Petric
