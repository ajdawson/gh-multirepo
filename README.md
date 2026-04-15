# gh-multirepo

This is a `gh` extension that provides a wrapper command to run the same `gh`
command across multiple repositories. Some `gh` commands support multiple
occurrences of `-R`, but others do not. This wrapper simply executes the given
command for each repository serially. You can only us `multirepo` to run `gh`
commands that accept a `-R` option.

## Install

```sh
gh extension install ajdawson/gh-multirepo
```

## Usage

The command takes a comma-separated list of repositories to work on, and a `gh`
command that supports the `-R` option to run against all of them. It is
recommended to always use `--` to separate the command from other arguments,
although it is only necessary if your command has option flags:

    gh multirepo cli/cli,cli/go-gh -- pr list

The `--` is necessary when your command uses option flags:

    gh multirepo cli/cli,cli/go-gh -- pr list --json title

By default the `multirepo` command prints a header consisting of the repository
name it is working on before the results for each repository. If you don't want
that you can turn it off with the `--no-header` flag. If you want a header but
don't like the default then you can provide a go template by using the
`--header-template` option, the only accepted variable for the template is
`.Repo`.
