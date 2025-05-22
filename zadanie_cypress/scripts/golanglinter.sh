#!/bin/bash
echo "Running golangci-lint at zadanie_cypress/backend/ ..."

cd zadanie_cypress/backend || exit 1

LINT_OUTPUT=$(golangci-lint run ./... 2>&1)
STATUS=$?

if [ $STATUS -ne 0 ]; then
  echo "$LINT_OUTPUT"
  echo "Golang linter failed failed. Commit aborted."
  exit 1
else
  echo "Golang linter passed."
fi
