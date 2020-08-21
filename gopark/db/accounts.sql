\connect Gopark

DROP TABLE IF EXISTS public.accounts;
CREATE TABLE public.accounts
(
    username character varying(100) COLLATE pg_catalog."default" NOT NULL,
    password character varying(20) COLLATE pg_catalog."default" NOT NULL,
    channel character varying(10) COLLATE pg_catalog."default" NOT NULL,
    display_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    active boolean NOT NULL,
    CONSTRAINT "Accounts_pkey" PRIMARY KEY (username)
);