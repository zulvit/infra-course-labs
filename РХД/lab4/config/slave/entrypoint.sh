#!/bin/sh
set -e

if [ -z "$(ls -A /var/lib/postgresql/data 2>/dev/null)" ]; then
    echo ">>> Data directory is empty. Running pg_basebackup from master..."

    PGPASSWORD=rut_replicator_password pg_basebackup \
        --host=postgres_master \
        --port=5432 \
        --username=replicator \
        --pgdata=/var/lib/postgresql/data \
        --slot=replication_slot_slave1 \
        --wal-method=stream \
        --checkpoint=fast \
        --write-recovery-conf \
        --progress

    echo ">>> pg_basebackup complete."
    chmod 0700 /var/lib/postgresql/data
fi

exec postgres \
    -c config_file=/etc/postgresql/postgresql.conf \
    -c hba_file=/etc/postgresql/pg_hba.conf
