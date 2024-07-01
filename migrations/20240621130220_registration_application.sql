-- +goose Up
CREATE TABLE registration_application (
    id int PRIMARY KEY,
    status text NOT NULL,
    date date NOT NULL,

    phone_number text,
    last_name text,
    first_name text,
    middle_name text,
    city text,

    license_number text,
    license_total_since_date date,
    license_issue_date date,
    license_expiry_date date,
    license_country varchar(3),

    car_brand text,
    car_model text,
    car_year int,
    car_color text,
    car_vin text,
    car_number text,
    car_license text,

    work_rule_id text
);

-- +goose Down
DROP TABLE registration_application;