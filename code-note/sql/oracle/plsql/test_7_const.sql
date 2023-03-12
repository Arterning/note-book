PI CONSTANT NUMBER := 3.141592654; 
DECLARE 
   -- constant declaration 
   pi constant number := 3.141592654; 
   -- other declarations 
   radius number(5,2);  
   dia number(5,2);  
   circumference number(7, 2); 
   area number (10, 2); 
BEGIN  
   -- processing 
   radius := 9.5;  
   dia := radius * 2;  
   circumference := 2.0 * pi * radius; 
   area := pi * radius * radius; 
   -- output 
   dbms_output.put_line('半径: ' || radius); 
   dbms_output.put_line('直径: ' || dia); 
   dbms_output.put_line('圆周: ' || circumference); 
   dbms_output.put_line('面积: ' || area); 
END; 
/

