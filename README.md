# TFEv4 Backup

Go packages to backup / restore a TFEv4 installation.

## Contents

The repo the following packages.

- a package with helper functions `helpers`. It contains functions to backup or restore PTFEv4 instance data from/to provided streams - [documentation](https://godoc.org/github.com/slavrd/go-ptfev4-backup/helpers).
- a command line tool `tfev4backup` to backup or restore a PTFEv4 instance data to a file. Check the [readme](./ptfev4backup/README.md) on details on how to use it.

## TODO

- [ ] - refactor helpers so that they can be tested
- [ ] - setup a CI/CD 
