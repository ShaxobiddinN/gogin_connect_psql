BEGIN;

DELETE FROM article WHERE id =  '26e2aebc-9771-45ba-8577-ef1a2e7b4170';
DELETE FROM article WHERE id =  '9900756f-e3ed-4dd7-a3a8-4e3cef248ccc';
DELETE FROM article WHERE id =  '3e451dc4-42e8-4dbc-a70b-edee8f6452ba';

DELETE FROM author WHERE id = '24000e82-9c48-4297-a442-ecd1ad55791e';
DELETE FROM author WHERE id = '3e1dfc06-dcf6-41fc-b3cc-7c0563fdfab3';

COMMIT;