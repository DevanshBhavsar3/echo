CREATE TYPE websiteStatus AS ENUM ('Up', 'Down', 'Unknown');

CREATE TABLE website (
                         "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         "url" TEXT NOT NULL,
	                     "health_check_route" TEXT DEFAULT '/',
                         "created_at" TIMESTAMP(3) NOT NULL DEFAULT NOW()
);

CREATE TABLE region (
                        "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        "name" TEXT NOT NULL
);

CREATE TABLE tick (
                              "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              "response_time_ms" INTEGER NOT NULL,
                              "status" websiteStatus NOT NULL,
                              "region_id" UUID NOT NULL,
                              "website_id" UUID NOT NULL,

                              FOREIGN KEY ("region_id")
                                  REFERENCES region("id") ON DELETE RESTRICT ON UPDATE CASCADE,

                              FOREIGN KEY ("website_id")
                                  REFERENCES website("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
