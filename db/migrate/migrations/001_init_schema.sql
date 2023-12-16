-- 001_init_schema.sql

-- Create deposits table 
CREATE TABLE IF NOT EXISTS deposits (
    guid             VARCHAR PRIMARY KEY,
    wallet_guid      VARCHAR NOT NULL,
    user_guid        VARCHAR NOT NULL,
    asset_guid       VARCHAR NOT NULL,
    transaction_guid VARCHAR NOT NULL,
    from_address     VARCHAR NOT NULL,
    to_address       VARCHAR NOT NULL,
    amount           VARCHAR NOT NULL,
    tx_hash          VARCHAR NOT NULL,
    tx_fee           VARCHAR NOT NULL,
    status           SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS deposits_guid ON deposits(guid);