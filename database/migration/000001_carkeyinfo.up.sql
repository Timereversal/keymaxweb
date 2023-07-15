CREATE TABLE IF NOT EXISTS carkeyinfo (
    id serial PRIMARY KEY ,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    brand text NOT NULL,
    model text NOT NULL,
    year integer NOT NULL,
    vin text NOT NULL,
    transponder text NOT NULL,
    blade text NOT NULL,
    frequency integer NOT NULL,
    partnumber text NOT NULL,
    keypicture text NOT NULL,
    remoteformat text NOT NULL,
    programmingprocedure text NOT NULL,
    cutinfo text NOT NULL ,
    cutinfopic text

)