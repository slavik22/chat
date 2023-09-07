CREATE TABLE "friends" (
                           "user_id" bigserial NOT NULL,
                           "friend_id" bigserial NOT NULL,
                           "created_at" timestamp DEFAULT NOW(),
                            PRIMARY KEY(user_id, friend_id)
);

ALTER TABLE "friends" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "friends" ADD FOREIGN KEY ("friend_id") REFERENCES "users" ("id");
