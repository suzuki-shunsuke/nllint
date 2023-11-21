# nllint

Linter to check newlines at the end of files

```console
$ echo -n {} > foo.json
$ echo -n "{}\n\n" > bar.json

$ nllint foo.json bar.json
ERRO[0000] a file is invalid                             error="a newline at the end of file is missing" file_path=foo.json nllint_version= program=nllint

# --empty-line (-e): Check leading and trailing white spaces in files
$ nllint -e foo.json bar.json
ERRO[0000] a file is invalid                             error="a newline at the end of file is missing" file_path=foo.json nllint_version= program=nllint
ERRO[0000] a file is invalid                             error="empty lines at the end of file should be trimmed" file_path=bar.json nllint_version= program=nllint

# -fix (-f): Fix files and outputs fixed file paths to the stdout
$ nllint -e -fix foo.json bar.json
WARN[0000] a newline at the end of file is missing       file_path=foo.json nllint_version= program=nllint
foo.json
WARN[0000] empty lines at the end of file should be trimmed  file_path=bar.json nllint_version= program=nllint
bar.json
```

## Motivation

Each line should have a newline of the end of line because [that’s how the POSIX standard defines a line](https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap03.html#tag_03_206):

Please see https://stackoverflow.com/a/729795/6364492

## Install

Coming soon.

## Usage

```sh
$ nllint help
nllint - Check newlines at the end of files

https://github.com/suzuki-shunsuke/nllint

Usage:
  nllint [-fix (-f)] [-trim-space (-s)] <file path> [<file path>...]

Options:
  -help, -h        Show help
  -version, -v     Show version
  -fix, -f         Fix files
  -trim-space, -s  Disallow leading and trailing white spaces in files
```

## :bulb: Auto Fix by CI

It's useful to format code automatically with nllint and push a commit to the remote branch in CI.

1. List changed files
1. Run `nllint -f [-e] [<changed files>...]`
1. Push a commit to the remote branch

Please refer to this repository's workflows as the example.

- [test.yaml](.github/workflows/test.yaml)
- [wc-test.yaml](.github/workflows/wc-test.yaml)

## LICENSE

[MIT](LICENSE)
