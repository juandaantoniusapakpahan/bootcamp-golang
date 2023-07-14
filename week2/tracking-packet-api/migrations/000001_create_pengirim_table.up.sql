create table pengirim(
    pengirim_id VARCHAR(64) NOT NULL,
    nama_pengirim VARCHAR(128) NOT NULL,
    no_telepon VARCHAR(64) NOT NULL,
    PRIMARY KEY (pengirim_id)
)ENGINE=InnoDB;