#!/bin/bash -eo pipefail

if [ -z "$1" ]; then
	echo "No selected port argument supplied"
	exit 1
fi

argIP="0.0.0.0"
argSleepInSeconds=1

i=$(( 60 / argSleepInSeconds ))
while [ $i -gt 0 ]; do
	i=$(( i - 1 ))
	(nc -vz $argIP $@ || nc -vzu $argIP $@) && break || sleep $argSleepInSeconds
done

if [ $i -eq 0 ]; then
	exit 1
fi
