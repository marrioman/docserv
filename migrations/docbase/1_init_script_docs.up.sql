CREATE SCHEMA IF NOT EXISTS docs;
CREATE EXTENSION IF NOT EXISTS "ltree";

CREATE TABLE IF NOT EXISTS docs.documents_ltree
(
    user_id  text,
    doc_type ltree,
    document_name varchar,
    document text
);

INSERT INTO docs.documents_ltree (user_id, doc_type, document_name, document)
VALUES ('48ab8b9d-e14c-49fe-bdd7-f439e64bb6ae', 'personal.ticket', 'билет в кино', '32548259fdfskfjkljlkajf'),
       ('ddc721a1-5aaf-4fee-a548-f65e678c28b7', 'personal.ticket', 'билет в кино', '32548259fdfskfjkljlkajf'),
       ('ddc721a1-5aaf-4fee-a548-f65e678c28b7', 'personal.avia', 'билет на самолет', 'пfапвы');
