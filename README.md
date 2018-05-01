## Codechain — code trust through hash chains

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/frankbraun/codechain) [![Build Status](https://img.shields.io/travis/frankbraun/codechain.svg?style=flat-square)](https://travis-ci.org/frankbraun/codechain) [![Go Report Card](https://goreportcard.com/badge/github.com/frankbraun/codechain?style=flat-square)](https://goreportcard.com/report/github.com/frankbraun/codechain)

Codechain establishes code trust via multi-party reviews recorded in
unmodifiable hash chains.

Codechain allows to only publish code that has been reviewed by a
preconfigured set of reviewers. The signing keys can be rotated and the
reviewer set flexibly changed.

Every published code state is uniquely identified by a deterministic
source tree hash stored in the hash chain, signed by a single
responsible developer.

Please note that Codechain uses files to store the hash chain, not a
distributed "blockchain".

### Installation

```
go get -u -v github.com/frankbraun/codechain
```

### Features

- [x] Minimal code base, Go only, cross-platform.

Currently Codechain depends on the `git` binary (for `git diff` and
`git apply`).

### Out of scope

- Source code management. Git and other VCS systems are good for that, Codechain
  can be used alongside them and solves a different problem.
- Single source of truth (SSOT). Codechain requires a SSOT to distribute the
  current hash chain head, but that's outside of the scope for now. DNS (plus
  DNSCrypt) or DNSSEC could be used. Gossiping of the current head would also
  work.
- Code distribution.

### Acknowledgments

Codechain has been heavily influenced by discussions with
[Jonathan Logan](https://github.com/JonathanLogan) of
[Cryptohippie](https://secure.cryptohippie.com/).
