ALTER TABLE "website"
ADD "created_by" UUID NOT NULL,

ADD CONSTRAINT website_created_by_fkey
FOREIGN KEY ("created_by") REFERENCES "user"("id") ON DELETE RESTRICT ON UPDATE CASCADE;