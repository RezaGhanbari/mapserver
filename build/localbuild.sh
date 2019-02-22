#!/usr/bin/env bash
export CC=clang &&
go get github.com/cnjack/throttle &&
go get github.com/gin-gonic/gin &&
go get github.com/gin-contrib/cache &&
go get github.com/satori/go.uuid &&
go get github.com/gomodule/redigo/redis &&
go get github.com/olivere/elastic &&
go get github.com/JamesMilnerUK/pip-go &&
go get github.com/corpix/uarand