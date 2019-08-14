# smartcd

[![Build Status](https://travis-ci.com/wtetsu/sillycd.svg?branch=master)](https://travis-ci.com/wtetsu/sillycd)

## Install

```
go get -u github.com/wtetsu/sillycd
```

## Define a shell function

### bash

```
#! /bin/sh

function c() {
  d=`sillycd $1`
  if [ $? -eq 0 ]; then
    echo $d
    cd $d
  fi
}
```
