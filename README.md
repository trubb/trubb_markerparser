# trubb_markerparser

Parses marker arrays generated from [Sweet markers system](https://steamcommunity.com/sharedfiles/filedetails/?id=324952672) and generates a trp array for use with [Tuntematon Firesupport](https://github.com/tuntematonjr/Tun-Firesupport).

Executable files can be found under releases.

To build from source, simply clone the repo, run `go get`, followed by `go build`.

```yaml
NAME:
   markerparser - Translate between marker and artillery target/bookmark arrays

USAGE:
   markerparser [global options] command [command options] [arguments...]

COMMANDS:
   toTuntematonFireSupport, totfs  from sweet marker array to tuntematon firesupport trps
   OPTIONS:
      --sourceFile FILE, -s FILE  load array from FILE

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```
