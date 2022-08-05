#!/bin/bash

setup_goenv() {
  	export GO111MODULE=on; export GOPROXY=https://goproxy.cn
  	go mod tidy; go mod vendor
}

update_for_linux() {
    sudo apt update; sudo apt upgrade
}

update_for_mac() {
    brew update; brew upgrade
}

install_mysql_for_linux() {
    update_for_linux
    sudo apt-get install -y mysql-server
    systemctl start mysql
    sudo mysql -uroot < ./scripts/sql/blog.sql
}

install_docker_for_linux() {
    update_for_linux
    sudo apt install -y docker docker-compose
    docker-compose up -d
}

install_mysql_for_mac() {
    update_for_mac
    brew install mysql
    brew services start mysql
    mysql -uroot < ./scripts/sql/blog.sql
}

install_docker_for_mac() {
    update_for_mac
    brew install docker docker-compose
    docker-compose up -d
}

install_mysql_and_docker() {
    platform=$(uname -s)

    if [ $platform == "Darwin" ]
    then
        install_mysql_for_mac
        install_docker_for_mac
    elif [ $platform == 'Linux' ]
    then
        install_mysql_for_linux
        install_docker_for_linux
    else
        echo "Unsupported platform!"
    fi
}

main() {
    setup_goenv
    install_mysql_and_docker
    echo -n "All done!"
}

main
