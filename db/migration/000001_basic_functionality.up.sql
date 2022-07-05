CREATE TABLE public.users (
  id bigint NOT NULL DEFAULT unique_rowid(),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  tuition text NULL,
  password text NULL,
  firstname text NULL,
  lastname text NULL,
  age smallint NULL,
  insitute_email text NULL,
  status text NULL,
  "role" text NULL,
  nationality_id bigint NULL
);

ALTER TABLE
  public.users
ADD
  CONSTRAINT users_pkey PRIMARY KEY (id);

CREATE TABLE public.nacionalities (
  id bigint NOT NULL DEFAULT unique_rowid(),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  country text NULL,
  city text NULL
);

ALTER TABLE
  public.nacionalities
ADD
  CONSTRAINT nacionalities_pkey PRIMARY KEY (id);

CREATE TABLE public.trimesters (
  id bigint NOT NULL DEFAULT unique_rowid(),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  name text NULL,
  init_date timestamp with time zone NULL,
  finish_date timestamp with time zone NULL
);

ALTER TABLE
  public.trimesters
ADD
  CONSTRAINT trimesters_pkey PRIMARY KEY (id);

CREATE TABLE public.students (
  id bigint NOT NULL DEFAULT unique_rowid(),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  user_id bigint NULL,
  career text NULL,
  trimester_completed smallint NULL,
  pensum text NULL,
  state text NULL,
  quarterly_index numeric NULL,
  general_index numeric NULL
);

ALTER TABLE
  public.students
ADD
  CONSTRAINT students_pkey PRIMARY KEY (id);

CREATE TABLE public.teachers (
  id bigint NOT NULL DEFAULT unique_rowid(),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  user_id bigint NULL,
  academic_area text NULL
);

ALTER TABLE
  public.teachers
ADD
  CONSTRAINT teachers_pkey PRIMARY KEY (id);

CREATE TABLE public.admins (
  id bigint NOT NULL DEFAULT unique_rowid(),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  user_id bigint NULL
);

ALTER TABLE
  public.admins
ADD
  CONSTRAINT admins_pkey PRIMARY KEY (id);
  
CREATE TABLE public.courses (
  id bigint NOT NULL DEFAULT unique_rowid(),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  code text NULL,
  name text NULL,
  career text NULL,
  credits bigint NULL,
  academic_area text NULL
);

ALTER TABLE
  public.courses
ADD
  CONSTRAINT courses_pkey PRIMARY KEY (id);
  
CREATE TABLE public.ratings (
  id bigint NOT NULL DEFAULT unique_rowid(),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  student_id bigint NULL,
  course_id bigint NULL,
  rating bigint NULL,
  rating_letter text NULL
);

ALTER TABLE
  public.ratings
ADD
  CONSTRAINT ratings_pkey PRIMARY KEY (id);
  
CREATE TABLE public.sections (
  id bigint NOT NULL DEFAULT unique_rowid(),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  teacher_id bigint NULL,
  course_id bigint NULL,
  code text NULL,
  schedule text NULL,
  quota bigint NULL
);

ALTER TABLE
  public.sections
ADD
  CONSTRAINT sections_pkey PRIMARY KEY (id);
  
CREATE TABLE public.selections (
  id bigint NOT NULL DEFAULT unique_rowid(),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  code text NULL,
  trimester_id bigint NULL,
  year text NULL,
  student_id bigint NULL
);

ALTER TABLE
  public.selections
ADD
  CONSTRAINT selections_pkey PRIMARY KEY (id);
  
CREATE TABLE public.selection_records (
  id bigint NOT NULL DEFAULT unique_rowid(),
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  deleted_at timestamp with time zone NULL,
  section_id bigint NULL,
  selection_id bigint NULL
);

ALTER TABLE
  public.selection_records
ADD
  CONSTRAINT selection_records_pkey PRIMARY KEY (id);
  