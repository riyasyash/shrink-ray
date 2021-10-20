CREATE TABLE "urls" (
        "id" BIGSERIAL PRIMARY KEY,
        "key" VARCHAR(64),
        "url" Text,
        "banned" Boolean,
        "created_at" DATE NULL,
        UNIQUE(key)
    );