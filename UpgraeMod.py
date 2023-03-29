#!/usr/bin/env python3
# -*- encoding: utf-8 -*-
from os import system

system("go clean")
for i in ["gns", "gcs", "gbm", "gf"]:
    for s in range(3):
        cmd = f"go get -u gitee.com/liumou_site/{i}"
        print(cmd)
        system(cmd)
system("go mod tidy")
