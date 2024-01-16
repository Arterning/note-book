
```sql

UPDATE ag_varieties
SET rice_type = (SELECT CASE
                            WHEN r.subspecies = '籼' THEN 1
                            WHEN r.subspecies = '粳' THEN 2
                            WHEN r.subspecies = '籼粳交' THEN 3
                            ELSE NULL
                            END
                 FROM variety_recommend r
                 WHERE r.id = ag_varieties.recommend_id),
    sow_type  = (SELECT CASE
                            WHEN r.rice_type = '双晚稻' THEN 3
                            WHEN r.rice_type = '早稻' THEN 4
                            WHEN r.rice_type = '单晚稻' THEN 5
                            WHEN r.rice_type = '中稻' THEN 6
                            WHEN r.rice_type = '再生稻' THEN 8
                            ELSE 3
                            END
                 FROM variety_recommend r
                 WHERE r.id = ag_varieties.recommend_id)
WHERE recommend_id IS NOT NULL;

```