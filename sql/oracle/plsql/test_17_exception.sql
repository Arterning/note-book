SET SERVEROUTPUT ON SIZE 99999;
DECLARE 
   c_id customers.id%type := 100; 
   c_name  customerS.name%type; 
   c_addr customers.address%type; 
BEGIN 
   SELECT  name, address INTO  c_name, c_addr 
   FROM customers 
   WHERE id = c_id;  
   DBMS_OUTPUT.PUT_LINE ('姓名: '||  c_name); 
   DBMS_OUTPUT.PUT_LINE ('地址: ' || c_addr); 

EXCEPTION 
   WHEN no_data_found THEN 
      dbms_output.put_line('没有找到符合条件的客户信息!'); 
   WHEN others THEN 
      dbms_output.put_line('Error!'); 
END; 
/

