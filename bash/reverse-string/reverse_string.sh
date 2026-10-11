#!/usr/bin/env bash


input=${1-}
reversed=''
for ((index = ${#input} - 1; index >= 0; index--)); do
    reversed+=${input:index:1}
done

printf '%s\n' "$reversed"
