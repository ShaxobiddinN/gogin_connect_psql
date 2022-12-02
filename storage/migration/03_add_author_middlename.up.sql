ALTER TABLE author ADD COLUMN middlename VARCHAR(100);

UPDATE author SET middlename = ' ' where middlename is NULL;