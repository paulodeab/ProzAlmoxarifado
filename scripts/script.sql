create database if not exists db_proz;
use db_proz;
create table if not exists produto(id integer primary key auto_increment, 
                    nomeItem varchar(45), quantidade float,unidadeMedida varchar(10), local varchar(45));