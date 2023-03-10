CREATE TABLE "shops_categories" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "link" varchar NOT NULL
);

CREATE TABLE "shops_shops_categories" (
  "shop_category_id" int NOT NULL,
  "shop_id" int NOT NULL
);

CREATE TABLE "shops" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "link" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products_categories" (
  "shop_id" int,
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "link" varchar NOT NULL
);

CREATE TABLE "products_products_categories" (
  "product_category_id" int NOT NULL,
  "product_id" int NOT NULL
);

CREATE TABLE "products" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "links" varchar[],
  "price" decimal,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "shops" ("id");

CREATE INDEX ON "products" ("id");

ALTER TABLE "shops_shops_categories" ADD FOREIGN KEY ("shop_category_id") REFERENCES "shops_categories" ("id");

ALTER TABLE "shops_shops_categories" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("id");

ALTER TABLE "products_products_categories" ADD FOREIGN KEY ("product_category_id") REFERENCES "products_categories" ("id");

ALTER TABLE "products_products_categories" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
