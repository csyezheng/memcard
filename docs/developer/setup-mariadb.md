# Setup Mysql

## Installation

```shell
sudo pacman -S mariadb
```

```shell
su - root
```

## Initial database
```shell
mariadb-install-db --user=mysql --basedir=/usr --datadir=/var/lib/mysql
```

## Start mariadb service

```shell
systemctl start mariadb.service
```
```shell
systemctl status mariadb.service
```

## Show user

```shell
mariadb
```
```shell
MariaDB [(none)]> SELECT User FROM mysql.user;
```

## Configuring root password
```
ALTER USER 'root'@'localhost' IDENTIFIED BY 'MyN3wP4ssw0rd';
flush privileges;
exit;
```
## Add user

```shell
mariadb
```

```shell
MariaDB> CREATE USER 'ye'@'localhost' IDENTIFIED BY '';
MariaDB> GRANT ALL PRIVILEGES ON mydb.* TO 'ye'@'localhost';
MariaDB> quit
```