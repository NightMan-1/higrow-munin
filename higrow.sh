#!/bin/sh

case $1 in
   config)
		curl http://127.0.0.1:8083/config
        exit 0;;
esac


curl http://127.0.0.1:8083/get
