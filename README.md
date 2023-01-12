#### Overview
Sample project to try buf!

#### IMPORTANT
> `release` branch should be an orphan branch.

To create orphan branch with empty commit:
``` bash
git checkout --orphan release
git commit --allow-empty -m "initial commit"
git push origin release
```

#### Missing
- Auto tagging every release.
- `buf generate` command to generate files.
