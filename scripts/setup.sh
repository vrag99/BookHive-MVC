#!/bin/bash

prettyPrint () {
    echo "-------------------------------------------------------"
    echo "$1"
    echo "-------------------------------------------------------"
}

setupConfig() {
    prettyPrint "Making config.yaml, enter database credentials"
    read -p "Username: " dbUser && echo "dbUser: $dbUser" > config.yaml
    read -s -p "Password: " password && echo "password: $password" >> config.yaml && echo ""
    read -p "Host: " host && echo "host: $host" >> config.yaml
    read -p "Port: " port && echo "port: $port" >> config.yaml
    read -p "Database Name: " dbName && echo "dbName: $dbName" >> config.yaml

    read -p "Generate a random JWT signing key (y/n)?" genRandom

    if [[ "$genRandom" == "y" && `which openssl` ]]; then
        accessTokenSecret=$(openssl rand -hex 16)
        echo "accessTokenSecret: $accessTokenSecret" >> config.yaml

    else
        echo "Couldn't generate random JWT signing key."
        read -s -p "Enter JWT signing key: " && echo "accessTokenSecret: $accessTokenSecret" >> config.yaml
    fi
}

migrateDB() {
    prettyPrint "Applying migrations..."
    migrate -path database/migration/ -database "mysql://$dbUser:$password@tcp($host:$port)/$dbName" -verbose up
    echo "Done."
}

############################################################################################

# Checking for the basic commands 
if [[ ! $(command -v go) ]]; then
    echo "Golang not found. Please install it to continue setup."
    echo "https://go.dev/doc/install"
    exit

elif [[ ! $(command -v mysql) ]]; then
    echo "MySQL not found. Please install it to continue setup."
    echo "https://ubuntu.com/server/docs/databases-mysql"
    exit

elif [[ ! $(command -v python3) ]]; then
    echo "Python3 not found. Please install it to continue setup."
    echo "https://www.python.org/downloads/"
    exit

elif [[ ! $(command -v pip) ]]; then
    echo "pip not found. Please install it to continue setup."
    exit

elif [[ ! $(command -v migrate) ]]; then
    echo "Migrate not found. Please install it to continue setup."
    echo "https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-ubuntu/"
    exit

else
    prettyPrint "All basic things in place :)"

fi

if [[ -f "config.yaml" ]]; then
    echo "config.yaml already exists. Do you want to overwrite it? (y/n) "
    read -p "Answer: " overwrite
    if [[ "$overwrite" == "y" ]]; then
        setupConfig
    fi
else
    setupConfig
fi

# Setting migrations
migrateDB

# Creating admin
prettyPrint "Creating an admin"
pip install mysql-connector-python
pip install bcrypt

python3 ./scripts/addAdmin.py

go mod tidy
go mod vendor