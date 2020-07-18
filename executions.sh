#!/bin/bash

PROG=$1
N=$2
TIMEOUT=$3

seq $N | xargs -I{} -P1 -n1 sh -c "echo Execution {}; $PROG | ts '%.T'; echo ; sleep $TIMEOUT"
