APP?=blackholemon
#RELEASE?=$(shell type c:\aworks\blackholemon\VERSION.txt)
RELEASE?=$(shell python version.py get)
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell powershell get-date -format "{yyyy-mm-dd_HH:mm:ss}")
PROJECT?=github.com/Jarover/BlackHoleMon

clean:
	del ${APP}
	del ${APP}.exe

build:	clean
	python version.py inc-patch
	
	go build \
                -ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
                -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
                -o ${APP}

run:	build
	./${APP} -f dev.json

test:
	go test -v -race ./...