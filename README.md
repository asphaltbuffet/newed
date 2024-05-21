# newed - a birthday suit for your project

<div align="center">

[![GitHub release (with filter)](https://img.shields.io/github/v/release/asphaltbuffet/newed)](https://github.com/asphaltbuffet/newed/releases)
[![go.mod](https://img.shields.io/github/go-mod/go-version/asphaltbuffet/newed)](go.mod)
[![GitHub License](https://img.shields.io/github/license/asphaltbuffet/newed)](LICENSE)
[![Common Changelog](https://common-changelog.org/badge.svg)](https://common-changelog.org)
[![wakatime](https://wakatime.com/badge/user/09307b0e-8348-4b4e-9b67-0026db3fe1f5/project/f64b1b24-0d96-4b60-aa34-231273d60dc1.svg)](https://wakatime.com/badge/user/09307b0e-8348-4b4e-9b67-0026db3fe1f5/project/f64b1b24-0d96-4b60-aa34-231273d60dc1)

[![CodeQL](https://github.com/asphaltbuffet/newed/workflows/CodeQL/badge.svg)](https://app.codecov.io/gh/asphaltbuffet/newed)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=asphaltbuffet_newed&metric=coverage)](https://sonarcloud.io/summary/new_code?id=asphaltbuffet_newed)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=asphaltbuffet_newed&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=asphaltbuffet_newed)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=asphaltbuffet_newed&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=asphaltbuffet_newed)

</div>

`newed`` creates projects from templates.

## Manifesto

1. MUST create a project in the specified directory from a template
2. The user MAY add/remove/modify project templates
3. MUST run and pass linting/testing as a template
4. MUST NOT require manual edits after generation to run and pass linting/testing
5. MUST support any file type or internal directory structure

## Install

```sh
go install github.com/asphaltbuffet/newed@latest
```

## Configuration

`.newed.toml`

application specific configuration

`bootstrap.toml`

Optional in each template directory. It defines post-generation key-value replacements.

## Build

Build scripts use [Task](https://taskfile.dev).

Install the necessary dependencies with `task install`. This should usually only be necessary to do once.

`task snapshot` will create a local build for your OS in `./dist/newed-<OS name>/`.
