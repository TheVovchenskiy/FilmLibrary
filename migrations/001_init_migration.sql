CREATE TABLE IF NOT EXISTS public.movie (
    id serial PRIMARY KEY CONSTRAINT id_is_positive CHECK (id > 0),
    "name" TEXT UNIQUE NOT NULL CONSTRAINT valid_name CHECK (
        length("name") > 0
        AND length("name") <= 150
    ),
    description TEXT DEFAULT '' CONSTRAINT valid_description CHECK (
        length(description) >= 0
        AND length(description) <= 1000
    ),
    release_date DATE NOT NULL,
    rating INT NOT NULL CONSTRAINT valid_rating CHECK (
        rating >= 0
        AND rating <= 10
    ),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS public.star (
    id serial PRIMARY KEY CONSTRAINT id_is_positive CHECK (id > 0),
    "name" TEXT UNIQUE NOT NULL CONSTRAINT valid_name CHECK (length("name") > 0),
    gender CHAR(1) CONSTRAINT valid_gender CHECK (gender IN ('F', 'M')),
    birthday DATE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS public.movie_star_assign (
    movie_id int REFERENCES public.movie ON DELETE CASCADE,
    star_id int REFERENCES public.star ON DELETE CASCADE,
    PRIMARY KEY (movie_id, star_id)
);

---- create above / drop below ----
DROP TABLE IF EXISTS public.movie;

DROP TABLE IF EXISTS public.star;

DROP TABLE IF EXISTS public.movie_star_assign;
