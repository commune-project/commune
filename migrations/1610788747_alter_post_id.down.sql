LOCK posts;
SELECT setval('posts_id_seq', (SELECT MAX(id) FROM posts));
ALTER TABLE posts
    ALTER COLUMN id
    SET DEFAULT nextval('posts_id_seq');