#! /bin/sh
setup_dir=${PWD}

set -e

curl -X POST -i "Content-Type: application/json" --data @./json/100_AF_NB_SUB_POST005.json http://localhost:8080/CNCA/1.0.1/subscriptions

exit 0

