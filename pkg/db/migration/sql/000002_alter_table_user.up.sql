DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'users'
        AND column_name = 'otp'
        AND column_name = 'is_active'
    ) THEN
ALTER TABLE users ADD COLUMN otp int DEFAULT 0;
ALTER TABLE users ADD COLUMN is_active bool DEFAULT false;
END IF;
END $$;
