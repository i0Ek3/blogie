#!/bin/bash

install_mysql_then_import() {
    platform=$(uname -s)

    if [ $platform == "Darwin" ]
    then
        brew update ; brew upgrade
        brew install mysql
        brew services start mysql
    elif [ $platform == 'Linux' ]
    then
        sudo apt update ; sudo apt install -y mysql
        sudo service mysql start
    else
        echo "Unsupported platform!"
    fi
    mysql -uroot < ./scripts/sql/blog.sql
}

main() {
    install_mysql_then_import
    echo -n "All done!"
}

main
