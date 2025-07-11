CREATE TABLE family_lists (
    fl_id SERIAL PRIMARY KEY,
    cst_id INTEGER NOT NULL,
    fl_name VARCHAR(50) NOT NULL,
    fl_relation VARCHAR(50) NOT NULL,
    fl_dob VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_cst
        FOREIGN KEY(cst_id)
        REFERENCES customers(cst_id)
        ON DELETE CASCADE
);

-- Menambahkan indeks untuk foreign key agar performa lebih baik
CREATE INDEX idx_family_list_cst_id ON family_lists (cst_id);