# nllint

[![DeepWiki](https://img.shields.io/badge/DeepWiki-suzuki--shunsuke%2Fnllint-blue.svg?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACwAAAAyCAYAAAAnWDnqAAAAAXNSR0IArs4c6QAAA05JREFUaEPtmUtyEzEQhtWTQyQLHNak2AB7ZnyXZMEjXMGeK/AIi+QuHrMnbChYY7MIh8g01fJoopFb0uhhEqqcbWTp06/uv1saEDv4O3n3dV60RfP947Mm9/SQc0ICFQgzfc4CYZoTPAswgSJCCUJUnAAoRHOAUOcATwbmVLWdGoH//PB8mnKqScAhsD0kYP3j/Yt5LPQe2KvcXmGvRHcDnpxfL2zOYJ1mFwrryWTz0advv1Ut4CJgf5uhDuDj5eUcAUoahrdY/56ebRWeraTjMt/00Sh3UDtjgHtQNHwcRGOC98BJEAEymycmYcWwOprTgcB6VZ5JK5TAJ+fXGLBm3FDAmn6oPPjR4rKCAoJCal2eAiQp2x0vxTPB3ALO2CRkwmDy5WohzBDwSEFKRwPbknEggCPB/imwrycgxX2NzoMCHhPkDwqYMr9tRcP5qNrMZHkVnOjRMWwLCcr8ohBVb1OMjxLwGCvjTikrsBOiA6fNyCrm8V1rP93iVPpwaE+gO0SsWmPiXB+jikdf6SizrT5qKasx5j8ABbHpFTx+vFXp9EnYQmLx02h1QTTrl6eDqxLnGjporxl3NL3agEvXdT0WmEost648sQOYAeJS9Q7bfUVoMGnjo4AZdUMQku50McDcMWcBPvr0SzbTAFDfvJqwLzgxwATnCgnp4wDl6Aa+Ax283gghmj+vj7feE2KBBRMW3FzOpLOADl0Isb5587h/U4gGvkt5v60Z1VLG8BhYjbzRwyQZemwAd6cCR5/XFWLYZRIMpX39AR0tjaGGiGzLVyhse5C9RKC6ai42ppWPKiBagOvaYk8lO7DajerabOZP46Lby5wKjw1HCRx7p9sVMOWGzb/vA1hwiWc6jm3MvQDTogQkiqIhJV0nBQBTU+3okKCFDy9WwferkHjtxib7t3xIUQtHxnIwtx4mpg26/HfwVNVDb4oI9RHmx5WGelRVlrtiw43zboCLaxv46AZeB3IlTkwouebTr1y2NjSpHz68WNFjHvupy3q8TFn3Hos2IAk4Ju5dCo8B3wP7VPr/FGaKiG+T+v+TQqIrOqMTL1VdWV1DdmcbO8KXBz6esmYWYKPwDL5b5FA1a0hwapHiom0r/cKaoqr+27/XcrS5UwSMbQAAAABJRU5ErkJggg==)](https://deepwiki.com/suzuki-shunsuke/nllint)

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
