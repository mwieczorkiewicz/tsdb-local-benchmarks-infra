-- public.load_profile definition

-- Drop table

DROP TABLE IF EXISTS public.load_profile;

CREATE TABLE public.load_profile (
	ts timestamptz NOT NULL,
	load_profile numeric NOT NULL,
	status text NOT NULL
);

DROP TABLE IF EXISTS public.weather_num;

CREATE TABLE public.weather_num (
	ts timestamptz NOT NULL,
	"location" text NOT NULL,
	"parameter" text NOT NULL,
	value numeric NOT NULL,
	unit text NOT NULL
);

DROP TABLE IF EXISTS public.weather_string;

CREATE TABLE public.weather_string (
	ts timestamptz NOT NULL,
	"location" text NOT NULL,
	"parameter" text NOT NULL,
	value text NOT NULL,
	unit text NOT NULL
);
alter table public.load_profile owner to postgres;
alter table public.weather_num owner to postgres;
alter table public.weather_string owner to postgres;