# benomad
is a bash script manager written on Go

The main idea was to make fast, simple and maintainable product, so its here!

## How it works?
Let me explain!
In ~/.benomad/ you have .ben files, like my_install_script.ben, and thats how it looks on your (my) computer
```
name: "my_install_script"
version: "rolling"
description: "installer for gmfi"
script: "/Users/egor/.benomad/scripts/install.sh"
```
But lets see how it looks on the internet, how developer of script [published it](https://raw.githubusercontent.com/Fynjirby/test-repo/refs/heads/main/my_install_script.ben)
```
name: "my_install_script"
version: "@"
description: "installer for gmfi"
# gmfi installer as script
script: "https://raw.githubusercontent.com/Fynjirby/gmfi/refs/heads/main/install.sh"
```
Oh! thats kinda different... why? Thats because when you install this ben, benomad's installer converts script path, removes comments, skips empty lines, changes version to "rolling" from "@", and in future will maybe do some more things.
So, now in ~/.benomad/ you have 1st file, where script path is on your PC and etc,  thats the file benomad works with.

In benomad you can install some bens, remove them, edit bens and scripts, run bens, etc...
Lets see full command list 👇

## Command list
- install \<ben\> - install some benomad metadata files (bens)
- remove \<ben\> - remove benomad 
- list - list all bens installed
- run \<ben\> - run a ben's script
- info \<ben\> - see information about any ben
- edit \<ben\> - edit ben file or script of ben

There will be more soon :) _(its all under development, check releases for new version)_

## Aliases to commands
- install - add
- remove - delete
- list - ls
- run - exec
- help - man

For sure will be more soon

Thanks so much for reading this all and using benomad!
