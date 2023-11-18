#!/usr/bin/env bash
set -euo pipefail

db_name="simoomdb"
mysql -u root -e "CREATE DATABASE IF NOT EXISTS ${db_name}"
mysql -u root ${db_name} < /schema.sql

test_db_name="${db_name}_test"
mysql -u root -e "CREATE DATABASE IF NOT EXISTS ${test_db_name}"
mysql -u root ${test_db_name} < /schema.sql
