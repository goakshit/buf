#!/usr/bin/env bash
if ! type buf >/dev/null 2>&1
then
    echo "buf tool is . You can install pre-commit using: brew install bufbuild/buf/buf"
    exit 1
fi
