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

## How to run

If you are unable to install Go on your system please get one of the compiled binaries that are available under [Releases](https://github.com/trubb/trubb_markerparser/releases/).

To run locally: `./trubb_markerparser totfs -s <input file>`

```bash
$ ./trubb_markerparser totfs -s samples/test.txt
2023/02/11 22:44:24 Running sweetToTun on file "samples/test.txt"
2023/02/11 22:44:24 Reading from file samples/test.txt
2023/02/11 22:44:24 Parsed the input to the following trp array:

[[],[["F4","01515","10485",[1515.12,10485.8]],["markertwo","03171","10742",[3171.69,10742.3]],["F6","00123","00098",[123.45,98.01]],["CT","03857","01855",[3857.04,1855.34]]]]

2023/02/11 22:44:24 No file with same name found, creating...
2023/02/11 22:44:24 Created output file parsed_markers_2023-02-11T22:44:24+01:00.txt
2023/02/11 22:44:24 Finished```
