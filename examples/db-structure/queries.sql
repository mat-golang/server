-- When a user click the 'Skip' button, we want to destroy that room
-- and the rows for which users were in that room

DELETE FROM room_user_map WHERE room_id = <room_id>;
DELETE FROM rooms where room_id = <room_id>;