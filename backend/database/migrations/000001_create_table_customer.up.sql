CREATE TABLE customers (
    cst_id SERIAL PRIMARY KEY,
    nationality_id INTEGER NOT NULL,
    cst_name VARCHAR(50) NOT NULL,
    cst_dob DATE NOT NULL,
    cst_phoneNum VARCHAR(20) NOT NULL,
    cst_email VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);