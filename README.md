# curry

## Install

```sh
$ go install github.com/lambdasawa/curry@latest
```

```sh
$ brew tap lambdasawa/tap
$ brew install lambdasawa/tap/curry
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

`curry` replace placeholder (`{}`).

```
$ curry echo foo {} bar
> hoge
foo hoge bar
> fuga
foo fuga bar
>
```

`curry` replace and expand placeholder (`{...}`).

```
$ curry git {}
> status -s
git: 'status -s' is not a git command. See 'git --help'.
exit status 1
>

$ curry git {...}
> status -s
 M README.md
 M main.go
>

$ curry mysql -uroot -proot -h127.0.0.1 -P3306 -e {...}
> show databases
ERROR 1049 (42000): Unknown database 'databases'
exit status 1
>

$ curry mysql -uroot -proot -h127.0.0.1 -P3306 -e {}
> show databases
Database
information_schema
mysql
performance_schema
sys
>
````

`curry` can read stdin.

```
$ curl -s https://httpbin.org/headers | curry jq
> .
{
  "headers": {
    "Accept": "*/*",
    "Host": "httpbin.org",
    "User-Agent": "curl/7.64.1"
  }
}
> .headers.Host
"httpbin.org"
>
```

Shortcut key available by [go-prompt](https://github.com/c-bata/go-prompt#keyboard-shortcuts).

## Examples

```
$ curry docker {...}
> ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
> run -d nginx
bcb800a4f7791beee3dda968063e2fc242097131dfe8b1cb9ab5d7e20376a69d
> ps
CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS        PORTS     NAMES
bcb800a4f779   nginx     "/docker-entrypoint.…"   2 seconds ago   Up 1 second   80/tcp    romantic_elbakyan
>
```
