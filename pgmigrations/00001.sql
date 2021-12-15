CREATE TABLE "urls" (
  "id" SERIAL,
  "url" varchar NOT NULL UNIQUE,
  PRIMARY KEY ("id")
);
