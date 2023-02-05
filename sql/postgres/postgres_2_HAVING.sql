-- HAVING 子句可以让我们筛选分组后的各组数据
-- HAVING 子句必须放置于 GROUP BY 子句后面
select * from company;
INSERT INTO COMPANY VALUES (8, 'Paul', 24, 'Houston', 20000.00);
INSERT INTO COMPANY VALUES (9, 'James', 44, 'Norway', 5000.00);
INSERT INTO COMPANY VALUES (10, 'James', 45, 'Texas', 5000.00);
SELECT NAME FROM COMPANY GROUP BY name HAVING count(name) < 2;
SELECT NAME FROM COMPANY GROUP BY name HAVING count(name) >= 2;


--DISTINCT
--Filter the duplicate name
SELECT DISTINCT name FROM COMPANY;