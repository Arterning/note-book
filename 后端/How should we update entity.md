
If the filed is null, we'd better keep it as same instead of set it to NULL

Because if we dont' do this way, the client must pass all the old fields

what if we want to set it to NULL?
We should set is to empty string ''


---

This things also happens in query

if the query thing is null, should we pass it into sql query ?

```sql

select * from user where name like '%abc%'; --this is what i expcted

select * from user where name like '%abc%' and email is null and pass is null;
```