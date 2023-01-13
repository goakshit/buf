#### Overview
Sample project to try buf!

#### IMPORTANT
- `release` branch should be an orphan branch.

To create orphan branch with empty commit:
``` bash
git checkout --orphan release
git commit --allow-empty -m "initial commit"
git push origin release
```

- Auto tagging every release. Defaults to: **#major**(X.x.x), **#minor**(x.X.x) & **#patch**(x.x.X)

#### Missing
- Use existing buf actions for linting, breaking change ... and so on.
- `buf generate` command to generate files.
