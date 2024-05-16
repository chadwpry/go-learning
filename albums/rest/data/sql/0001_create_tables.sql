USE recordings;

DROP TABLE IF EXISTS albums;

CREATE TABLE albums (
  id         VARCHAR(40) DEFAULT (uuid()) NOT NULL PRIMARY KEY,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL
);
