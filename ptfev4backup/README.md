# ptfev4backup

A basic CLI utility to backup / restore a PTFEv4 data to a file.

## Building the command

To build the utility need to have Go version 1.14 or greater installed.

- clone the repository

```bash
go get -v -d  github.com/slavrd/go-ptfev4-backup/ptfev4backup
```

This will download the source code of the command to `$GOPATH/src/slavrd/go-ptfev4-backup/ptfev4backup`

If `$GOPATH` is not set the default value used is `$HOME/go`

- build the binary

```bash
go install github.com/slavrd/go-ptfev4-backup/ptfev4backup
```

The compiled binary will be `$GOPATH/bin/ptfev4backup`

## Usage

To use basic command usage is

```bash
./ptfev4backup [args] <backup | restore>
```

The arguments can be provided via CLI flags or environment variables according to the table below.

| CLI flag | Environment Variable | Mandatory | Description |
| -------- | -------------------- | :-------: | ----------- |
| -host | PTFE_HOSTNAME | yes | Hostname of the PTFE instance. |
| -token | PTFE_BACKUP_TOKEN | yes | PTFE backup token. |
| -pass | PTFE_BACKUP_PASSWORD | yes | Encryption password for the backed up data. |
| -file | PTFE_BACKUP_FILE | yes | File to write to or restore from. |

**Note:** CLI flags take precedence over environment variables.

### Examples

- backup

```bash
./ptfev4backup \
    -host my.ptfev4.com \
    -token agk93Jcd*@%dsa13 \
    -pass myPassw0rd \
    -file backup.blob \
    backup
```

- restore
  
```bash
export PTFE_HOSTNAME=my.ptfev4.com
export PTFE_BACKUP_TOKEN=vgr^34*!dgFtkyaA
export PTFE_BACKUP_PASSWORD=myPassw0rd
export PTFE_BACKUP_FILE=backup.blob

./ptfev4backup restore
```
