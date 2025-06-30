ALTER TABLE "website"
DROP CONSTRAINT IF EXISTS website_created_by_fkey;

ALTER TABLE "website"
DROP COLUMN IF EXISTS "created_by";