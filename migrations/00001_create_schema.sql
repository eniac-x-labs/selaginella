CREATE TABLE IF NOT EXISTS l2_contract_events (
    -- Searchable fields
                                                  guid             VARCHAR PRIMARY KEY,
                                                  block_hash       VARCHAR NOT NULL REFERENCES l2_block_headers(hash) ON DELETE CASCADE,
    contract_address VARCHAR NOT NULL,
    transaction_hash VARCHAR NOT NULL,
    log_index        INTEGER NOT NULL,
    event_signature  VARCHAR NOT NULL, -- bytes32(0x0) when topics are missing
    timestamp        INTEGER NOT NULL CHECK (timestamp > 0),

    -- Raw Data
    rlp_bytes VARCHAR NOT NULL
    );
CREATE INDEX IF NOT EXISTS l2_contract_events_timestamp ON l2_contract_events(timestamp);
CREATE INDEX IF NOT EXISTS l2_contract_events_block_hash ON l2_contract_events(block_hash);
CREATE INDEX IF NOT EXISTS l2_contract_events_event_signature ON l2_contract_events(event_signature);
CREATE INDEX IF NOT EXISTS l2_contract_events_contract_address ON l2_contract_events(contract_address);

