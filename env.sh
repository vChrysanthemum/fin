#!/usr/bin/env bash

# adjust GOPATH
case ":$GOPATH:" in
    *":$(pwd):"*) :;;
    *) GOPATH=$(pwd):$GOPATH;;
esac
export GOPATH


# adjust PATH
while IFS=':' read -ra ADDR; do
    for i in "${ADDR[@]}"; do
        case ":$PATH:" in
            *":$i/bin:"*) :;;
            *) PATH=$i/bin:$PATH
        esac
    done
done <<< "$GOPATH"
export PATH

if [ ! -d "$(pwd)/bin" ];
then
    mkdir -p "$(pwd)/bin"
fi

if [ ! -d "$(pwd)/src" ];
then
    mkdir -p "$(pwd)/src"
fi

if [ ! -d "$(pwd)/src/in/" ];
then 
    mkdir src/in
fi

for folder in `ls ./lib/`
do 
    if [ ! -d "$(pwd)/src/in/$folder" ];
    then
        ln -s "$(pwd)/lib/$folder" "$(pwd)/src/in/$folder"
    fi
done


for folder in `ls ./3rdlib/`
do 
    if [ ! -d "$(pwd)/src/$folder" ];
    then
        ln -s "$(pwd)/3rdlib/$folder" "$(pwd)/src/$folder"
    fi
done

for folder in `ls ./lib/`
do 
    if [ -d "$(pwd)/lib/$folder/lua" ];
    then
        if [ ! -d "$(pwd)/lua/$folder" ];
        then
            ln -s "$(pwd)/lib/$folder/lua" "$(pwd)/lua/$folder"
        fi
    fi
done
