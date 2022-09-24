CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
 "id" UUID NOT NULL DEFAULT uuid_generate_v4(),
 "user" varchar(100) NOT NULL,
 "password" varchar(100) NOT NULL,
 "valid" boolean DEFAULT true,
 "expires" timestamp NOT NULL DEFAULT (NOW() + interval '1 year'),
 "paying" boolean DEFAULT NULL,
 PRIMARY KEY ("id"),
 CONSTRAINT "user" UNIQUE ("user")
);

CREATE INDEX "expires" ON "users" ("expires");