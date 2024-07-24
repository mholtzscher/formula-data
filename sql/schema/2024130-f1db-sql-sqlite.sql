DROP VIEW IF EXISTS pre_qualifying_result;
DROP VIEW IF EXISTS free_practice_1_result;
DROP VIEW IF EXISTS free_practice_2_result;
DROP VIEW IF EXISTS free_practice_3_result;
DROP VIEW IF EXISTS free_practice_4_result;
DROP VIEW IF EXISTS qualifying_1_result;
DROP VIEW IF EXISTS qualifying_2_result;
DROP VIEW IF EXISTS qualifying_result;
DROP VIEW IF EXISTS sprint_qualifying_result;
DROP VIEW IF EXISTS sprint_starting_grid_position;
DROP VIEW IF EXISTS sprint_race_result;
DROP VIEW IF EXISTS warming_up_result;
DROP VIEW IF EXISTS starting_grid_position;
DROP VIEW IF EXISTS race_result;
DROP VIEW IF EXISTS fastest_lap;
DROP VIEW IF EXISTS pit_stop;
DROP VIEW IF EXISTS driver_of_the_day_result;

DROP TABLE IF EXISTS race_constructor_standing;
DROP TABLE IF EXISTS race_driver_standing;
DROP TABLE IF EXISTS race_data;
DROP TABLE IF EXISTS race;
DROP TABLE IF EXISTS season_constructor_standing;
DROP TABLE IF EXISTS season_driver_standing;
DROP TABLE IF EXISTS season_entrant_driver;
DROP TABLE IF EXISTS season_entrant_tyre_manufacturer;
DROP TABLE IF EXISTS season_entrant_constructor;
DROP TABLE IF EXISTS season_entrant;
DROP TABLE IF EXISTS season;
DROP TABLE IF EXISTS grand_prix;
DROP TABLE IF EXISTS circuit;
DROP TABLE IF EXISTS entrant;
DROP TABLE IF EXISTS tyre_manufacturer;
DROP TABLE IF EXISTS engine_manufacturer;
DROP TABLE IF EXISTS constructor_previous_next_constructor;
DROP TABLE IF EXISTS constructor;
DROP TABLE IF EXISTS driver_family_relationship;
DROP TABLE IF EXISTS driver;
DROP TABLE IF EXISTS country;
DROP TABLE IF EXISTS continent;

CREATE TABLE continent (
  id varchar(100) NOT NULL,
  code varchar(2) NOT NULL,
  name varchar(100) NOT NULL,
  demonym varchar(100) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (code),
  UNIQUE (name)
);

CREATE TABLE country (
  id varchar(100) NOT NULL,
  alpha2_code varchar(2) NOT NULL,
  alpha3_code varchar(3) NOT NULL,
  name varchar(100) NOT NULL,
  demonym varchar(100),
  continent_id varchar(100) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (alpha2_code),
  UNIQUE (alpha3_code),
  UNIQUE (name),
  FOREIGN KEY (continent_id) REFERENCES continent (id)
);

CREATE INDEX coun_continent_id_idx ON country(continent_id);

CREATE TABLE driver (
  id varchar(100) NOT NULL,
  name varchar(100) NOT NULL,
  first_name varchar(100) NOT NULL,
  last_name varchar(100) NOT NULL,
  full_name varchar(100) NOT NULL,
  abbreviation varchar(3) NOT NULL,
  permanent_number varchar(2),
  gender varchar(6) NOT NULL,
  date_of_birth date NOT NULL,
  date_of_death date,
  place_of_birth varchar(100) NOT NULL,
  country_of_birth_country_id varchar(100) NOT NULL,
  nationality_country_id varchar(100) NOT NULL,
  second_nationality_country_id varchar(100),
  best_championship_position int,
  best_starting_grid_position int,
  best_race_result int,
  total_championship_wins int NOT NULL,
  total_race_entries int NOT NULL,
  total_race_starts int NOT NULL,
  total_race_wins int NOT NULL,
  total_race_laps int NOT NULL,
  total_podiums int NOT NULL,
  total_points decimal(8, 2) NOT NULL,
  total_championship_points decimal(8, 2) NOT NULL,
  total_pole_positions int NOT NULL,
  total_fastest_laps int NOT NULL,
  total_driver_of_the_day int NOT NULL,
  total_grand_slams int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (country_of_birth_country_id) REFERENCES country (id),
  FOREIGN KEY (nationality_country_id) REFERENCES country (id),
  FOREIGN KEY (second_nationality_country_id) REFERENCES country (id)
);

CREATE INDEX driv_abbreviation_idx ON driver(abbreviation);
CREATE INDEX driv_country_of_birth_country_id_idx ON driver(country_of_birth_country_id);
CREATE INDEX driv_date_of_birth_idx ON driver(date_of_birth);
CREATE INDEX driv_date_of_death_idx ON driver(date_of_death);
CREATE INDEX driv_first_name_idx ON driver(first_name);
CREATE INDEX driv_full_name_idx ON driver(full_name);
CREATE INDEX driv_gender_idx ON driver(gender);
CREATE INDEX driv_last_name_idx ON driver(last_name);
CREATE INDEX driv_name_idx ON driver(name);
CREATE INDEX driv_nationality_country_id_idx ON driver(nationality_country_id);
CREATE INDEX driv_permanent_number_idx ON driver(permanent_number);
CREATE INDEX driv_place_of_birth_idx ON driver(place_of_birth);
CREATE INDEX driv_second_nationality_country_id_idx ON driver(second_nationality_country_id);

CREATE TABLE driver_family_relationship (
  driver_id varchar(100) NOT NULL,
  other_driver_id varchar(100) NOT NULL,
  type varchar(50) NOT NULL,
  PRIMARY KEY (driver_id, other_driver_id, type),
  FOREIGN KEY (driver_id) REFERENCES driver (id),
  FOREIGN KEY (other_driver_id) REFERENCES driver (id)
);

CREATE INDEX drfr_driver_id_idx ON driver_family_relationship(driver_id);
CREATE INDEX drfr_other_driver_id_idx ON driver_family_relationship(other_driver_id);

CREATE TABLE constructor (
  id varchar(100) NOT NULL,
  name varchar(100) NOT NULL,
  full_name varchar(100) NOT NULL,
  country_id varchar(100) NOT NULL,
  best_championship_position int,
  best_starting_grid_position int,
  best_race_result int,
  total_championship_wins int NOT NULL,
  total_race_entries int NOT NULL,
  total_race_starts int NOT NULL,
  total_race_wins int NOT NULL,
  total_1_and_2_finishes int NOT NULL,
  total_race_laps int NOT NULL,
  total_podiums int NOT NULL,
  total_podium_races int NOT NULL,
  total_championship_points decimal(8, 2) NOT NULL,
  total_pole_positions int NOT NULL,
  total_fastest_laps int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (country_id) REFERENCES country (id)
);

CREATE INDEX cons_country_id_idx ON constructor(country_id);
CREATE INDEX cons_full_name_idx ON constructor(full_name);
CREATE INDEX cons_name_idx ON constructor(name);

CREATE TABLE constructor_previous_next_constructor (
  constructor_id varchar(100) NOT NULL,
  previous_next_constructor_id varchar(100) NOT NULL,
  year_from int NOT NULL,
  year_to int,
  PRIMARY KEY (constructor_id, previous_next_constructor_id, year_from),
  FOREIGN KEY (constructor_id) REFERENCES constructor (id),
  FOREIGN KEY (previous_next_constructor_id) REFERENCES constructor (id)
);

CREATE INDEX cpnc_constructor_id_idx ON constructor_previous_next_constructor(constructor_id);
CREATE INDEX cpnc_previous_next_constructor_id_idx ON constructor_previous_next_constructor(previous_next_constructor_id);

CREATE TABLE engine_manufacturer (
  id varchar(100) NOT NULL,
  name varchar(100) NOT NULL,
  country_id varchar(100) NOT NULL,
  best_championship_position int,
  best_starting_grid_position int,
  best_race_result int,
  total_championship_wins int NOT NULL,
  total_race_entries int NOT NULL,
  total_race_starts int NOT NULL,
  total_race_wins int NOT NULL,
  total_race_laps int NOT NULL,
  total_podiums int NOT NULL,
  total_podium_races int NOT NULL,
  total_championship_points decimal(8, 2) NOT NULL,
  total_pole_positions int NOT NULL,
  total_fastest_laps int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (country_id) REFERENCES country (id)
);

CREATE INDEX enma_country_id_idx ON engine_manufacturer(country_id);
CREATE INDEX enma_name_idx ON engine_manufacturer(name);

CREATE TABLE tyre_manufacturer (
  id varchar(100) NOT NULL,
  name varchar(100) NOT NULL,
  country_id varchar(100) NOT NULL,
  best_starting_grid_position int,
  best_race_result int,
  total_race_entries int NOT NULL,
  total_race_starts int NOT NULL,
  total_race_wins int NOT NULL,
  total_race_laps int NOT NULL,
  total_podiums int NOT NULL,
  total_podium_races int NOT NULL,
  total_pole_positions int NOT NULL,
  total_fastest_laps int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (country_id) REFERENCES country (id)
);

CREATE INDEX tyma_country_id_idx ON tyre_manufacturer(country_id);
CREATE INDEX tyma_name_idx ON tyre_manufacturer(name);

CREATE TABLE entrant (
  id varchar(100) NOT NULL,
  name varchar(100) NOT NULL,
  PRIMARY KEY (id)
);

CREATE INDEX entr_name_idx ON entrant(name);

CREATE TABLE circuit (
  id varchar(100) NOT NULL,
  name varchar(100) NOT NULL,
  full_name varchar(100) NOT NULL,
  previous_names varchar(255),
  type varchar(6) NOT NULL,
  place_name varchar(100) NOT NULL,
  country_id varchar(100) NOT NULL,
  latitude decimal(10, 6) NOT NULL,
  longitude decimal(10, 6) NOT NULL,
  total_races_held int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (country_id) REFERENCES country (id)
);

CREATE INDEX circ_country_id_idx ON circuit(country_id);
CREATE INDEX circ_full_name_idx ON circuit(full_name);
CREATE INDEX circ_name_idx ON circuit(name);
CREATE INDEX circ_place_name_idx ON circuit(place_name);
CREATE INDEX circ_type_idx ON circuit(type);

CREATE TABLE grand_prix (
  id varchar(100) NOT NULL,
  name varchar(100) NOT NULL,
  full_name varchar(100) NOT NULL,
  short_name varchar(100) NOT NULL,
  abbreviation varchar(3) NOT NULL,
  country_id varchar(100),
  total_races_held int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (country_id) REFERENCES country (id)
);

CREATE INDEX grpr_abbreviation_idx ON grand_prix(abbreviation);
CREATE INDEX grpr_country_id_idx ON grand_prix(country_id);
CREATE INDEX grpr_full_name_idx ON grand_prix(full_name);
CREATE INDEX grpr_name_idx ON grand_prix(name);
CREATE INDEX grpr_short_name_idx ON grand_prix(short_name);

CREATE TABLE season (
  year int NOT NULL,
  PRIMARY KEY (year)
);

CREATE TABLE season_entrant (
  year int NOT NULL,
  entrant_id varchar(100) NOT NULL,
  country_id varchar(100) NOT NULL,
  PRIMARY KEY (year, entrant_id),
  FOREIGN KEY (year) REFERENCES season (year),
  FOREIGN KEY (entrant_id) REFERENCES entrant (id),
  FOREIGN KEY (country_id) REFERENCES country (id)
);

CREATE INDEX seen_country_id_idx ON season_entrant(country_id);
CREATE INDEX seen_entrant_id_idx ON season_entrant(entrant_id);
CREATE INDEX seen_year_idx ON season_entrant(year);

CREATE TABLE season_entrant_constructor (
  year int NOT NULL,
  entrant_id varchar(100) NOT NULL,
  constructor_id varchar(100) NOT NULL,
  engine_manufacturer_id varchar(100) NOT NULL,
  PRIMARY KEY (year, entrant_id, constructor_id, engine_manufacturer_id),
  FOREIGN KEY (year) REFERENCES season (year),
  FOREIGN KEY (entrant_id) REFERENCES entrant (id),
  FOREIGN KEY (constructor_id) REFERENCES constructor (id),
  FOREIGN KEY (engine_manufacturer_id) REFERENCES engine_manufacturer (id)
);

CREATE INDEX seec_constructor_id_idx ON season_entrant_constructor(constructor_id);
CREATE INDEX seec_engine_manufacturer_id_idx ON season_entrant_constructor(engine_manufacturer_id);
CREATE INDEX seec_entrant_id_idx ON season_entrant_constructor(entrant_id);
CREATE INDEX seec_year_idx ON season_entrant_constructor(year);

CREATE TABLE season_entrant_tyre_manufacturer (
  year int NOT NULL,
  entrant_id varchar(100) NOT NULL,
  constructor_id varchar(100) NOT NULL,
  engine_manufacturer_id varchar(100) NOT NULL,
  tyre_manufacturer_id varchar(100) NOT NULL,
  PRIMARY KEY (
    year,
    entrant_id,
    constructor_id,
    engine_manufacturer_id,
    tyre_manufacturer_id
  ),
  FOREIGN KEY (year) REFERENCES season (year),
  FOREIGN KEY (entrant_id) REFERENCES entrant (id),
  FOREIGN KEY (constructor_id) REFERENCES constructor (id),
  FOREIGN KEY (engine_manufacturer_id) REFERENCES engine_manufacturer (id),
  FOREIGN KEY (tyre_manufacturer_id) REFERENCES tyre_manufacturer (id)
);

CREATE INDEX setm_constructor_id_idx ON season_entrant_tyre_manufacturer(constructor_id);
CREATE INDEX setm_engine_manufacturer_id_idx ON season_entrant_tyre_manufacturer(engine_manufacturer_id);
CREATE INDEX setm_entrant_id_idx ON season_entrant_tyre_manufacturer(entrant_id);
CREATE INDEX setm_tyre_manufacturer_id_idx ON season_entrant_tyre_manufacturer(tyre_manufacturer_id);
CREATE INDEX setm_year_idx ON season_entrant_tyre_manufacturer(year);

CREATE TABLE season_entrant_driver (
  year int NOT NULL,
  entrant_id varchar(100) NOT NULL,
  constructor_id varchar(100) NOT NULL,
  engine_manufacturer_id varchar(100) NOT NULL,
  driver_id varchar(100) NOT NULL,
  rounds varchar(100),
  rounds_text varchar(100),
  test_driver boolean NOT NULL,
  PRIMARY KEY (
    year,
    entrant_id,
    constructor_id,
    engine_manufacturer_id,
    driver_id
  ),
  FOREIGN KEY (year) REFERENCES season (year),
  FOREIGN KEY (entrant_id) REFERENCES entrant (id),
  FOREIGN KEY (constructor_id) REFERENCES constructor (id),
  FOREIGN KEY (engine_manufacturer_id) REFERENCES engine_manufacturer (id),
  FOREIGN KEY (driver_id) REFERENCES driver (id)
);

CREATE INDEX seed_constructor_id_idx ON season_entrant_driver(constructor_id);
CREATE INDEX seed_driver_id_idx ON season_entrant_driver(driver_id);
CREATE INDEX seed_engine_manufacturer_id_idx ON season_entrant_driver(engine_manufacturer_id);
CREATE INDEX seed_entrant_id_idx ON season_entrant_driver(entrant_id);
CREATE INDEX seed_year_idx ON season_entrant_driver(year);

CREATE TABLE season_driver_standing (
  year int NOT NULL,
  position_display_order int NOT NULL,
  position_number int,
  position_text varchar(4) NOT NULL,
  driver_id varchar(100) NOT NULL,
  points decimal(8, 2) NOT NULL,
  PRIMARY KEY (year, position_display_order),
  FOREIGN KEY (year) REFERENCES season (year),
  FOREIGN KEY (driver_id) REFERENCES driver (id)
);

CREATE INDEX seds_driver_id_idx ON season_driver_standing(driver_id);
CREATE INDEX seds_position_display_order_idx ON season_driver_standing(position_display_order);
CREATE INDEX seds_position_number_idx ON season_driver_standing(position_number);
CREATE INDEX seds_position_text_idx ON season_driver_standing(position_text);
CREATE INDEX seds_year_idx ON season_driver_standing(year);

CREATE TABLE season_constructor_standing (
  year int NOT NULL,
  position_display_order int NOT NULL,
  position_number int,
  position_text varchar(4) NOT NULL,
  constructor_id varchar(100) NOT NULL,
  engine_manufacturer_id varchar(100) NOT NULL,
  points decimal(8, 2) NOT NULL,
  PRIMARY KEY (year, position_display_order),
  FOREIGN KEY (year) REFERENCES season (year),
  FOREIGN KEY (constructor_id) REFERENCES constructor (id),
  FOREIGN KEY (engine_manufacturer_id) REFERENCES engine_manufacturer (id)
);

CREATE INDEX secs_constructor_id_idx ON season_constructor_standing(constructor_id);
CREATE INDEX secs_engine_manufacturer_id_idx ON season_constructor_standing(engine_manufacturer_id);
CREATE INDEX secs_position_display_order_idx ON season_constructor_standing(position_display_order);
CREATE INDEX secs_position_number_idx ON season_constructor_standing(position_number);
CREATE INDEX secs_position_text_idx ON season_constructor_standing(position_text);
CREATE INDEX secs_year_idx ON season_constructor_standing(year);

CREATE TABLE race (
  id int NOT NULL,
  year int NOT NULL,
  round int NOT NULL,
  date date NOT NULL,
  time clob,
  grand_prix_id varchar(100) NOT NULL,
  official_name varchar(100) NOT NULL,
  qualifying_format varchar(20) NOT NULL,
  sprint_qualifying_format varchar(20),
  circuit_id varchar(100) NOT NULL,
  circuit_type varchar(6) NOT NULL,
  course_length decimal(6, 3) NOT NULL,
  laps int NOT NULL,
  distance decimal(6, 3) NOT NULL,
  scheduled_laps int,
  scheduled_distance decimal(6, 3),
  pre_qualifying_date date,
  pre_qualifying_time varchar(5),
  free_practice_1_date date,
  free_practice_1_time varchar(5),
  free_practice_2_date date,
  free_practice_2_time varchar(5),
  free_practice_3_date date,
  free_practice_3_time varchar(5),
  free_practice_4_date date,
  free_practice_4_time varchar(5),
  qualifying_1_date date,
  qualifying_1_time varchar(5),
  qualifying_2_date date,
  qualifying_2_time varchar(5),
  qualifying_date date,
  qualifying_time varchar(5),
  sprint_qualifying_date date,
  sprint_qualifying_time varchar(5),
  sprint_race_date date,
  sprint_race_time varchar(5),
  warming_up_date date,
  warming_up_time varchar(5),
  PRIMARY KEY (id),
  UNIQUE (year, round),
  FOREIGN KEY (year) REFERENCES season (year),
  FOREIGN KEY (grand_prix_id) REFERENCES grand_prix (id),
  FOREIGN KEY (circuit_id) REFERENCES circuit (id)
);

CREATE INDEX race_circuit_id_idx ON race(circuit_id);
CREATE INDEX race_date_idx ON race(date);
CREATE INDEX race_grand_prix_id_idx ON race(grand_prix_id);
CREATE INDEX race_official_name_idx ON race(official_name);
CREATE INDEX race_round_idx ON race(round);
CREATE INDEX race_year_idx ON race(year);

CREATE TABLE race_data (
  race_id int NOT NULL,
  type varchar(50) NOT NULL,
  position_display_order int NOT NULL,
  position_number int,
  position_text varchar(4) NOT NULL,
  driver_number varchar(3) NOT NULL,
  driver_id varchar(100) NOT NULL,
  constructor_id varchar(100) NOT NULL,
  engine_manufacturer_id varchar(100) NOT NULL,
  tyre_manufacturer_id varchar(100) NOT NULL,
  practice_time varchar(20),
  practice_time_millis int,
  practice_gap varchar(20),
  practice_gap_millis int,
  practice_interval varchar(20),
  practice_interval_millis int,
  practice_laps int,
  qualifying_time varchar(20),
  qualifying_time_millis int,
  qualifying_q1 varchar(20),
  qualifying_q1_millis int,
  qualifying_q2 varchar(20),
  qualifying_q2_millis int,
  qualifying_q3 varchar(20),
  qualifying_q3_millis int,
  qualifying_gap varchar(20),
  qualifying_gap_millis int,
  qualifying_interval varchar(20),
  qualifying_interval_millis int,
  qualifying_laps int,
  starting_grid_position_grid_penalty varchar(20),
  starting_grid_position_grid_penalty_positions int,
  starting_grid_position_time varchar(20),
  starting_grid_position_time_millis int,
  race_shared_car boolean,
  race_laps int,
  race_time varchar(20),
  race_time_millis int,
  race_time_penalty varchar(20),
  race_time_penalty_millis int,
  race_gap varchar(20),
  race_gap_millis int,
  race_gap_laps int,
  race_interval varchar(20),
  race_interval_millis int,
  race_reason_retired varchar(100),
  race_points decimal(8, 2),
  race_grid_position_number int,
  race_grid_position_text varchar(2),
  race_positions_gained int,
  race_pit_stops int,
  race_fastest_lap boolean,
  race_driver_of_the_day boolean,
  race_grand_slam boolean,
  fastest_lap_lap int,
  fastest_lap_time varchar(20),
  fastest_lap_time_millis int,
  fastest_lap_gap varchar(20),
  fastest_lap_gap_millis int,
  fastest_lap_interval varchar(20),
  fastest_lap_interval_millis int,
  pit_stop_stop int,
  pit_stop_lap int,
  pit_stop_time varchar(20),
  pit_stop_time_millis int,
  driver_of_the_day_percentage decimal(4, 1),
  PRIMARY KEY (race_id, type, position_display_order),
  FOREIGN KEY (race_id) REFERENCES race (id),
  FOREIGN KEY (driver_id) REFERENCES driver (id),
  FOREIGN KEY (constructor_id) REFERENCES constructor (id),
  FOREIGN KEY (engine_manufacturer_id) REFERENCES engine_manufacturer (id),
  FOREIGN KEY (tyre_manufacturer_id) REFERENCES tyre_manufacturer (id)
);

CREATE INDEX rada_constructor_id_idx ON race_data(constructor_id);
CREATE INDEX rada_driver_id_idx ON race_data(driver_id);
CREATE INDEX rada_driver_number_idx ON race_data(driver_number);
CREATE INDEX rada_engine_manufacturer_id_idx ON race_data(engine_manufacturer_id);
CREATE INDEX rada_position_number_idx ON race_data(position_number);
CREATE INDEX rada_position_text_idx ON race_data(position_text);
CREATE INDEX rada_race_id_idx ON race_data(race_id);
CREATE INDEX rada_type_idx ON race_data(type);
CREATE INDEX rada_tyre_manufacturer_id_idx ON race_data(tyre_manufacturer_id);

CREATE TABLE race_driver_standing (
  race_id int NOT NULL,
  position_display_order int NOT NULL,
  position_number int,
  position_text varchar(4) NOT NULL,
  driver_id varchar(100) NOT NULL,
  points decimal(8, 2) NOT NULL,
  positions_gained int,
  PRIMARY KEY (race_id, position_display_order),
  FOREIGN KEY (race_id) REFERENCES race (id),
  FOREIGN KEY (driver_id) REFERENCES driver (id)
);

CREATE INDEX rads_driver_id_idx ON race_driver_standing(driver_id);
CREATE INDEX rads_position_display_order_idx ON race_driver_standing(position_display_order);
CREATE INDEX rads_position_number_idx ON race_driver_standing(position_number);
CREATE INDEX rads_position_text_idx ON race_driver_standing(position_text);
CREATE INDEX rads_race_id_idx ON race_driver_standing(race_id);

CREATE TABLE race_constructor_standing (
  race_id int NOT NULL,
  position_display_order int NOT NULL,
  position_number int,
  position_text varchar(4) NOT NULL,
  constructor_id varchar(100) NOT NULL,
  engine_manufacturer_id varchar(100) NOT NULL,
  points decimal(8, 2) NOT NULL,
  positions_gained int,
  PRIMARY KEY (race_id, position_display_order),
  FOREIGN KEY (race_id) REFERENCES race (id),
  FOREIGN KEY (constructor_id) REFERENCES constructor (id),
  FOREIGN KEY (engine_manufacturer_id) REFERENCES engine_manufacturer (id)
);

CREATE INDEX racs_constructor_id_idx ON race_constructor_standing(constructor_id);
CREATE INDEX racs_engine_manufacturer_id_idx ON race_constructor_standing(engine_manufacturer_id);
CREATE INDEX racs_position_display_order_idx ON race_constructor_standing(position_display_order);
CREATE INDEX racs_position_number_idx ON race_constructor_standing(position_number);
CREATE INDEX racs_position_text_idx ON race_constructor_standing(position_text);
CREATE INDEX racs_race_id_idx ON race_constructor_standing(race_id);

CREATE VIEW pre_qualifying_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  time,
  time_millis,
  gap,
  gap_millis,
  interval,
  interval_millis,
  laps
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.qualifying_time AS time,
  race_data.qualifying_time_millis AS time_millis,
  race_data.qualifying_gap AS gap,
  race_data.qualifying_gap_millis AS gap_millis,
  race_data.qualifying_interval AS interval,
  race_data.qualifying_interval_millis AS interval_millis,
  race_data.qualifying_laps AS laps
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'PRE_QUALIFYING_RESULT'
  AND race_data.race_id = race.id
);

CREATE VIEW free_practice_1_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  time,
  time_millis,
  gap,
  gap_millis,
  interval,
  interval_millis,
  laps
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.practice_time AS time,
  race_data.practice_time_millis AS time_millis,
  race_data.practice_gap AS gap,
  race_data.practice_gap_millis AS gap_millis,
  race_data.practice_interval AS interval,
  race_data.practice_interval_millis AS interval_millis,
  race_data.practice_laps AS laps
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'FREE_PRACTICE_1_RESULT'
  AND race_data.race_id = race.id
);

CREATE VIEW free_practice_2_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  time,
  time_millis,
  gap,
  gap_millis,
  interval,
  interval_millis,
  laps
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.practice_time AS time,
  race_data.practice_time_millis AS time_millis,
  race_data.practice_gap AS gap,
  race_data.practice_gap_millis AS gap_millis,
  race_data.practice_interval AS interval,
  race_data.practice_interval_millis AS interval_millis,
  race_data.practice_laps AS laps
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'FREE_PRACTICE_2_RESULT'
  AND race_data.race_id = race.id
);

CREATE VIEW free_practice_3_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  time,
  time_millis,
  gap,
  gap_millis,
  interval,
  interval_millis,
  laps
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.practice_time AS time,
  race_data.practice_time_millis AS time_millis,
  race_data.practice_gap AS gap,
  race_data.practice_gap_millis AS gap_millis,
  race_data.practice_interval AS interval,
  race_data.practice_interval_millis AS interval_millis,
  race_data.practice_laps AS laps
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'FREE_PRACTICE_3_RESULT'
  AND race_data.race_id = race.id
);

CREATE VIEW free_practice_4_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  time,
  time_millis,
  gap,
  gap_millis,
  interval,
  interval_millis,
  laps
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.practice_time AS time,
  race_data.practice_time_millis AS time_millis,
  race_data.practice_gap AS gap,
  race_data.practice_gap_millis AS gap_millis,
  race_data.practice_interval AS interval,
  race_data.practice_interval_millis AS interval_millis,
  race_data.practice_laps AS laps
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'FREE_PRACTICE_4_RESULT'
  AND race_data.race_id = race.id
);

CREATE VIEW qualifying_1_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  time,
  time_millis,
  gap,
  gap_millis,
  interval,
  interval_millis,
  laps
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.qualifying_time AS time,
  race_data.qualifying_time_millis AS time_millis,
  race_data.qualifying_gap AS gap,
  race_data.qualifying_gap_millis AS gap_millis,
  race_data.qualifying_interval AS interval,
  race_data.qualifying_interval_millis AS interval_millis,
  race_data.qualifying_laps AS laps
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'QUALIFYING_1_RESULT'
  AND race_data.race_id = race.id
);

CREATE VIEW qualifying_2_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  time,
  time_millis,
  gap,
  gap_millis,
  interval,
  interval_millis,
  laps
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.qualifying_time AS time,
  race_data.qualifying_time_millis AS time_millis,
  race_data.qualifying_gap AS gap,
  race_data.qualifying_gap_millis AS gap_millis,
  race_data.qualifying_interval AS interval,
  race_data.qualifying_interval_millis AS interval_millis,
  race_data.qualifying_laps AS laps
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'QUALIFYING_2_RESULT'
  AND race_data.race_id = race.id
);

CREATE VIEW qualifying_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  time,
  time_millis,
  q1,
  q1_millis,
  q2,
  q2_millis,
  q3,
  q3_millis,
  gap,
  gap_millis,
  interval,
  interval_millis,
  laps
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.qualifying_time AS time,
  race_data.qualifying_time_millis AS time_millis,
  race_data.qualifying_q1 AS q1,
  race_data.qualifying_q1_millis AS q1_millis,
  race_data.qualifying_q2 AS q2,
  race_data.qualifying_q2_millis AS q2_millis,
  race_data.qualifying_q3 AS q3,
  race_data.qualifying_q3_millis AS q3_millis,
  race_data.qualifying_gap AS gap,
  race_data.qualifying_gap_millis AS gap_millis,
  race_data.qualifying_interval AS interval,
  race_data.qualifying_interval_millis AS interval_millis,
  race_data.qualifying_laps AS laps
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'QUALIFYING_RESULT'
  AND race_data.race_id = race.id
);

CREATE VIEW sprint_qualifying_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  time,
  time_millis,
  q1,
  q1_millis,
  q2,
  q2_millis,
  q3,
  q3_millis,
  gap,
  gap_millis,
  interval,
  interval_millis,
  laps
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.qualifying_time AS time,
  race_data.qualifying_time_millis AS time_millis,
  race_data.qualifying_q1 AS q1,
  race_data.qualifying_q1_millis AS q1_millis,
  race_data.qualifying_q2 AS q2,
  race_data.qualifying_q2_millis AS q2_millis,
  race_data.qualifying_q3 AS q3,
  race_data.qualifying_q3_millis AS q3_millis,
  race_data.qualifying_gap AS gap,
  race_data.qualifying_gap_millis AS gap_millis,
  race_data.qualifying_interval AS interval,
  race_data.qualifying_interval_millis AS interval_millis,
  race_data.qualifying_laps AS laps
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'SPRINT_QUALIFYING_RESULT'
  AND race_data.race_id = race.id
);

CREATE VIEW sprint_starting_grid_position(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  grid_penalty,
  grid_penalty_positions,
  time,
  time_millis
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.starting_grid_position_grid_penalty AS grid_penalty,
  race_data.starting_grid_position_grid_penalty_positions AS grid_penalty_positions,
  race_data.starting_grid_position_time AS time,
  race_data.starting_grid_position_time_millis AS time_millis
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'SPRINT_STARTING_GRID_POSITION'
  AND race_data.race_id = race.id
);

CREATE VIEW sprint_race_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  laps,
  time,
  time_millis,
  time_penalty,
  time_penalty_millis,
  gap,
  gap_millis,
  gap_laps,
  interval,
  interval_millis,
  reason_retired,
  points,
  grid_position_number,
  grid_position_text,
  positions_gained
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.race_laps AS laps,
  race_data.race_time AS time,
  race_data.race_time_millis AS time_millis,
  race_data.race_time_penalty AS time_penalty,
  race_data.race_time_penalty_millis AS time_penalty_millis,
  race_data.race_gap AS gap,
  race_data.race_gap_millis AS gap_millis,
  race_data.race_gap_laps AS gap_laps,
  race_data.race_interval AS interval,
  race_data.race_interval_millis AS interval_millis,
  race_data.race_reason_retired AS reason_retired,
  race_data.race_points AS points,
  race_data.race_grid_position_number AS grid_position_number,
  race_data.race_grid_position_text AS grid_position_text,
  race_data.race_positions_gained AS positions_gained
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'SPRINT_RACE_RESULT'
  AND race_data.race_id = race.id
);

CREATE VIEW warming_up_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  time,
  time_millis,
  gap,
  gap_millis,
  interval,
  interval_millis,
  laps
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.practice_time AS time,
  race_data.practice_time_millis AS time_millis,
  race_data.practice_gap AS gap,
  race_data.practice_gap_millis AS gap_millis,
  race_data.practice_interval AS interval,
  race_data.practice_interval_millis AS interval_millis,
  race_data.practice_laps AS laps
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'WARMING_UP_RESULT'
  AND race_data.race_id = race.id
);

CREATE VIEW starting_grid_position(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  grid_penalty,
  grid_penalty_positions,
  time,
  time_millis
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.starting_grid_position_grid_penalty AS grid_penalty,
  race_data.starting_grid_position_grid_penalty_positions AS grid_penalty_positions,
  race_data.starting_grid_position_time AS time,
  race_data.starting_grid_position_time_millis AS time_millis
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'STARTING_GRID_POSITION'
  AND race_data.race_id = race.id
);

CREATE VIEW race_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  shared_car,
  laps,
  time,
  time_millis,
  time_penalty,
  time_penalty_millis,
  gap,
  gap_millis,
  gap_laps,
  interval,
  interval_millis,
  reason_retired,
  points,
  grid_position_number,
  grid_position_text,
  positions_gained,
  pit_stops,
  fastest_lap,
  driver_of_the_day,
  grand_slam
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.race_shared_car AS shared_car,
  race_data.race_laps AS laps,
  race_data.race_time AS time,
  race_data.race_time_millis AS time_millis,
  race_data.race_time_penalty AS time_penalty,
  race_data.race_time_penalty_millis AS time_penalty_millis,
  race_data.race_gap AS gap,
  race_data.race_gap_millis AS gap_millis,
  race_data.race_gap_laps AS gap_laps,
  race_data.race_interval AS interval,
  race_data.race_interval_millis AS interval_millis,
  race_data.race_reason_retired AS reason_retired,
  race_data.race_points AS points,
  race_data.race_grid_position_number AS grid_position_number,
  race_data.race_grid_position_text AS grid_position_text,
  race_data.race_positions_gained AS positions_gained,
  race_data.race_pit_stops AS pit_stops,
  race_data.race_fastest_lap AS fastest_lap,
  race_data.race_driver_of_the_day AS driver_of_the_day,
  race_data.race_grand_slam AS grand_slam
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'RACE_RESULT'
  AND race_data.race_id = race.id
);

CREATE VIEW fastest_lap(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  lap,
  time,
  time_millis,
  gap,
  gap_millis,
  interval,
  interval_millis
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.fastest_lap_lap AS lap,
  race_data.fastest_lap_time AS time,
  race_data.fastest_lap_time_millis AS time_millis,
  race_data.fastest_lap_gap AS gap,
  race_data.fastest_lap_gap_millis AS gap_millis,
  race_data.fastest_lap_interval AS interval,
  race_data.fastest_lap_interval_millis AS interval_millis
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'FASTEST_LAP'
  AND race_data.race_id = race.id
);

CREATE VIEW pit_stop(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  stop,
  lap,
  time,
  time_millis
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.pit_stop_stop AS stop,
  race_data.pit_stop_lap AS lap,
  race_data.pit_stop_time AS time,
  race_data.pit_stop_time_millis AS time_millis
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'PIT_STOP'
  AND race_data.race_id = race.id
);

CREATE VIEW driver_of_the_day_result(
  race_id,
  year,
  round,
  position_display_order,
  position_number,
  position_text,
  driver_number,
  driver_id,
  constructor_id,
  engine_manufacturer_id,
  tyre_manufacturer_id,
  percentage
)
AS
SELECT
  race.id AS race_id,
  race.year,
  race.round,
  race_data.position_display_order,
  race_data.position_number,
  race_data.position_text,
  race_data.driver_number,
  race_data.driver_id,
  race_data.constructor_id,
  race_data.engine_manufacturer_id,
  race_data.tyre_manufacturer_id,
  race_data.driver_of_the_day_percentage AS percentage
FROM race_data
  JOIN race
    ON 1 = 1
WHERE (
  race_data.type = 'DRIVER_OF_THE_DAY_RESULT'
  AND race_data.race_id = race.id
);







































