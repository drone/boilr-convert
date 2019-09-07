# boilr-convert

This is a [boilr template](https://github.com/tmrts/boilr) for creating a conversion extension. Use a conversion extension to convert or modify the Drone configuration file. For example, convert from Jsonnet to Yaml. Get started by cloning the project and installing the template:

```console
$ cd boilr-convert
$ boilr template save . drone-convert -f
```

create a project in directory my-converter:

```console
$ boilr template use drone-convert my-converter
```

enter the docker registry name for this project:

```text
[?] Please choose a value for "DockerRepository" [default: "foo/bar"]:
```

enter the go module name:

```text
[?] Please choose a value for "GoModule" [default: "github.com/foo/bar":
```
