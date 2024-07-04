#!/usr/bin/env bash

set -e

host="$1"
shift
port="$1"
shift

cmd="$@"

until nc -z "$host" "$port"; do
  >&2 echo "PostgreSQL is unavailable - sleeping"
  sleep 1
done

>&2 echo "PostgreSQL is up - executing command"
exec $cmd