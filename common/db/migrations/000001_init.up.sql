CREATE TYPE "website_status" AS ENUM ('up', 'down', 'unknown');

CREATE TABLE "website" (
                         "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         "url" TEXT NOT NULL,
	                     "frequency" INTERVAL DEFAULT '3min',
                         "created_at" TIMESTAMP(3) NOT NULL DEFAULT NOW()
);

CREATE TABLE "region" (
                        "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        "name" TEXT UNIQUE NOT NULL
);

CREATE TABLE "website_tick" (
                              "id" UUID DEFAULT gen_random_uuid(),
                              "time" TIMESTAMPTZ NOT NULL,
                              "response_time_ms" INTEGER NOT NULL,
                              "status" "website_status" NOT NULL,
                              "region_id" UUID NOT NULL,
                              "website_id" UUID NOT NULL,

                              FOREIGN KEY ("region_id")
                                  REFERENCES region("id") ON DELETE RESTRICT ON UPDATE CASCADE,

                              FOREIGN KEY ("website_id")
                                  REFERENCES website("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
