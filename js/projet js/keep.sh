#!/bin/bash
stop0= $(kill -STOP $0)
echo $stop0
data = `jobs`
echo $data
bg 
echo 'end '
