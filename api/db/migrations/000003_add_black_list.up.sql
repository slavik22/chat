CREATE TABLE "black_list" (
                           "id" bigserial PRIMARY KEY,
                           "user_id" bigserial NOT NULL,
                           "friend_id" bigserial NOT NULL,
                           "created_at" timestamp DEFAULT NOW()
);

ALTER TABLE "black_list" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "black_list" ADD FOREIGN KEY ("friend_id") REFERENCES "users" ("id");