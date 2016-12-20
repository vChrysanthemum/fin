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

if [ ! -d "$(pwd)/src/main/" ];
then 
    ln -s "$(pwd)/main" "$(pwd)/src/main"
fi

if [ ! -d "$(pwd)/src/ui/" ];
then 
    mkdir -p "$(pwd)/src/ui"
    ln -s "$(pwd)/test/ui.go" "$(pwd)/src/ui/ui.go"
fi

if [ ! -d "$(pwd)/src/fin/" ];
then 
    mkdir src/fin
fi

for folder in `ls ./lib/`
do 
    if [ ! -d "$(pwd)/src/fin/$folder" ];
    then
        ln -s "$(pwd)/lib/$folder" "$(pwd)/src/fin/$folder"
    fi
done


for folder in `ls ./3rdlib/`
do 
    if [ ! -d "$(pwd)/src/$folder" ];
    then
        ln -s "$(pwd)/3rdlib/$folder" "$(pwd)/src/$folder"
    fi
done

if [ ! -e "$(pwd)/lua" ];
then
    mkdir -p "$(pwd)/lua"
fi


for folder in `ls ./lib/`
do 
    if [ -e "$(pwd)/lib/$folder/lua" ];
    then
        if [ ! -e "$(pwd)/lua/$folder" ];
        then
            ln -s "$(pwd)/lib/$folder/lua" "$(pwd)/lua/$folder"
        fi
    fi
done
