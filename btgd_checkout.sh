#!/bin/bash

git checkout $(cat go.sum | grep github.com/btgsuite/btgd | tail -n1 | awk -F " " '{ print $2 }' | awk -F "-" '{ print $3 }' | awk -F "/" '{ print $1 }')
