# nllint

Linter and formatter of newlines and trailing spaces in files

<img width="870" alt="image" src="https://github.com/suzuki-shunsuke/nllint/assets/13323303/6962481e-d74e-47a6-bdd0-cd31edae1771">

--

<img width="982" alt="image" src="https://github.com/suzuki-shunsuke/nllint/assets/13323303/5cc5cf3f-874b-4465-834c-5bba55285fa8">

```console
$ echo -n {} > foo.json
$ echo -n "{}\n\n" > bar.json

$ nllint foo.json bar.json
ERRO[0000] a file is invalid                             error="a newline at the end of file is missing" file_path=foo.json nllint_version= program=nllint

# --trim-space (-s): Check leading and trailing white spaces in files
$ nllint -s foo.json bar.json
ERRO[0000] a file is invalid                             error="a newline at the end of file is missing" file_path=foo.json nllint_version= program=nllint
ERRO[0000] a file is invalid                             error="leading and trailing white spaces in files should be trimmed" file_path=bar.json nllint_version= program=nllint

# -fix (-f): Fix files and outputs fixed file paths to the stdout
$ nllint -s -fix foo.json bar.json
WARN[0000] a newline at the end of file is missing       file_path=foo.json nllint_version= program=nllint
foo.json
WARN[0000] leading and trailing white spaces in files should be trimmed  file_path=bar.json nllint_version= program=nllint
bar.json
```

## Motivation

Each line should have a newline of the end of line because [that’s how the POSIX standard defines a line](https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap03.html#tag_03_206):

Please see https://stackoverflow.com/a/729795/6364492

## Install

nllint is a single binary written in [Go](https://go.dev/). So you only need to install an executable file into `$PATH`.

1. [Homebrew](https://brew.sh/)

```sh
brew install suzuki-shunsuke/nllint/nllint
```

2. [Scoop](https://scoop.sh/)

```sh
scoop bucket add suzuki-shunsuke https://github.com/suzuki-shunsuke/scoop-bucket
scoop install nllint
```

3. [aqua](https://aquaproj.github.io/)

```sh
aqua g -i suzuki-shunsuke/nllint
```

4. Download a prebuilt binary from [GitHub Releases](https://github.com/suzuki-shunsuke/nllint/releases) and install it into `$PATH`

## Usage

```sh
$ nllint help
nllint - Check newlines at the end of files

https://github.com/suzuki-shunsuke/nllint

Usage:
  nllint [-fix (-f)] [-trim-space (-s)] [-trim-trailing-space (-S)] [-ignore-notfound (-i)] <file path> [<file path>...]

Options:
  -help, -h                 Show help
  -version, -v              Show version
  -fix, -f                  Fix files
  -trim-space, -s           Disallow leading and trailing white spaces in files
  -trim-trailing-space, -S  Disallow trailing white spaces in each line
  -ignore-notfound, -i      Ignore not found files
```

## :bulb: Auto Fix by CI

<img width="870" alt="image" src="https://github.com/suzuki-shunsuke/nllint/assets/13323303/6962481e-d74e-47a6-bdd0-cd31edae1771">

--

<img width="982" alt="image" src="https://github.com/suzuki-shunsuke/nllint/assets/13323303/5cc5cf3f-874b-4465-834c-5bba55285fa8">

It's useful to format code automatically with nllint and push a commit to the remote branch in CI.

1. List changed files
1. Run `nllint -f [-s] [<changed files>...]`
1. Push a commit to the remote branch

Please refer to this repository's workflows as the example.

- [test.yaml](.github/workflows/test.yaml)
- [wc-test.yaml](.github/workflows/wc-test.yaml)

`nllint` is enough fast, so we think it's also okay to lint all files instead of only changed files.

```sh
git ls-files | xargs nllint -f -s
```

## LICENSE

[MIT](LICENSE)
