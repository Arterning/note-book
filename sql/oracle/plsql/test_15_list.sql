SET SERVEROUTPUT ON SIZE 99999;
DECLARE 
   CURSOR c_customers is 
      select name from customers; 

   TYPE c_list IS TABLE of customers.name%type INDEX BY binary_integer; 
   name_list c_list; 
   counter integer :=0; 
BEGIN 
   FOR n IN c_customers LOOP 
      counter := counter +1; 
      name_list(counter) := n.name; 
      dbms_output.put_line('Customer('||counter||'):'||name_list(counter)); 
   END LOOP; 
END; 
/

