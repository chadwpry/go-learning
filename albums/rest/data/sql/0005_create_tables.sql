\connect recordings

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS albums;

CREATE TABLE albums (
  id         uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  title      text NOT NULL,
  artist     text NOT NULL,
  price      numeric(5,2) NOT NULL
);
