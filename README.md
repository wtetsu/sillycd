[![Build Status](https://travis-ci.com/wtetsu/sillycd.svg?branch=master)](https://travis-ci.com/wtetsu/sillycd)

# SillyCd

![iconfinder_kdisknav_1311](https://user-images.githubusercontent.com/515948/71973704-02de0300-3253-11ea-88da-8e8129647a98.png)

The fastest way to type cd command

# Usage

SillyCd helps you reduce keystrokes dramatically for changing current directory.

## Bash

```sh
$ c /u/l/b
/usr/local/bin

$ c /l/as/as
/Library/Application Support/App Store

$ c /a/v/c/r/a/e/
/Applications/Visual Studio Code.app/Contents/Resources/app/extensions

```

## Windows

```
c:\>c /pf/mss/100
c:\Program Files\Microsoft SQL Server\100

c:\>c /pf/mss
c:\Program Files\Microsoft SQL Server

c:\>c /pf/mss/v
C:\Program Files\Microsoft SQL Server Compact Edition\v4.0
```

## Installation

### Installation: OSX

1. Get a binary file
1. Define a function

### Get binary

```sh
brew tap wtetsu/sillycd
brew install sillycd
```

or [download](https://github.com/wtetsu/sillycd/releases/download/v1.0.0/sillycd-v1.0.0-darwin-amd64.zip) and add it to a PATH directory.

### Define function

sillycd just writes a matched directory to stdout. If you use sillycd as alternative cd, you have to define a shell function.

```sh
# For bash:
function c() {
  d=`sillycd $1`
  if [ $? -eq 0 ]; then
    echo $d
    cd "$d"
  else
    echo "$1: No such file or directory" >&2
  fi
}
```

### Give it a try!

```
$ c /u/l/b
```

## Installation: Windows

### Get files

[Download it!](https://github.com/wtetsu/sillycd/releases/download/v1.0.0/sillycd-v1.0.0-windows-amd64.zip)

And add these files into a PATH directory.

- sillycd.exe
- c.bat

### Give it a try!

```
c:\>c /pf
```

## Rules

- You have to type the first character of each directory
- You can omit middle characters
- If matched multiple directories, sillycd picks the most "probable" directory

## Third-party data

### Images

Folder, lightning, power icon(LGPL)

https://www.iconfinder.com/icons/1311/folder_lightning_power_icon
