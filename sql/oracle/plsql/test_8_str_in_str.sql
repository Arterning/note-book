-- 要在字符串文字中嵌入单引号，请将两个单引号放在一起
DECLARE 
   message  varchar2(30):= 'What''s yiibai.com!'; 
BEGIN 
   dbms_output.put_line(message); 
END; 
/

-- =检查两个操作数的值是否相等，如果是，则条件成立。
-- :=就是赋值