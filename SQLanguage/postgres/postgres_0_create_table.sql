CREATE DATABASE ning;
create database huang;
create database finance;
create database gold;



CREATE TABLE COMPANY(
                        ID INT PRIMARY KEY     NOT NULL,
                        NAME           TEXT    NOT NULL,
                        AGE            INT     NOT NULL,
                        ADDRESS        CHAR(50),
                        SALARY         REAL
);

INSERT INTO COMPANY (ID,NAME,AGE,ADDRESS,SALARY)
VALUES (1, 'Paul', 32, 'California', 20000.00);

select * from COMPANY t for update;

CREATE TABLE DEPARTMENT(
                           ID INT PRIMARY KEY      NOT NULL,
                           DEPT           CHAR(50) NOT NULL,
                           EMP_ID         INT      NOT NULL
);


