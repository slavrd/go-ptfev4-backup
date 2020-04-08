# tfev4backup

A basic CLI utility to backup / restore a TFEv4 data to a file.

## Building the command

To build the utility need to have Go version 1.14 or greater installed.

- clone the repository

```bash
go get -v -d  github.com/slavrd/go-tfev4-backup/tfev4backup
```

This will download the source code of the command to `$GOPATH/src/slavrd/go-tfev4-backup/tfev4backup`

If `$GOPATH` is not set the default value used is `$HOME/go`

- build the binary

```bash
go install github.com/slavrd/go-tfev4-backup/tfev4backup
```

The compiled binary will be `$GOPATH/bin/tfev4backup`

## Usage

To use basic command usage is

```bash
./tfev4backup [args] <backup | restore>
```

The arguments can be provided via CLI flags or environment variables according to the table below.

| CLI flag | Environment Variable | Mandatory | Description |
| -------- | -------------------- | :-------: | ----------- |
| -host | TFE_HOSTNAME | yes | Hostname of the TFE instance. |
| -token | TFE_BACKUP_TOKEN | yes | TFE backup token. |
| -pass | TFE_BACKUP_PASSWORD | yes | Encryption password for the backed up data. |
| -file | TFE_BACKUP_FILE | yes | File to write to or restore from. |

**Note:** CLI flags take precedence over environment variables.

### Examples

- backup

```bash
./tfev4backup \
    -host my.tfev4.com \
    -token agk93Jcd*@%dsa13 \
    -pass myPassw0rd \
    -file backup.blob \
    backup
```

- restore
  
```bash
export TFE_HOSTNAME=my.tfev4.com
export TFE_BACKUP_TOKEN=vgr^34*!dgFtkyaA
export TFE_BACKUP_PASSWORD=myPassw0rd
export TFE_BACKUP_FILE=backup.blob

./tfev4backup restore
```
