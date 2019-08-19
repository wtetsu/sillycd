# sillycd

[![Build Status](https://travis-ci.com/wtetsu/sillycd.svg?branch=master)](https://travis-ci.com/wtetsu/sillycd)

Reduce your typing dramatically when doing cd.

# Usage

## Bash

```sh
$ c /l/as/as
/Library/Application Support/App Store
```

## Windows

```bat
c:\>c /pf/mss/100
c:\Program Files\Microsoft SQL Server\100

c:\>c /pf/mss
c:\Program Files\Microsoft SQL Server

c:\>c /pf/mss/v
C:\Program Files\Microsoft SQL Server Compact Edition\v4.0
```

# Install

## Install a binary file

```
go get -u github.com/wtetsu/sillycd
```

## For bash

Define a shell function

```sh
#! /bin/sh

function c() {
  d=`sillycd $1`
  if [ $? -eq 0 ]; then
    echo $d
    cd "$d"
  fi
}
```

## For Windows

Put a bat file like c.bat.

```bat
@set __sillycd_target=
@for /f "delims=" %%i in ('sillycd %1') do @set __sillycd_target=%%i
@if not "%__sillycd_target%" == "" (
  @echo %__sillycd_target%
  @cd /d "%__sillycd_target%"
)
@set __sillycd_target=
```
