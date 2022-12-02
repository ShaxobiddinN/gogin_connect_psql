ALTER TABLE author ALTER COLUMN firstname DROP NOT NULL;
ALTER TABLE author ALTER COLUMN lastname DROP NOT NULL;
ALTER TABLE author ALTER COLUMN middlename DROP NOT NULL;

ALTER TABLE author 
RENAME COLUMN firstname TO temp_firstname;
ALTER TABLE author 
RENAME COLUMN lastname TO temp_lastname;
ALTER TABLE author 
RENAME COLUMN middlename TO temp_middlename;

