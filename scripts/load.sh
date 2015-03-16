#!/bin/bash

HOST=`hostname`
DATE=`date +%s`
LOAD=`cat /proc/loadavg | awk '{print $1}'`

cat << EOF
{
    "time": $DATE,
    "state": "d",
    "service": "system",
    "host": "$HOST",
    "metric_f": $LOAD
}
EOF
