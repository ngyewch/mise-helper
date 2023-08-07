# rtx-helper

[![Build](https://github.com/ngyewch/rtx-helper/actions/workflows/build.yml/badge.svg)](https://github.com/ngyewch/rtx-helper/actions/workflows/build.yml)

Helper tool for [rtx](https://github.com/jdxcode/rtx).

## Installation

rtx-helper can be installed via [rtx](https://github.com/jdxcode/rtx). See https://github.com/ngyewch/rtx-helper-plugin

## Features

### Install (recursive)

```
rtx-helper install
```

### Latest version (recursive)

```
rtx-helper latest
```

Version constraints can be specified in `.rtx-helper.toml`. See https://github.com/Masterminds/semver for more details.

```
[constraints]
nodejs = '^16'
java = '^11'
golang = '>= 1.17, < 1.19'
```
