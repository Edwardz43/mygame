------------------------ Table: public.member ------------------------------------

-- DROP TABLE public.member;

CREATE TABLE public.member
(
    id integer NOT NULL DEFAULT nextval('member_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    password text COLLATE pg_catalog."default" NOT NULL,
    email character varying(100) COLLATE pg_catalog."default" NOT NULL,
    balance money DEFAULT 0,
    islogin boolean DEFAULT false,
    CONSTRAINT member_pkey PRIMARY KEY (id),
    CONSTRAINT email_unique UNIQUE (email)
,
    CONSTRAINT name_unique UNIQUE (name)

)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.member
    OWNER to admin;

COMMENT ON COLUMN public.member.balance
    IS 'account balance of member';

COMMENT ON COLUMN public.member.islogin
    IS 'check if login';

-- Index: idx_member_deleted_at

-- DROP INDEX public.idx_member_deleted_at;

CREATE INDEX idx_member_deleted_at
    ON public.member USING btree
    (deleted_at)
    TABLESPACE pg_default;

-- Index: name_email

-- DROP INDEX public.name_email;

CREATE UNIQUE INDEX name_email
    ON public.member USING btree
    (name COLLATE pg_catalog."C.UTF-8" text_pattern_ops, email COLLATE pg_catalog."C.UTF-8" varchar_ops)
    TABLESPACE pg_default;

-- Index: uix_member_email

-- DROP INDEX public.uix_member_email;

CREATE UNIQUE INDEX uix_member_email
    ON public.member USING btree
    (email COLLATE pg_catalog."default")
    TABLESPACE pg_default;

-------------------------- Table: public.lobby ---------------------------------

-- DROP TABLE public.lobby;

CREATE TABLE public.lobby
(
    id integer NOT NULL DEFAULT nextval('lobby_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    game_id smallint NOT NULL,
    run bigint NOT NULL,
    inn integer NOT NULL,
    status smallint NOT NULL,
    CONSTRAINT lobby_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.lobby
    OWNER to admin;

-- Index: idx_lobby_deleted_at

-- DROP INDEX public.idx_lobby_deleted_at;

CREATE INDEX idx_lobby_deleted_at
    ON public.lobby USING btree
    (deleted_at)
    TABLESPACE pg_default;


-------------------------- Table: public.game_result --------------------------

-- DROP TABLE public.game_result;

CREATE TABLE public.game_result
(
    id integer NOT NULL DEFAULT nextval('game_result_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    game_id integer,
    run bigint,
    inn integer,
    detail text COLLATE pg_catalog."default",
    mod_times integer,
    CONSTRAINT game_result_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.game_result
    OWNER to admin;

-- Index: idx_game_result_deleted_at

-- DROP INDEX public.idx_game_result_deleted_at;

CREATE INDEX idx_game_result_deleted_at
    ON public.game_result USING btree
    (deleted_at)
    TABLESPACE pg_default;


----------------------- Table: public.bet_distincts --------------------------

-- DROP TABLE public.bet_distincts;

CREATE TABLE public.bet_distincts
(
    id integer NOT NULL DEFAULT nextval('bet_distincts_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    game_id integer,
    "distinct" text COLLATE pg_catalog."default",
    win_flag boolean,
    CONSTRAINT bet_distincts_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.bet_distincts
    OWNER to admin;

-- Index: idx_bet_distincts_deleted_at

-- DROP INDEX public.idx_bet_distincts_deleted_at;

CREATE INDEX idx_bet_distincts_deleted_at
    ON public.bet_distincts USING btree
    (deleted_at)
    TABLESPACE pg_default;

-------------------- Table: public.bet_records ----------------------------------

-- DROP TABLE public.bet_records;

CREATE TABLE public.bet_records
(
    id integer NOT NULL DEFAULT nextval('bet_records_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    game_id integer,
    run bigint,
    inn integer,
    "distinct" integer,
    amount bigint,
    CONSTRAINT bet_records_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.bet_records
    OWNER to admin;

-- Index: idx_bet_records_deleted_at

-- DROP INDEX public.idx_bet_records_deleted_at;

CREATE INDEX idx_bet_records_deleted_at
    ON public.bet_records USING btree
    (deleted_at)
    TABLESPACE pg_default;


-------------------------- Table: public.bet_results --------------------------------

-- DROP TABLE public.bet_results;

CREATE TABLE public.bet_results
(
    id integer NOT NULL DEFAULT nextval('bet_results_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    bet_record_id integer,
    win_flag boolean,
    win_distinct text COLLATE pg_catalog."default",
    "distinct" text COLLATE pg_catalog."default",
    bet_amount integer,
    bet_win integer,
    CONSTRAINT bet_results_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.bet_results
    OWNER to admin;

-- Index: idx_bet_results_deleted_at

-- DROP INDEX public.idx_bet_results_deleted_at;

CREATE INDEX idx_bet_results_deleted_at
    ON public.bet_results USING btree
    (deleted_at)
    TABLESPACE pg_default;
