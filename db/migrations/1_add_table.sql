-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
Create Table Buckets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    Placename TEXT(255),
    Latitude TEXT(255),
    Longitude TEXT(255),
    Visited int(1)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
Drop Table Buckets;