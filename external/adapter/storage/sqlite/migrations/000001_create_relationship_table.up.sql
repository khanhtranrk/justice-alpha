CREATE TABLE "relationships" (
    "owner_key" VARCHAR(32) NOT NULL,
    "guest_key" VARCHAR(32) NOT NULL,
    "safe_key" VARCHAR(32) NOT NULL,
    "gate" VARCHAR(255) NOT NULL
);

CREATE UNIQUE INDEX "relationships_index_on_owner_key" ON "relationships" ("owner_key");
