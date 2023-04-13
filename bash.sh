#!/bin/bash

go run main.go input

if [[ $ALL_KEYS_PRESENT == "true"]]; then
    echo $(go run yo.go)
else
    echo "Exit from Bash"
    return
fi

go run main.go output