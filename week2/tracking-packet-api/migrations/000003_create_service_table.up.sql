create table services(
    service_id varchar(64) not null,
    nama_service varchar(128) not null,
    harga_per_kg float(10, 2) not null, 
    primary key(service_id)
)ENGINE=InnoDB;