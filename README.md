# rtx-helper

[![Build](https://github.com/ngyewch/rtx-helper/actions/workflows/build.yml/badge.svg)](https://github.com/ngyewch/rtx-helper/actions/workflows/build.yml)

Helper tool for [rtx](https://github.com/jdxcode/rtx).

## Installation

asdf-helper can be installed via [asdf](https://asdf-vm.com/). See https://github.com/ngyewch/asdf-helper-plugin

## Features

### Install (recursive)

```
rtx-helper install
```

### Latest version (recursive)

```
rtx-helper latest
```

Version constraints can be specified in `.tool-versions`. See https://github.com/Masterminds/semver for more details.

```
nodejs 16.17.1 # (constraint ^16)
java openjdk-11.0.2 # (constraint ^11)
golang 1.17.13 # (constraint >= 1.17, < 1.19)
```
