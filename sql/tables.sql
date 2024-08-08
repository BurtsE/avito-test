CREATE TABLE houses IF NOT EXISTS(
    uuid SERIAL NOT NULL PRIMARY KEY,
    street VARCHAR(128) NOT NULL,
    construction_date DATE NOT NULL,
    developer   VARCHAR(128) DEFAULT "",
    initialization_date DATE NOT NULL
    last_update_time TIMESTAMP NOT NULL,
    flats_number INTEGER
)

CREATE TABLE flats(
    id SERIAL NOT NULL,
    unit_number SMALLINT NOT NULL,
    price INTEGER NOT NULL,
    room_number,
    moderation_status,
    house_id INTEGER NOT NULL REFERENCES houses(uuid)

)