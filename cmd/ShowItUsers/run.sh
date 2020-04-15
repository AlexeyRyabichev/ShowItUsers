#!/bin/bash
go build
nohup ./ShowItUsers >>out.log 2>&1 &
