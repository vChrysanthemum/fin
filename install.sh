#!/usr/bin/env bash

if [ ! -d "$HOME/.in/lua" ];
then
    mkdir -p "$HOME/.in"
    ln -s "$(pwd)/lua" "$HOME/.in/lua"
fi

if [ ! -d "$HOME/.in/project" ];
then
    mkdir -p "$HOME/.in"
    ln -s "$(pwd)/project" "$HOME/.in/project"
fi
