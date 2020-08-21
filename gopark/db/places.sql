\connect Gopark

DROP TABLE IF EXISTS public.places;
CREATE TABLE public.places
(
    id integer NOT NULL DEFAULT nextval('places_id_seq'::regclass),
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT places_pkey PRIMARY KEY (id)
);