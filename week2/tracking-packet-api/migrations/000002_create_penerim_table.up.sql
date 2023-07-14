create table penerima (
    penerima_id varchar(64) not null,
    nama_penerima varchar(128) not null,
    no_telepon varchar(64) not null,
    PRIMARY KEY(penerima_id)
)ENGINE=InnoDB;