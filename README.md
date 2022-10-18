# ttg

**T**ravel **T**erra**g**runt directory tree as a first class passenger! :sunglasses:

## How to use?

* `cd` to the directory of your [locally cloned] Terragrunt/Terraform Git repo;
* run `ttg` command there :rocket: ([optional] arguments are "path filter" matches);
* use arrow keys to navigate the list and `/` key to search for specific projects;

## Shell aliases

It is **highly** recommended to use `bash` (or `zsh`) aliases. Start from something like this:
```
alias ttg='_ttg(){ ttg -outFile ~/.ttg-path ${@} && cd "$(cat ~/.ttg-path)" }; _ttg'
alias ttt='ttg --top'
```

:bulb: `ttg --top` is a "shortcut" that brings you to the top level path of your repository.

### Why aliases?
Core aspect of this program is the ability to change working directory while staying **inside the current shell**.
This can not be done by the program itself, because of obvious security related `POSIX` limitations. Without instrumenting
the shell with aliases `ttg` still can kinda work, but will provide you with much more awkward and second class user
experience, i.e. you will need to exit subshell before you "jump" to the next project. :weary:

## How to build?

* `make dep` - install dependencies;
* `make build` - build the `ttg` binary in `cmd/ttg` path;
* `make install` - [optional] install built `ttg` binary under the `${PREFIX}/bin` location;

## How to release a new version?

:bulb: Set `GITHUB_TOKEN` environmental variable.

Run `make full-release` recipe, which is equal to run following, one by one:

* `make clean` - cleanup project tree from previously built artifacts;
* `make dep` - ensure all dependencies are installed;
* `make release` - create release artifacts;
* `make next-version-tag` - tag your HEAD with the incremented version;
* `make push-tags` - push your new tag to Git;
* `make github` - create a GitHub release using your artifacts;
