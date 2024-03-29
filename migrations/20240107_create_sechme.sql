DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'uint256') THEN
        CREATE DOMAIN UINT256 AS NUMERIC
            CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
ELSE
    ALTER DOMAIN UINT256 DROP CONSTRAINT uint256_check;
        ALTER DOMAIN UINT256 ADD
            CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
END IF;
END $$;

CREATE TABLE IF NOT EXISTS cross_chain_transfer (
    guid                  VARCHAR PRIMARY KEY,
    source_chain_id       UINT256 NOT NULL,
    dest_chain_id         UINT256 NOT NULL,
    source_hash           VARCHAR NOT NULL,
    tx_hash               VARCHAR NOT NULL,
    fee                   UINT256 NOT NULL,
    nonce                 UINT256 NOT NULL,
    source_sender_address VARCHAR NOT NULL,
    dest_receive_address  VARCHAR NOT NULL,
    token_address         VARCHAR NOT NULL,
    amount                UINT256 NOT NULL,
    status                SMALLINT DEFAULT 0,
    timestamp             INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS cross_chain_transfer_timestamp ON cross_chain_transfer(timestamp);
CREATE INDEX IF NOT EXISTS cross_chain_transfer_source_hash ON cross_chain_transfer(source_hash);
CREATE INDEX IF NOT EXISTS cross_chain_transfer_tx_hash ON cross_chain_transfer(tx_hash);
CREATE INDEX IF NOT EXISTS cross_chain_transfer_source_sender_address ON cross_chain_transfer(source_sender_address);
CREATE INDEX IF NOT EXISTS cross_chain_transfer_dest_receive_address ON cross_chain_transfer(dest_receive_address);
CREATE INDEX IF NOT EXISTS cross_chain_transfer_token_address ON cross_chain_transfer(token_address);

CREATE TABLE IF NOT EXISTS update_deposit_funding_pool_balance (
    guid                  VARCHAR PRIMARY KEY,
    source_chain_id       UINT256 NOT NULL,
    dest_chain_id         UINT256 NOT NULL,
    receive_address       VARCHAR NOT NULL,
    token_address         VARCHAR NOT NULL,
    amount                UINT256 NOT NULL,
    source_hash           VARCHAR NOT NULL,
    tx_hash               VARCHAR NOT NULL,
    status                SMALLINT DEFAULT 0,
    timestamp             INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS update_deposit_funding_pool_balance_timestamp ON update_deposit_funding_pool_balance(timestamp);
CREATE INDEX IF NOT EXISTS update_deposit_funding_pool_balance_source_hash ON update_deposit_funding_pool_balance(source_hash);
CREATE INDEX IF NOT EXISTS update_deposit_funding_pool_balance_tx_hash ON update_deposit_funding_pool_balance(tx_hash);
CREATE INDEX IF NOT EXISTS update_deposit_funding_pool_balance_dest_receive_address ON update_deposit_funding_pool_balance(receive_address);
CREATE INDEX IF NOT EXISTS update_deposit_funding_pool_balance_token_address ON update_deposit_funding_pool_balance(token_address);

CREATE TABLE IF NOT EXISTS update_withdraw_funding_pool_balance (
    guid                  VARCHAR PRIMARY KEY,
    source_chain_id       UINT256 NOT NULL,
    dest_chain_id         UINT256 NOT NULL,
    receive_address       VARCHAR NOT NULL,
    token_address         VARCHAR NOT NULL,
    amount                UINT256 NOT NULL,
    source_hash           VARCHAR NOT NULL,
    tx_hash               VARCHAR NOT NULL,
    status                SMALLINT DEFAULT 0,
    timestamp             INTEGER NOT NULL CHECK (timestamp > 0)
);
CREATE INDEX IF NOT EXISTS update_withdraw_funding_pool_balance_timestamp ON update_withdraw_funding_pool_balance(timestamp);
CREATE INDEX IF NOT EXISTS update_withdraw_funding_pool_balance_source_hash ON update_withdraw_funding_pool_balance(source_hash);
CREATE INDEX IF NOT EXISTS update_withdraw_funding_pool_balance_tx_hash ON update_withdraw_funding_pool_balance(tx_hash);
CREATE INDEX IF NOT EXISTS update_withdraw_funding_pool_balance_dest_receive_address ON update_withdraw_funding_pool_balance(receive_address);
CREATE INDEX IF NOT EXISTS update_withdraw_funding_pool_balance_token_address ON update_withdraw_funding_pool_balance(token_address);