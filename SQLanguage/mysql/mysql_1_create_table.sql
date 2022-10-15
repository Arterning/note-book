CREATE TABLE IF NOT EXISTS `runoob_tbl`(
                                           `runoob_id` INT UNSIGNED AUTO_INCREMENT,
                                           `runoob_title` VARCHAR(100) NOT NULL,
                                           `runoob_author` VARCHAR(40) NOT NULL,
                                           `submission_date` DATE,
                                           PRIMARY KEY ( `runoob_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO runoob_tbl
(runoob_title, runoob_author, submission_date)
VALUES ("学习 PHP", "菜鸟教程", NOW());


select * from runoob_tbl;


-- enum
use test;
create table enum_test ( ee enum('fish', 'apple', 'bird') not null);

select *
from test.enum_test order by ee;

insert into test.enum_test(ee) values ('aa');
insert into test.enum_test(ee) values ('apple');




-- address to number
select inet_aton('255.255.255.255');

-- number to address
select inet_ntoa(3232235777);