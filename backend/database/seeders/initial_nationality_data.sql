INSERT INTO nationalities (country, code, created_at, updated_at) VALUES
    ('Indonesia', 'IDN', NOW(), NOW()),
    ('Malaysia', 'MYS', NOW(), NOW()),
    ('Singapore', 'SGP', NOW(), NOW()),
    ('United States', 'USA', NOW(), NOW()) 
ON CONFLICT (code) DO NOTHING;