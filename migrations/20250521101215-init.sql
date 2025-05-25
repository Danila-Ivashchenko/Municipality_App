
-- +migrate Up

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

ALTER TABLE municipality_object_to_passport_partition
    ADD FOREIGN KEY (municipality_passport_partitition_id) REFERENCES municipality_passport_partitition(id) ON delete CASCADE;

ALTER TABLE municipality_object_to_passport_partition
    ADD FOREIGN KEY (municipality_object_id) REFERENCES municipality_object(id) ON delete CASCADE;

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

CREATE TABLE IF NOT EXISTS municipality_entity_type (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS municipality_entity_template (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    entity_type_id bigint NOT NULL,
    municipality_id bigint NOT NULL,
    FOREIGN KEY (entity_type_id) REFERENCES municipality_entity_type(id) ON delete cascade,
    FOREIGN KEY (municipality_id) REFERENCES municipality(id) ON delete cascade,
    UNIQUE (municipality_id, name)
);

CREATE TABLE IF NOT EXISTS municipality_entity (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    municipality_entity_template_id bigint NOT NULL,
    description text,
    FOREIGN KEY (municipality_entity_template_id) REFERENCES municipality_entity_template(id) ON delete cascade
);

CREATE TABLE IF NOT EXISTS municipality_entity_to_passport_partition (
    municipality_passport_partitition_id bigint NOT NULL,
    municipality_entity_id bigint NOT NULL,
    UNIQUE (municipality_passport_partitition_id, municipality_entity_id)
);

ALTER TABLE municipality_entity_to_passport_partition
    ADD FOREIGN KEY (municipality_passport_partitition_id) REFERENCES municipality_passport_partitition(id) ON delete CASCADE;

ALTER TABLE municipality_entity_to_passport_partition
    ADD FOREIGN KEY (municipality_entity_id) REFERENCES municipality_entity(id) ON delete CASCADE;

CREATE TABLE IF NOT EXISTS municipality_entity_attribute (
    id BIGSERIAL PRIMARY KEY,
    entity_template_id bigint NOT NULL,
    name VARCHAR(255) NOT NULL,
    default_value VARCHAR(255) NOT NULL,
    to_show bool NOT NULL default false,
    UNIQUE (entity_template_id, name),
    FOREIGN KEY (entity_template_id) REFERENCES municipality_entity_template(id) ON delete CASCADE
);

CREATE TABLE IF NOT EXISTS municipality_entity_attribute_value (
    id BIGSERIAL PRIMARY KEY,
    entity_attribute_id bigint NOT NULL,
    entity_id bigint NOT NULL,
    value VARCHAR(255) NOT NULL,
    UNIQUE (id, entity_attribute_id),
    FOREIGN KEY (entity_attribute_id) REFERENCES municipality_entity_attribute(id) ON delete CASCADE,
    FOREIGN KEY (entity_id) REFERENCES municipality_entity(id) ON delete CASCADE
);

CREATE TABLE IF NOT EXISTS route (
    id BIGSERIAL PRIMARY KEY,
    municipality_passport_partitition_id bigint NOT NULL,
    name VARCHAR(255) NOT NULL,
    length bigint,
    duration bigint NOT NULL,
    level bigint NOT NULL,
    movement_way text,
    seasonality text,
    personal_equipment text,
    dangers text,
    rules text,
    route_equipment text,
    geometry jsonb,
    FOREIGN KEY (municipality_passport_partitition_id) REFERENCES municipality_passport_partitition(id) ON delete cascade
);

CREATE TABLE IF NOT EXISTS municipality_route_object (
    id BIGSERIAL PRIMARY KEY,
    route_id bigint NOT NULL,
    object_name VARCHAR(255) NOT NULL,
    order_number int not null,
    source_object_id bigint,
    location_id bigint,
    FOREIGN KEY (route_id) REFERENCES route(id) ON delete cascade,
    FOREIGN KEY (source_object_id) REFERENCES municipality_object(id) ON delete set null,
    FOREIGN KEY (location_id) REFERENCES location(id) ON delete SET NULL,
    UNIQUE (object_name, route_id)
);

CREATE TABLE IF NOT EXISTS municipality_passport_file (
    id BIGSERIAL PRIMARY KEY,
    path varchar(500) NOT NULL,
    passport_id bigint NOT NULL,
    file_name varchar(500) NOT NULL,
    created_at TIMESTAMP Default now(),
    FOREIGN KEY (passport_id) REFERENCES municipality_passport(id) ON delete cascade,
    UNIQUE (passport_id, path)
);

CREATE TABLE IF NOT EXISTS user_permission (
    user_id bigint NOT NULL,
    permission bigint NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON delete SET NULL,
    UNIQUE (user_id, permission)
);

-- +migrate Down
-- (здесь должен быть код отката миграции, если нужно)
