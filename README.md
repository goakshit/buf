#### Overview [test]
Sample project to try buf to compile protobuffers.

#### IMPORTANT
- `release` branch should be an orphan branch.

To create orphan branch with empty commit:
``` bash
git checkout --orphan release
git commit --allow-empty -m "initial commit"
git push origin release
```

- Auto tagging every release. 
  Defaults to: **#major**(X.x.x), **#minor**(x.X.x) & **#patch**(x.x.X)
  If none of the tags are present in commit message, new tags are ignored(defaults to minor tag; configurable).

- Added Workflow for verifying buf linting, breaking change ... and so on.
- `buf generate` command to generate files.
