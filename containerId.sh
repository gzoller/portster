#!/bin/sh

cat /proc/self/cgroup | grep "docker" | head -1 | sed s/.*\\///g