#!/bin/bash

netstat -vnatp 2&> /dev/null | grep 9090 | awk '{print $7}' | cut -d '/' -f1 | xargs kill -9
