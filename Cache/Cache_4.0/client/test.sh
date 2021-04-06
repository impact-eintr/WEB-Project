#!/bin/bash
for i in {1..10}
do
    echo $i
    ./client S test$i 你好$i
done

for i in {1..10}
do
    ./client G test$i
done
