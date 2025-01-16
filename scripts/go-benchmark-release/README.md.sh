#!/bin/bash
cat README.md.tpl
echo
echo "# Running everything"
echo
task -l | sed 's/:$/:\n/g'
echo

TASKS=$(task -l --json | jq .tasks[].name -r)
for NAME in $TASKS; do
	echo "## $(task ${NAME} --summary)"
	echo
done
