CREATE TABLE IF NOT EXISTS movies (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    ranking integer NOT NULL,
    floors integer NOT NULL,
    location text[] NOT NULL,
    description text[] NOT NULL,
    );