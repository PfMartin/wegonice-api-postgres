CREATE TABLE "users" (
  "email" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "authors" (
  "id" bigserial PRIMARY KEY,
  "author_name" varchar NOT NULL,
  "website" varchar NOT NULL DEFAULT '',
  "instagram" varchar NOT NULL DEFAULT '',
  "youtube" varchar NOT NULL DEFAULT '',
  "user_created" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "recipes" (
  "id" bigserial PRIMARY KEY,
  "recipe_name" varchar NOT NULL,
  "link" varchar NOT NULL DEFAULT '',
  "author_id" bigserial NOT NULL,
  "prep_time" float NOT NULL DEFAULT 0,
  "prep_time_unit" varchar NOT NULL DEFAULT 'min',
  "user_created" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "recipe_ingredients" (
  "id" bigserial PRIMARY KEY,
  "rank" int NOT NULL,
  "ingredient_name" varchar NOT NULL,
  "unit" varchar NOT NULL,
  "amount" float NOT NULL DEFAULT 0,
  "recipe_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "recipe_steps" (
  "id" bigserial PRIMARY KEY,
  "rank" int NOT NULL,
  "step_description" varchar NOT NULL,
  "recipe_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "authors" ("id");

CREATE INDEX ON "authors" ("author_name");

CREATE INDEX ON "recipes" ("id");

CREATE INDEX ON "recipes" ("recipe_name");

CREATE INDEX ON "recipe_ingredients" ("id");

CREATE INDEX ON "recipe_ingredients" ("ingredient_name");

CREATE INDEX ON "recipe_steps" ("id");

COMMENT ON COLUMN "recipe_ingredients"."amount" IS 'cannot be negative';

ALTER TABLE "authors" ADD FOREIGN KEY ("user_created") REFERENCES "users" ("email");

ALTER TABLE "recipes" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "recipes" ADD FOREIGN KEY ("user_created") REFERENCES "users" ("email");

ALTER TABLE "recipe_ingredients" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipes" ("id");

ALTER TABLE "recipe_steps" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipes" ("id");
