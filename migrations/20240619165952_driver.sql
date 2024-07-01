-- +goose Up
CREATE TABLE cars (
    id uuid PRIMARY KEY,
    fleet_id varchar(32) UNIQUE,
    brand text NOT NULL,
    model text NOT NULL,
    year int NOT NULL,
    color text NOT NULL,
    vin text NOT NULL,
    license_plate_number text NOT NULL,
    registration_certificate text NOT NULL
);

CREATE TABLE IF NOT EXISTS driver_license (
    id uuid PRIMARY KEY,
    registration_certificate text NOT NULL,
    driving_experience timestamptz NOT NULL,
    issue_date timestamptz NOT NULL,
    expiry_date timestamptz NOT NULL,
    country varchar(3) NOT NULL
);

CREATE TABLE drivers
(
    id uuid NOT NULL PRIMARY KEY,
    telegram_id int NOT NULL UNIQUE,
    fleet_id varchar(32) NOT NULL UNIQUE,
    jump_id int NOT NULL UNIQUE,
    first_name text NOT NULL,
    last_name text NOT NULL,
    middle_name text,
    city text NOT NULL,
    phone_number text NOT NULL,
    accept_cash bool NOT NULL,
    work_rule_id text NOT NULL,
    work_rule_updated_at timestamptz NOT NULL,
    is_self_employed bool NOT NULL,
    car_id uuid REFERENCES cars(id) NOT NULL,
    driver_license_id uuid REFERENCES driver_license(id) NOT NULL,
    created_at timestamptz NOT NULL
);


-- +goose Down
DROP TABLE drivers;
DROP TABLE driver_license;
DROP TABLE cars;