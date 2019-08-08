# smartcd

[![Build Status](https://travis-ci.com/wtetsu/stupidcd.svg?branch=master)](https://travis-ci.com/wtetsu/stupidcd)

## Install

```
go get -u github.com/wtetsu/stupidcd
```

## Define a shell function

### bash

```
#! /bin/sh

function c() {
  d=`stupidcd $1`
  if [ $? -eq 0 ]; then
    echo $d
    cd $d
  fi
}
```
