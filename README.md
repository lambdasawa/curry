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
