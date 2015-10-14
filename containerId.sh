#!/bin/sh

cat /proc/self/cgroup | grep "docker" | sed s/\\//\\n/g | tail -1
