#!/bin/bash
for (( i=2; i<=700; i++ )); do
            echo "Create vscode $i"
		docker rm -f  "user-$i"
	echo "removed $i"
        done
