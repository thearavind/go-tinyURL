CREATE TABLE short_url(
   id SERIAL PRIMARY KEY     NOT NULL,
   hash           CHAR(50)    NOT NULL,
   url            TEXT     NOT NULL,
   clicks         INT
);