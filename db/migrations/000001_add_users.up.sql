CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "name" varchar NOT NULL,
                         "login" varchar UNIQUE NOT NULL,
                         "hashed_password" varchar NOT NULL
);

CREATE UNIQUE INDEX ON "users" ("login");