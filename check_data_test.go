package main

var data = `github.com/allowed/packageApache,version,Apache-2.0
github.com/allowed/packageMIT,version,MIT
github.com/allowed/packageBSD2,version,BSD-2-Clause
github.com/allowed/packageBSD3,version,BSD-3-Clause
github.com/allowed/packageISC,version,ISC
github.com/allowed/packageRND,version,RND
`

var configAllAllowed = `allowed:
  type:
    - Apache-2.0
    - MIT
    - BSD-2-Clause
    - BSD-3-Clause
    - ISC
    - RND
blocked:
  type:
    - MPL-2.0 # this is not compatible with other allowed licenses
`

var configPackageRNDAllowed = `allowed:
  type:
    - Apache-2.0
    - MIT
    - BSD-2-Clause
    - BSD-3-Clause
    - ISC
  packages:
    - package: github.com/allowed/packageRND
      license: RND
blocked:
  type:
    - MPL-2.0 # this is not compatible with other allowed licenses
`

var configRNDNotAllowed = `allowed:
  type:
    - Apache-2.0
    - MIT
    - BSD-2-Clause
    - BSD-3-Clause
    - ISC
blocked:
  type:
    - MPL-2.0 # this is not compatible with other allowed licenses
`

var configRNDBlocked = `allowed:
  type:
    - Apache-2.0
    - MIT
    - BSD-2-Clause
    - BSD-3-Clause
    - ISC
blocked:
  type:
    - RND
    - MPL-2.0 # this is not compatible with other allowed licenses
`

var configRNDIsOKPackageRNDBlocked = `allowed:
  type:
    - Apache-2.0
    - MIT
    - BSD-2-Clause
    - BSD-3-Clause
    - ISC
    - RND
blocked:
  packages:
    - github.com/allowed/packageRND
  type:
    - MPL-2.0 # this is not compatible with other allowed licenses
`
