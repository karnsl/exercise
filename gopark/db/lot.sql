\connect Gopark

DROP TABLE IF EXISTS public.lot;
CREATE TABLE public.lot
(
    id integer NOT NULL DEFAULT nextval('lot_id_seq'::regclass),
    place_id integer NOT NULL,
    building character varying(50) COLLATE pg_catalog."default" NOT NULL,
    floor character varying(10) COLLATE pg_catalog."default" NOT NULL,
    zone character varying(50) COLLATE pg_catalog."default" NOT NULL,
    "number" smallint NOT NULL,
    username character varying(100) COLLATE pg_catalog."default",
    CONSTRAINT lot_pkey PRIMARY KEY (id),
    CONSTRAINT place_id FOREIGN KEY (place_id)
        REFERENCES public.places (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT reserved_user FOREIGN KEY (username)
        REFERENCES public.accounts (username) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);