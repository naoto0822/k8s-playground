#!/bin/bash

for i in {0..200}
do
    curl http://localhost:10001/v1/soba;
    sleep 5;
done
