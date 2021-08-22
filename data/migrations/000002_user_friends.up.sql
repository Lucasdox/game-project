CREATE table "user_friends" (
                        user_id uuid not null REFERENCES game.public.user (id) ON DELETE CASCADE,
                        friend_id uuid not null REFERENCES game.public.user (id) ON DELETE CASCADE,
                        PRIMARY KEY (user_id, friend_id),
                        CONSTRAINT no_duplicate_friends UNIQUE (user_id, friend_id)
);
