DECLARE 
   c_id customers.id%type := &cc_id; 
   c_name  customerS.name%type; 
   c_addr customers.address%type;  
   -- user defined exception 
   ex_invalid_id  EXCEPTION; 
BEGIN 
   IF c_id <= 0 THEN 
   -- raise user exception
      RAISE ex_invalid_id; 
   ELSE 
      SELECT  name, address INTO  c_name, c_addr 
      FROM customers 
      WHERE id = c_id;
      DBMS_OUTPUT.PUT_LINE ('姓名: '||  c_name);  
      DBMS_OUTPUT.PUT_LINE ('地址: ' || c_addr); 
   END IF; 

EXCEPTION 
   WHEN ex_invalid_id THEN 
      dbms_output.put_line('编号ID必须要大于0!'); 
   WHEN no_data_found THEN 
      dbms_output.put_line('未找到指定ID的客户信息!'); 
   WHEN others THEN 
      dbms_output.put_line('Error!');  
END; 
/

