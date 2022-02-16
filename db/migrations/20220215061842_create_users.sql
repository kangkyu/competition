-- migrate:up
CREATE TABLE users (
  id serial PRIMARY KEY,
  email varchar(255) NOT NULL UNIQUE,
  time_created timestamp DEFAULT now() NOT NULL
);

-- migrate:down
DROP TABLE users;
