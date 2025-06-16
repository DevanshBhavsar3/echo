CREATE TABLE "website_region" (
    "website_id" UUID NOT NULL,
    "region_id" UUID NOT NULL,

    PRIMARY KEY ("website_id", "region_id"),

    FOREIGN KEY ("website_id")
        REFERENCES website("id") ON DELETE RESTRICT ON UPDATE CASCADE,

    FOREIGN KEY ("region_id")
        REFERENCES region("id") ON DELETE RESTRICT ON UPDATE CASCADE
)