CREATE TABLE IF NOT EXISTS carkeyinfo (
    id bigserial PRIMARY KEY ,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    brand text NOT NULL,
    model text NOT NULL,
    year integer NOT NULL,
    vin text NOT NULL ,

)