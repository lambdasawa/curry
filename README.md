# curry

## Install

```sh
$ go install github.com/lambdasawa/curry@latest
```

## Usage

Basic usage.

```
$ curry echo foo bar
> hoge
foo bar hoge
> fuga
foo bar fuga
>
```

`curry` replace placeholder (`{}`) like xargs.

```
$ curry echo foo {} bar
> hoge
foo hoge bar
> fuga
foo fuga bar
>
```

Shortcut key available by [liner](https://github.com/peterh/liner#line-editing).

## Examples

```
$ curry docker
> ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
> run -d nginx
bcb800a4f7791beee3dda968063e2fc242097131dfe8b1cb9ab5d7e20376a69d
> ps
CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS        PORTS     NAMES
bcb800a4f779   nginx     "/docker-entrypoint.…"   2 seconds ago   Up 1 second   80/tcp    romantic_elbakyan
>
```
