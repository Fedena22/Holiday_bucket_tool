-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
Insert into Buckets (Placename, Latitude, Longitude, Visited ) Values ("Kyoto palace", "35.02509", "135.76193", 1);
Insert into Buckets (Placename, Latitude, Longitude, Visited ) Values ("Osaka trainstation", "34.7332", "135.49928", 0);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
Delete from Buckets;

