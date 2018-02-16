set -e

export GOPATH=$(cd `dirname $0`; pwd)

if [ -n "$1" ]; then
  go install golang/$1
else
	go install golang/main
fi

