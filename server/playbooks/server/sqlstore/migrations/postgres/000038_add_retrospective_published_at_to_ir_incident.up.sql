ALTER TABLE IR_Incident ADD COLUMN IF NOT EXISTS RetrospectivePublishedAt BIGINT NOT NULL DEFAULT 0;
