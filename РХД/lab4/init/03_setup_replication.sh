#!/bin/sh
set -e

psql -v ON_ERROR_STOP=1 \
    --username "$POSTGRES_USER" \
    --dbname "$POSTGRES_DB" <<-'EOSQL'
    DO $$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'replicator') THEN
            CREATE USER replicator
                WITH REPLICATION
                ENCRYPTED PASSWORD 'rut_replicator_password';
        END IF;
    END $$;

    SELECT pg_create_physical_replication_slot('replication_slot_slave1', true)
    WHERE NOT EXISTS (
        SELECT FROM pg_replication_slots
        WHERE slot_name = 'replication_slot_slave1'
    );
EOSQL
