-- Table: admin

-- DROP TABLE admin;
DROP TABLE IF EXISTS
  admin ;

CREATE TABLE admin
(
    id serial NOT NULL,
    full_name varchar(255) NOT NULL, 
    admin_email varchar(255), 
    userName varchar(255) NOT NULL,
    pass varchar(255) NOT NULL, 
    CONSTRAINT pk_admin PRIMARY KEY(id )
);

DROP TABLE IF EXISTS
  barang;

CREATE TABLE barang
(
    id_barang serial NOT NULL, 
    nama varchar(255) NOT NULL, 
    stok int NOT NULL, 
    harga float NOT NULL, 
    persen_laba float NOT NULL, 
    diskon float,
    --id_jenis_brg int NOT NULL, 
    --id_supplier int NOT NULL, 
    PRIMARY KEY (id_barang),
    --FOREIGN KEY (id_jenis_brg) REFERENCES jenis_barang(id_jenis_brg),
   	--FOREIGN KEY (id_supplier) REFERENCES supplier(id_supplier)
);

DROP TABLE jenis_barang CASCADE;

CREATE TABLE jenis_barang
(
    id_jenis_brg serial NOT NULL,
    nama_jenis varchar(255) NOT NULL, 
    PRIMARY KEY(id_jenis_brg)
);

DROP TABLE supplier CASCADE;

CREATE TABLE supplier(
    id_supplier serial NOT NULL, 
    nama varchar(255),
    alamat varchar(255),
    telephon varchar(255),
    email varchar(255),
    PRIMARY KEY(id_supplier)
);

DROP TABLE member;

CREATE TABLE member(
    id_member serial NOT NULL, 
    nama varchar(255),
    alamat varchar(255),
    telepon varchar(255),
	PRIMARY KEY(id_member)
)

