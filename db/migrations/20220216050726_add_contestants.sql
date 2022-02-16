-- migrate:up
CREATE TABLE contestants (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL UNIQUE,
  time_created timestamp DEFAULT now() NOT NULL,
  time_updated timestamp DEFAULT now() NOT NULL
);

CREATE TABLE scores (
  id bigserial PRIMARY KEY,
  contestant_id integer NOT NULL,
  user_id integer NOT NULL,
  time_created timestamp DEFAULT now() NOT NULL
);

ALTER TABLE users
ADD COLUMN time_updated timestamp DEFAULT now() NOT NULL;

-- migrate:down
ALTER TABLE users DROP COLUMN time_updated;
DROP TABLE scores;
DROP TABLE contestants;
