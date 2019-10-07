--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5 (Debian 11.5-3.pgdg90+1)
-- Dumped by pg_dump version 11.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: clients; Type: TABLE; Schema: public; Owner: hosting
--

CREATE TABLE public.clients (
    id bigint NOT NULL,
    login text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    source smallint
);


ALTER TABLE public.clients OWNER TO hosting;

--
-- Name: clients_id_seq; Type: SEQUENCE; Schema: public; Owner: hosting
--

CREATE SEQUENCE public.clients_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.clients_id_seq OWNER TO hosting;

--
-- Name: clients_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hosting
--

ALTER SEQUENCE public.clients_id_seq OWNED BY public.clients.id;


--
-- Name: pools; Type: TABLE; Schema: public; Owner: hosting
--

CREATE TABLE public.pools (
    id bigint NOT NULL,
    address inet NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    active boolean DEFAULT false NOT NULL,
    workload bigint DEFAULT 0 NOT NULL
);


ALTER TABLE public.pools OWNER TO hosting;

--
-- Name: pools_id_seq; Type: SEQUENCE; Schema: public; Owner: hosting
--

CREATE SEQUENCE public.pools_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pools_id_seq OWNER TO hosting;

--
-- Name: pools_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hosting
--

ALTER SEQUENCE public.pools_id_seq OWNED BY public.pools.id;


--
-- Name: virtual_hosts; Type: TABLE; Schema: public; Owner: hosting
--

CREATE TABLE public.virtual_hosts (
    id bigint NOT NULL,
    pool_id bigint NOT NULL,
    client_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.virtual_hosts OWNER TO hosting;

--
-- Name: virtual_hosts_id_seq; Type: SEQUENCE; Schema: public; Owner: hosting
--

CREATE SEQUENCE public.virtual_hosts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.virtual_hosts_id_seq OWNER TO hosting;

--
-- Name: virtual_hosts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hosting
--

ALTER SEQUENCE public.virtual_hosts_id_seq OWNED BY public.virtual_hosts.id;


--
-- Name: clients id; Type: DEFAULT; Schema: public; Owner: hosting
--

ALTER TABLE ONLY public.clients ALTER COLUMN id SET DEFAULT nextval('public.clients_id_seq'::regclass);


--
-- Name: pools id; Type: DEFAULT; Schema: public; Owner: hosting
--

ALTER TABLE ONLY public.pools ALTER COLUMN id SET DEFAULT nextval('public.pools_id_seq'::regclass);


--
-- Name: virtual_hosts id; Type: DEFAULT; Schema: public; Owner: hosting
--

ALTER TABLE ONLY public.virtual_hosts ALTER COLUMN id SET DEFAULT nextval('public.virtual_hosts_id_seq'::regclass);


--
-- Data for Name: clients; Type: TABLE DATA; Schema: public; Owner: hosting
--

COPY public.clients (id, login, created_at, deleted_at, source) FROM stdin;
\.


--
-- Data for Name: pools; Type: TABLE DATA; Schema: public; Owner: hosting
--

COPY public.pools (id, address, created_at, updated_at, deleted_at, active, workload) FROM stdin;
\.


--
-- Data for Name: virtual_hosts; Type: TABLE DATA; Schema: public; Owner: hosting
--

COPY public.virtual_hosts (id, pool_id, client_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Name: clients_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hosting
--

SELECT pg_catalog.setval('public.clients_id_seq', 1, false);


--
-- Name: pools_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hosting
--

SELECT pg_catalog.setval('public.pools_id_seq', 1, false);


--
-- Name: virtual_hosts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hosting
--

SELECT pg_catalog.setval('public.virtual_hosts_id_seq', 1, false);


--
-- Name: clients clients_login_key; Type: CONSTRAINT; Schema: public; Owner: hosting
--

ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_login_key UNIQUE (login);


--
-- Name: clients clients_pkey; Type: CONSTRAINT; Schema: public; Owner: hosting
--

ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_pkey PRIMARY KEY (id);


--
-- Name: pools pools_pkey; Type: CONSTRAINT; Schema: public; Owner: hosting
--

ALTER TABLE ONLY public.pools
    ADD CONSTRAINT pools_pkey PRIMARY KEY (id);


--
-- Name: virtual_hosts virtual_hosts_pkey; Type: CONSTRAINT; Schema: public; Owner: hosting
--

ALTER TABLE ONLY public.virtual_hosts
    ADD CONSTRAINT virtual_hosts_pkey PRIMARY KEY (id);


--
-- Name: virtual_hosts virtual_hosts_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hosting
--

ALTER TABLE ONLY public.virtual_hosts
    ADD CONSTRAINT virtual_hosts_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.clients(id);


--
-- Name: virtual_hosts virtual_hosts_pool_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: hosting
--

ALTER TABLE ONLY public.virtual_hosts
    ADD CONSTRAINT virtual_hosts_pool_id_fkey FOREIGN KEY (pool_id) REFERENCES public.pools(id);


--
-- PostgreSQL database dump complete
--

