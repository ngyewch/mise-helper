# mise-helper

[![Build](https://github.com/ngyewch/mise-helper/actions/workflows/build.yml/badge.svg)](https://github.com/ngyewch/mise-helper/actions/workflows/build.yml)

Helper tool for [mise](https://github.com/jdx/mise).

## Installation

mise-helper can be installed via [mise](https://github.com/jdx/mise). See https://github.com/ngyewch/mise-helper-plugin

## Features

### Install (recursive)

```
mise-helper install
```

### Latest version (recursive)

```
mise-helper latest
```

Version constraints can be specified in `.mise-helper.toml`. See https://github.com/Masterminds/semver for more details.

```
[constraints]
nodejs = '^16'
java = '^11'
golang = '>= 1.17, < 1.19'
```
