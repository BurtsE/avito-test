CREATE TABLE IF NOT EXISTS houses (
    uuid SERIAL NOT NULL PRIMARY KEY,
    street VARCHAR(128) NOT NULL,
    construction_date DATE NOT NULL,
    developer   VARCHAR(128) DEFAULT '',
    initialization_date DATE NOT NULL,
    last_update_time TIMESTAMP NOT NULL,
    flats_number INTEGER NOT NULL DEFAULT
);

CREATE TABLE flats(
    id SERIAL NOT NULL,
    unit_number SMALLINT NOT NULL,
    price INTEGER NOT NULL,
    room_number SMALLINT,
    moderation_status VARCHAR(128),
    house_id INTEGER NOT NULL REFERENCES houses(uuid),
    moderator_id VARCHAR(128) default ''
);

CREATE TABLE users(
    user_id  VARCHAR(128) NOT NULL,
    email VARCHAR(128) NOT NULL,
    pass VARCHAR(128) NOT NULL,
    user_type VARCHAR(128) NOT NULL
);