  RelationshipKey string
  OwnerKey string
  GuestKey string
  GuestRequestQueue string
  GuestResponseQueue string
  RequestKey string
  ResponseKey string
  CreatedAt time.Time
  UpdatedAt time.Time

CREATE TABLE "relationships" (
    "id" BIGINT PRIMARY KEY AUTOINCREMENT,
    "owner_key" VARCHAR(32) NOT NULL,
    "guest_key" VARCHAR(32) NOT NULL,
    "guest_request_queue" VARCHAR(255) NOT NULL,
    "request_response_queue" VARCHAR(255) NOT NULL,
    "refresh_key_queue" VARCHAR(255) NOT NULL,
    "created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX "ownerKey" ON "guestKey" ("email");
