# benomad
is a bash script manager written on Go

The main idea was to make fast, simple and maintainable product, so its here!

## Install
#### Fastest way 
Just run this command in terminal and it will install everything itself
```sh
curl -L https://sh.fynjirby.dev/benomad | sh
```
or if you prefer GoLang package manager use
```sh
go install github.com/fynjirby/benomad@latest
```
#### Manual way
Go to [releases](https://github.com/Fynjirby/benomad/releases/) and download latest binary for your OS, then move it to `/usr/local/bin/` and enjoy with simple `benomad` in terminal!

## Building
- Install [Go](https://go.dev/) and make sure it's working with `go version`
- Clone repo
- Run `go build` in repo directory, then move it to `/usr/local/bin/`



## Command list
 - install \<url\> - install some script by url
 - remove \<script\> - remove script
 - list - list all scripts installed
 - run \<script\> - run a script
 - edit \<script\> - edit script

There will be more soon :) _(its all under development, check releases for new version)_

## Aliases to commands
- install - add
- remove - delete
- list - ls
- run - exec
- help - man

For sure will be more soon

Thanks so much for reading this all and using benomad!
