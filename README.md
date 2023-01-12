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

- Auto tagging every release [Understand naming commit message]
> Setup another github workflow that bumps the tag version on the release branch.
> To bump the version, commit message should be formulated properly. It follows (Angular convention)[https://github.com/angular/angular/blob/main/CONTRIBUTING.md#-commit-message-format]

#### Missing
- Use existing buf actions for linting, breaking change ... and so on.
- `buf generate` command to generate files.
