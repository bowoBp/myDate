

-- User table
CREATE TABLE IF NOT EXISTS "users" (
                        "user_id" SERIAL PRIMARY KEY,
                        "user_name" VARCHAR(255) NOT NULL,
                        "email" VARCHAR(255) UNIQUE NOT NULL,
                        "password" VARCHAR(255) NOT NULL,
                        "registration_date" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        "premium_status" BOOLEAN DEFAULT false
);

-- Profile table
CREATE TABLE IF NOT EXISTS "profile" (
                           "profile_id" SERIAL PRIMARY KEY,
                           "user_id" INT NOT NULL,
                           "photo_url" VARCHAR(255),
                           "description" TEXT,
                           "age" INT,
                           "location" VARCHAR(255),
                           FOREIGN KEY ("user_id") REFERENCES "users" ("user_id")
);

-- Swipe History table
CREATE TABLE IF NOT EXISTS "swipe_history" (
                                 "swipe_id" SERIAL PRIMARY KEY,
                                 "user_id" INT NOT NULL,
                                 "profile_id" INT NOT NULL,
                                 "action" INT,
                                 "swipe_time" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 FOREIGN KEY ("user_id") REFERENCES "users" ("user_id"),
                                 FOREIGN KEY ("profile_id") REFERENCES "profile" ("profile_id")
);

-- Premium Package table
CREATE TABLE IF NOT EXISTS "premium_package" (
                                   "package_id" SERIAL PRIMARY KEY,
                                   "package_name" VARCHAR(255) NOT NULL,
                                   "duration_days" INT,
                                   "price" DECIMAL(10,2),
                                   "feature_description" TEXT
);

-- Transaction table
CREATE TABLE IF NOT EXISTS "transaction" (
                               "transaction_id" SERIAL PRIMARY KEY,
                               "user_id" INT NOT NULL,
                               "package_id" INT NOT NULL,
                               "purchase_time" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               "payment_status" BOOLEAN DEFAULT false,
                               FOREIGN KEY ("user_id") REFERENCES "users" ("user_id"),
                               FOREIGN KEY ("package_id") REFERENCES "premium_package" ("package_id")
);
