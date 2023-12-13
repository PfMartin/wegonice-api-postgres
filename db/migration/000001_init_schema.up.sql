CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "authors" (
  "name" varchar PRIMARY KEY,
  "website" varchar,
  "instagram" varchar,
  "youtube" varchar,
  "user_created" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "recipes" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "link" varchar,
  "author_name" varchar NOT NULL,
  "prep_time" float NOT NULL DEFAULT 0,
  "prep_time_unit" varchar NOT NULL DEFAULT 'min',
  "user_created" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "recipe_ingredients" (
  "name" varchar PRIMARY KEY,
  "unit" varchar NOT NULL,
  "amount" float NOT NULL DEFAULT 0,
  "recipe_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "recipe_steps" (
  "id" bigserial PRIMARY KEY,
  "description" varchar NOT NULL,
  "recipe_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "authors" ("name");

CREATE INDEX ON "recipes" ("name");

CREATE INDEX ON "recipe_ingredients" ("name");

COMMENT ON COLUMN "recipe_ingredients"."amount" IS 'cannot be negative';

ALTER TABLE "authors" ADD FOREIGN KEY ("user_created") REFERENCES "users" ("username");

ALTER TABLE "recipes" ADD FOREIGN KEY ("author_name") REFERENCES "authors" ("name");

ALTER TABLE "recipes" ADD FOREIGN KEY ("user_created") REFERENCES "users" ("username");

ALTER TABLE "recipe_ingredients" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipes" ("id");

ALTER TABLE "recipe_steps" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipes" ("id");
