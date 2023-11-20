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

## Install

Coming soon.

## Usage

```sh
$ nllint help
nllint - Check newlines at the end of files

https://github.com/suzuki-shunsuke/nllint

Usage:
  nllint [-fix] [-empty-line] <file path> [<file path>...]

Options:
  -help, -h        Show help
  -version, -v     Show version
  -fix, -f         Fix files
  -empty-line, -e  Disallow leading and trailing white spaces in files
```

## LICENSE

[MIT](LICENSE)
