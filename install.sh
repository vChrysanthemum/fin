#!/usr/bin/env bash

if [ ! -d "$HOME/.in/lua" ];
then
    mkdir -p "$HOME/.in"
    cp -rf "$(pwd)/lua" "$HOME/.in/lua"
fi

if [ ! -d "$HOME/.in/project" ];
then
    mkdir -p "$HOME/.in"
    cp -rf "$(pwd)/project" "$HOME/.in/project"
fi
