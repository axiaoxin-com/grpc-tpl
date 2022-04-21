#! /usr/bin/env bash

main() {
    gopath=`go env GOPATH`
    if [ $? = 127 ]; then
        echo "GOPATH not exists"
        exit -1
    fi
    echo -e "New project will be create in ${gopath}/src/"
    echo -ne "Enter your new project full name (eg. github.com/my_username/my_projname): "
    read projname

    # get template project
    echo -e "Downloading the template..."
    if !(curl https://codeload.github.com/axiaoxin-com/grpc-tpl/zip/main -o /tmp/grpc-tpl.zip && unzip /tmp/grpc-tpl.zip -d /tmp)
    then
        echo "Downloading failed."
        exit -2
    fi

    echo -e "Generating the project..."
    mv /tmp/grpc-tpl-main ${gopath}/src/${projname} && cd ${gopath}/src/${projname}

    if [ `uname` = 'Darwin' ]; then
        sed -i '' -e "s|github.com/axiaoxin-com/grpc-tpl|${projname}|g" `grep "grpc-tpl" --include "*.proto" --include ".travis.yml" --include "*.go" --include "go.*" -rl .`
    else
        sed -i "s|github.com/axiaoxin-com/grpc-tpl|${projname}|g" `grep "grpc-tpl" --include "*.proto" --include ".travis.yml" --include "*.go" --include "go.*" -rl .`
    fi

    if [ $? -ne 0 ]
    then
        echo -e "set project name failed."
        exit -3
    fi

    echo -e "Create project ${projname} in ${gopath}/src succeed."

    # init a git repo
    echo -ne "Do you want to init a git repo[N/y]: "
    read initgit
    if [ "${initgit}" == "y" ] || [ "${initgit}" == "Y" ]; then
        cd ${gopath}/src/${projname} && git init && git add . && git commit -m "init project with grpc-tpl"
        cp ${gopath}/src/${projname}/misc/scripts/pre-push.githook ${gopath}/src/${projname}/.git/hooks/pre-push
        chmod +x ${gopath}/src/${projname}/.git/hooks/pre-push
    fi
}
main
