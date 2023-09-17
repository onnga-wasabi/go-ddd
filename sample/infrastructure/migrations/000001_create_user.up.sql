BEGIN;
CREATE TABLE IF NOT EXISTS users (
    id varchar(26),
    name varchar(50) NOT NULL,
    PRIMARY KEY (id)
);
END;
