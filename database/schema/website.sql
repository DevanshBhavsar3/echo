CREATE TYPE websiteStatus AS ENUM ('Up', 'Down', 'Unknown');

CREATE TABLE website (
    "id" UUID PRIMARY KEY DEFAULT uuid(),
    "url" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT NOW()
);

CREATE TABLE region (
    "id" UUID PRIMARY KEY DEFAULT uuid(),
    "name" TEXT NOT NULL
);

CREATE TABLE websiteTicks (
    "id" UUID PRIMARY KEY DEFAULT uuid(),
    "response_time_ms" INTEGER NOT NULL,
    "status" websiteStatus NOT NULL,
    "regionId" UUID NOT NULL,
    "websiteId" UUID NOT NULL,

    FOREIGN KEY ("regionId")
    REFERENCES region("id") ON DELETE RESTRICT ON UPDATE CASCADE,

    FOREIGN KEY ("websiteId")
    REFERENCES website("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
