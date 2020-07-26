CREATE TABLE `urls` (
        `id` INTEGER PRIMARY KEY AUTOINCREMENT,
        `key` VARCHAR(64),
        `url` Text,
        `banned` Boolean,
        `created_at` DATE NULL,
        UNIQUE(key)
    );