CREATE TABLE IF NOT EXISTS public.employees
(
    id serial NOT NULL,
    first_name text,
    last_name text,
    email text,
    hire_date date,
    CONSTRAINT employees_pkey PRIMARY KEY (id)
);