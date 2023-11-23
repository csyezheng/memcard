#!/bin/zsh

# usage: ./scripts/dev_set_envs.sh --db-engine postgresql --log-level INFO --log-form text

format=$(getopt -n "$0" -l "db-engine:,log-level:,log-form:" -- -- "$@")

eval set -- "$format"

#Read the argument values
while [ $# -gt 0 ]
do
     case "$1" in
          --db-engine) DB_ENGINE="$2"; shift;;
          --log-level) LOG_LEVEL="$2"; shift;;
          --log-form) LOG_FORM="$2"; shift;;
          --) shift;;
     esac
     shift;
done

if [ -z "$DB_ENGINE" ]; then
	DB_ENGINE="postgresql"
fi

if [ -z "$LOG_LEVEL" ]; then
	LOG_LEVEL="INFO"
fi

if [ -z "$LOG_FORM" ]; then
	LOG_FORM="text"
fi

echo $DB_ENGINE
echo $LOG_LEVEL
echo $LOG_FORM

export DB_ENGINE=$DB_ENGINE
export $LOG_LEVEL=$LOG_LEVEL
export $LOG_FORM=$LOG_FORM

function set_postgresql_env() {
  export DB_ENGINE="postgresql"
  export DB_HOST="127.0.0.1"
  export DB_PORT=5432
  export DB_USER="root"
  export DB_PASSWORD="root"
  export DB_NAME="postgresql"
  export DB_SSLMODE="disable"
  echo "set postgresql env"
  echo "set postgresql environment variables successfully"
}

function set_mysql_env() {
  export DB_ENGINE="mysql"
  export DB_HOST="127.0.0.1"
  export DB_PORT=3306
  export DB_USER="root"
  export DB_PASSWORD="root"
  export DB_NAME="mysql"
  export DB_SSLMODE="false"
  echo "set mysql environment variables successfully"
}

case $DB_ENGINE in
   "postgresql")
      set_postgresql_env
      ;;
   "mysql")
      set_mysql_env
      ;;
   *)
     print "not supported database engine: $1\n"
     ;;
esac

