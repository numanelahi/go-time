#!/bin/bash
if [ $1 ]; then
    /app/go-time $1 $2
else
    /app/go-time
fi