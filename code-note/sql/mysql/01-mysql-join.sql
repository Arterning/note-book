SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl a, tcount_tbl b WHERE a.runoob_author = b.runoob_author;


SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl a 
LEFT JOIN tcount_tbl b ON a.runoob_author = b.runoob_author;


SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl a 
RIGHT JOIN tcount_tbl b ON a.runoob_author = b.runoob_author;
