go build .

mkdir -p .test/abc-def-ghi
mkdir -p .test/foo-bar-abc


CURRNET_DIR=`./sillycd .`

function test() {
  param=$1
  expected=$2
  actual=`./sillycd $1`
  if [ "$actual" == "$expected" ]; then
    echo "OK: $param $actual"
  else
    echo "ERROR:$param $actual $expected"
  fi
}

test . $CURRNET_DIR
test .. `dirname $CURRNET_DIR`
test .test/adg `./sillycd .test/abc-def-ghi`
test .test/f `./sillycd .test/foo-bar-abc`
test .test/fb `./sillycd .test/foo-bar-abc`
test .test/fba `./sillycd .test/foo-bar-abc`
test .test/fbab `./sillycd .test/foo-bar-abc`
