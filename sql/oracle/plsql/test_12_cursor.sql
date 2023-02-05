-- 每当发出DML语句(INSERT，UPDATE和DELETE)时，隐式游标与此语句相关联
SET SERVEROUTPUT ON SIZE 99999;
DECLARE  
   total_rows number(2); 
BEGIN 
   UPDATE customers 
   SET salary = salary + 500; 
   IF sql%notfound THEN 
      dbms_output.put_line('没有找到客户信息~'); 
   ELSIF sql%found THEN 
      total_rows := sql%rowcount;
      dbms_output.put_line('一共有：' || total_rows || ' 个客户的工资被更新！ '); 
   END IF;  
END; 
/

