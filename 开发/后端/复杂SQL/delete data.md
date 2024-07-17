
```sql
delete from ag_message  
where (weather_alert_code = '11B01' or weather_alert_code = '11B03')  
and (create_date > '2023-09-18 12:00:00')  
and (business_id, farm_id, create_date) NOT IN (  
SELECT business_id, farm_id, MAX(create_date)  
FROM ag_message where create_date > '2023-09-18 12:00:00' and (weather_alert_code = '11B01' or weather_alert_code = '11B03')  
GROUP BY business_id, farm_id  
);
```


Here means only remain the max created_date record
```sql
NOT IN (  
SELECT business_id, farm_id, MAX(create_date)  ..
```