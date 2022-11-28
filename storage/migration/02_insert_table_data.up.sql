BEGIN;

INSERT INTO author (id, firstname, lastname) VALUES ('3e1dfc06-dcf6-41fc-b3cc-7c0563fdfab3', 'John', 'Doe') ON CONFLICT DO NOTHING;
INSERT INTO author (id, firstname, lastname) VALUES ('24000e82-9c48-4297-a442-ecd1ad55791e', 'Shaxobiddin', 'Najmiddinov' ) ON CONFLICT DO NOTHING;

INSERT INTO article (id, title, body, author_id) VALUES ('26e2aebc-9771-45ba-8577-ef1a2e7b4170', 'Lorem 1', 'Body 1', '3e1dfc06-dcf6-41fc-b3cc-7c0563fdfab3') ON CONFLICT DO NOTHING;
INSERT INTO article (id, title, body, author_id) VALUES ('9900756f-e3ed-4dd7-a3a8-4e3cef248ccc', 'Lorem 2', 'Body 2', '24000e82-9c48-4297-a442-ecd1ad55791e') ON CONFLICT DO NOTHING;

INSERT INTO article (id, title, body, author_id) VALUES ('3e451dc4-42e8-4dbc-a70b-edee8f6452ba', 'Lorem 3', 'Body 3', '3e1dfc06-dcf6-41fc-b3cc-7c0563fdfab3') ON CONFLICT DO NOTHING;


COMMIT;