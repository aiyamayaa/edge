#! /bin/sh

setup_dir=${PWD}

set -e

curl -X POST -i "Content-Type: application/json" --data @./json/100_AF_NB_SUB_POST01.json http://localhost:8080/CNCA/1.0.1/subscriptions
curl -X POST -i "Content-Type: application/json" --data @./json/100_AF_NB_SUB_POST02.json http://localhost:8080/CNCA/1.0.1/subscriptions
curl -X POST -i "Content-Type: application/json" --data @./json/100_AF_NB_SUB_POST03.json http://localhost:8080/CNCA/1.0.1/subscriptions
curl -X POST -i "Content-Type: application/json" --data @./json/100_AF_NB_SUB_POST04.json http://localhost:8080/CNCA/1.0.1/subscriptions
curl -X POST -i "Content-Type: application/json" --data @./json/100_AF_NB_SUB_POST05.json http://localhost:8080/CNCA/1.0.1/subscriptions

exit 0

