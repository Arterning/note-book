CREATE OR REPLACE FUNCTION totalRecords (arg1 integer, arg2 integer)
    RETURNS integer AS $total$
declare
    total integer;
BEGIN
   raise notice '%,yeah , print..%', arg1 , arg2;
   SELECT count(*) into total FROM COMPANY;
   RETURN total;
END;
$total$ LANGUAGE plpgsql;

--invoke
select totalRecords(12, 33);


