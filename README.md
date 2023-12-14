# ![HAProxy](assets/images/haproxy-weblogo-210x49.png "HAProxy")

## Go Logger

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![Github tag](https://badgen.net/github/tag/oktalz/licence-checker)](https://github.com/oktalz/licence-checker/tags/)
[![Go Report Card](https://goreportcard.com/badge/github.com/oktalz/licence-checker)](https://goreportcard.com/report/github.com/oktalz/licence-checker)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/oktalz/licence-checker)
[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/oktalz/licence-checker.svg)](https://github.com/oktalz/licence-checker)
[![Contributors](https://img.shields.io/github/contributors/oktalz/licence-checker?color=purple)](https://github.com/haproxy/haproxy/blob/master/CONTRIBUTING)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

### Description

Go License checker is an tool to check is packages have correct licenses

### Prerequisite

this tool depends on [github.com/google/go-licenses](https://github.com/google/go-licenses) being installed

```bash
go install github.com/google/go-licenses@latest
```

### Instalation

```bash
go install github.com/haproxytech/go-licence-checker@latest
```

### Running

application expects `.license-checker.yml` file in same folder as `go.mod` file

location can also be specified with ENV variable `LICENSE_CHECKER`

### Example

```yaml
allowed:
    type:
        - Apache-2.0
        - MIT
        - BSD-2-Clause
        - BSD-3-Clause
        - ISC
blocked:
    type:
        - MPL-2.0 # this is not compatible with other allowed licenses
```

### Full Example

```yaml
allowed:
    type:
        - Apache-2.0
        - MIT
        - BSD-2-Clause
        - BSD-3-Clause
        - ISC
    packages:
        - package: github.com/group/package # with custom license
          licence: RND # type is needed since this can change over time
blocked:
    type: # this is not needed. all not listed in allowed are blocked
        - MPL-2.0 # this is not compatible with other allowed licenses
    packages:
        - github.com/group/package/ # with right license but not allowed

```


### Contributing

Thanks for your interest in the project and your willing to contribute:

- this project uses [taskfile](https://taskfile.dev/)
- Pull requests are welcome!
- For commit messages and general style please follow the haproxy project's [CONTRIBUTING guide](https://github.com/haproxy/haproxy/blob/master/CONTRIBUTING) and use that where applicable.

### Discussion

A Github issue is the right place to discuss feature requests, bug reports or any other subject that needs tracking.

To ask questions, get some help or even have a little chat, you can join our #ingress-controller channel in [HAProxy Community Slack](https://slack.haproxy.org).

## License

[Apache License 2.0](LICENSE)
