### Container Daemon (like docker)

* Actual docker api:
```sh
$ docker <options> run <img> <cmd> <args>
```

* Our api:
```sh
$ ./container-daemon run <img>
```

> For sanity, we define `<img>` as a json file as shown in the `stub` folder.

#### Build and run instructions
* From project root
```sh
$ go build
$ sudo ./container-daemon run stub/image.json
```

#### Structure of `image.json`
```json
{
    "Cmd": "/path/to/executble",
    "Hostname": "hostname of container(string)",
    "MaxMem": "max memory usage(in bytes, as string)",
    "MaxPids": "max no. of process allowed(in integer, as string)
}
```
