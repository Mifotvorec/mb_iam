// CREATE database
CREATE DATABASE mbclients
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LOCALE_PROVIDER = 'libc'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

 add UUID module

CREATE  TABLE "public".clients_main ( 
	login                varchar(20)  NOT NULL  ,
	guid                 uuid    ,
	name                 varchar(100)    ,
	middlename           varchar(100)    ,
	surname              varchar(100)    ,
	email                varchar(100)    ,
	phone                varchar(10)    ,
	datecreate           timestamp    ,
	dateupdate           timestamp    ,
	datedelete           timestamp    ,
	CONSTRAINT pk_clients_main PRIMARY KEY ( login )
 );
CREATE  TABLE "public".logs ( 
	"date"               timestamp    ,
	uuid                 uuid    ,
	text                 text    
 );
