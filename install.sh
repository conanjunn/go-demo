set -e

export GOPATH=$(cd `dirname $0`; pwd)

if [ -n "$1" ]; then
  go install $1
else
	go install main
fi

