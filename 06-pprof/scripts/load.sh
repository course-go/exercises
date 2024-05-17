#!/bin/sh

while true; do
        count=$(( ( RANDOM % 128 )  + 1 ))
        curl -X POST http://localhost:${PORT}/data?count=${count} &
        sleep 0.1
done

