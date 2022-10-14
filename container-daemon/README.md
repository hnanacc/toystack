### Container Daemon (like docker)

* Actual docker api:
```sh
$ docker <options> run <img> <cmd> <args>
```

* Our api:
```sh
$ ./main run <img>
```

> For sanity, we define `<img>` as path to an executable.
