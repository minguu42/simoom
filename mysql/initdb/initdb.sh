#!/usr/bin/env bash
set -euo pipefail

db_name="simoomdb"
mysql -u root -e "CREATE DATABASE IF NOT EXISTS ${db_name}"
mysql -u root ${db_name} < /schema.sql
mysql -u root ${db_name} <<EOF
INSERT INTO users (id, name, api_key, created_at, updated_at)
VALUES ('01DXF6DT000000000000000000', 'minguu42', 'rAM9Fm9huuWEKLdCwHBcju9Ty_-TL2tDsAicmMrXmUnaCGp3RtywzYpMDPdEtYtR', '2020-01-01 00:00:00', '2020-01-01 00:00:00')
EOF

test_db_name="${db_name}_test"
mysql -u root -e "CREATE DATABASE IF NOT EXISTS ${test_db_name}"
mysql -u root ${test_db_name} < /schema.sql
