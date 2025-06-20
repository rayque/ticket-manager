CREATE TABLE packages
(
    id          SERIAL PRIMARY KEY,
    uuid     VARCHAR(255) UNIQUE,
    product     VARCHAR(255) NOT NULL,
    weight      FLOAT NOT NULL,
    destination CHAR(2)      NOT NULL,
    status      VARCHAR(50)  NOT NULL,
    carrier_uuid     VARCHAR(255),
    created_at  TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP(3) NOT NULL
);

-- CREATE INDEX idx_packages_uuid ON packages (uuid);