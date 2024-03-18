CREATE TABLE IF NOT EXISTS public.role (
    id SERIAL PRIMARY KEY CONSTRAINT id_is_positive CHECK (id > 0),
    name TEXT UNIQUE NOT NULL
);

INSERT INTO
    public.role ("name")
VALUES
    ('admin');

CREATE TABLE IF NOT EXISTS public.user (
    id SERIAL PRIMARY KEY CONSTRAINT id_is_positive CHECK (id > 0),
    username TEXT UNIQUE NOT NULL CONSTRAINT valid_name CHECK (length(username) > 0)
);

INSERT INTO
    public.user (username)
VALUES
    ('test_user'),
    ('admin_user');

CREATE TABLE IF NOT EXISTS public.user_role_assign (
    user_id INT REFERENCES public.user(id),
    role_id INT REFERENCES public.role(id),
    PRIMARY KEY (user_id, role_id)
);

INSERT
    INTO
    public.user_role_assign
SELECT
    u.id,
    r.id
FROM
    public."user" u
JOIN public."role" r
ON
    TRUE
WHERE
    u.username = 'admin_user';


---- create above / drop below ----
DROP TABLE IF EXISTS public.user_roles_assign;

DROP TABLE IF EXISTS public.roles;

DROP TABLE IF EXISTS public.users;
