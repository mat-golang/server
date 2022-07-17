-- Entities use a UUID for their id, whilst relationships use incrementing bigints

-- My demo assumes that users are registered users, and not anonymous
CREATE TABLE users (
  user_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  username varchar(20) NOT NULL,
  email text NOT NULL,
  created_at timestamp NOT NULL DEFAULT now() -- when the user registered
);

-- Rooms are pretty much just IDs...
CREATE TABLE rooms (
  room_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  user_limit integer NOT NULL DEFAULT 2,
  created_at timestamp NOT NULL DEFAULT now() -- when the room was created (uptime)
);

-- Map users to which room they are in. This means we can start with 2 users per room but easily change that limit
CREATE TABLE room_user_map (
  room_user_map_id bigserial PRIMARY KEY,
  user_id uuid REFERENCES users (user_id),
  created_at timestamp NOT NULL DEFAULT now() -- what time the user joined
);