
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_admin bool NOT NULL default false,
    is_blocked bool NOT NULL default false,
    created_at TIMESTAMP Default now()
);

CREATE TABLE IF NOT EXISTS user_auth_token (
    id BIGSERIAL PRIMARY KEY,
    user_id bigint NOT NULL,
    token VARCHAR(128) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP Default now(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS region (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    code char(3) null unique
);

CREATE TABLE IF NOT EXISTS municipality (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    is_hidden bool NOT NULL default false,
    region_id bigint NOT NULL,
    created_at TIMESTAMP Default now(),
    FOREIGN KEY (region_id) REFERENCES region(id)
);

CREATE TABLE IF NOT EXISTS municipality_passport (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    municipality_id bigint NOT NULL,
    description text,
    year char(4),
    revision_code varchar(11) UNIQUE NOT NULL,
    is_main bool NOT NULL,
    is_hidden bool NOT NULL default false,
    updated_at TIMESTAMP Default now(),
    created_at TIMESTAMP Default now(),
    FOREIGN KEY (municipality_id) REFERENCES municipality(id),
    UNIQUE (name, municipality_id)
);

CREATE UNIQUE INDEX unique_main_passport_per_municipality
    ON municipality_passport (municipality_id)
    WHERE is_main = true;

CREATE TABLE IF NOT EXISTS municipality_passport_chapter (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    municipality_passport_id bigint NOT NULL,
    description text,
    chapter_text text,
    order_number INTEGER CHECK (order_number > 0),
    FOREIGN KEY (municipality_passport_id) REFERENCES municipality_passport(id) ON delete cascade,
    UNIQUE (name, municipality_passport_id)
);

CREATE TABLE IF NOT EXISTS municipality_passport_partitition (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    municipality_passport_chapter_id bigint NOT NULL,
    description text,
    chapter_text text,
    order_number INTEGER CHECK (order_number > 0),
    FOREIGN KEY (municipality_passport_chapter_id) REFERENCES municipality_passport_chapter(id) ON delete cascade,
    UNIQUE (name, municipality_passport_chapter_id)
);

CREATE TABLE IF NOT EXISTS municipality_object_type (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS location (
    id BIGSERIAL PRIMARY KEY,
    address VARCHAR(255),
    latitude DECIMAL(9,6)  CHECK (latitude BETWEEN -90 AND 90),
    longitude DECIMAL(9,6)  CHECK (longitude BETWEEN -180 AND 180),
    geometry jsonb
);

CREATE TABLE IF NOT EXISTS municipality_object_template (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    object_type_id bigint NOT NULL,
    municipality_id bigint NOT NULL,
    FOREIGN KEY (object_type_id) REFERENCES municipality_object_type(id) ON delete cascade,
    FOREIGN KEY (municipality_id) REFERENCES municipality(id) ON delete cascade,
    UNIQUE (municipality_id, name)
);

CREATE TABLE IF NOT EXISTS municipality_object (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    municipality_object_template_id bigint NOT NULL,
    location_id bigint,
    description text,
    FOREIGN KEY (location_id) REFERENCES location(id) ON delete SET NULL,
    FOREIGN KEY (municipality_object_template_id) REFERENCES municipality_object_template(id) ON delete cascade
);

CREATE TABLE IF NOT EXISTS municipality_object_to_passport_partition (
    municipality_passport_partitition_id bigint NOT NULL,
    municipality_object_id bigint NOT NULL,
    UNIQUE (municipality_passport_partitition_id, municipality_object_id)
);

CREATE TABLE IF NOT EXISTS municipality_object_attribute (
    id BIGSERIAL PRIMARY KEY,
    object_template_id bigint NOT NULL,
    name VARCHAR(255) NOT NULL,
    default_value VARCHAR(255) NOT NULL,
    to_show bool NOT NULL default false,
    UNIQUE (object_template_id, name),
    FOREIGN KEY (object_template_id) REFERENCES municipality_object_template(id) ON delete CASCADE
);

CREATE TABLE IF NOT EXISTS municipality_object_attribute_value (
    id BIGSERIAL PRIMARY KEY,
    object_attribute_id bigint NOT NULL,
    object_id bigint NOT NULL,
    value VARCHAR(255) NOT NULL,
    UNIQUE (id, object_attribute_id),
    FOREIGN KEY (object_attribute_id) REFERENCES municipality_object_attribute(id) ON delete CASCADE,
    FOREIGN KEY (object_id) REFERENCES municipality_object(id) ON delete CASCADE
);

CREATE TABLE IF NOT EXISTS municipality_object_to_passport_partition (
    object_id bigint NOT NULL,
    partition_id bigint NOT NULL,
    UNIQUE (object_id, partition_id)
);