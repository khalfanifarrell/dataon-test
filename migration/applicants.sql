CREATE TABLE IF NOT EXISTS applicants
(
	id              INT NOT NULL AUTO_INCREMENT,
    first_name      VARCHAR(150) NOT NULL,
    last_name       VARCHAR(150) NOT NULL,
    email           VARCHAR(150) NOT NULL,
    phone           VARCHAR(150) NOT NULL,
    home_address    TEXT NOT NULL,
    title 			VARCHAR(150) NOT NULL,
    years_of_exp 	INT NOT NULL,
	created_at      TIMESTAMP NOT NULL DEFAULT current_timestamp(),
	deleted_at      TIMESTAMP,
    PRIMARY KEY id (id),
    KEY deleted_at (deleted_at)
)
ENGINE=InnoDB
DEFAULT CHARSET=latin1
COLLATE=latin1_swedish_ci;