where condition should put at the end

```sql

		select
            guide.*,
            mission.farm_id as farmId,
            org.name as farmName
        from
            sys_expert_mission mission
            left join sys_expert_guide guide on guide.mission_id = mission.id
            left join sys_org org on org.id = mission.farm_id
        <where>
            and mission.is_active = 'Y'
            and guide.is_active = 'Y'
            and org.is_active = 'Y'
            and mission.user_id = #{userId}
            and mission.farm_id = #{farmId}
            <if test="currentDate != null">
                and guide.start_date &lt;= #{currentDate}
                and guide.end_date &gt;= #{currentDate}
            </if>
        </where>
```


compare with this one 


```sql

		select
            guide.*,
            mission.farm_id as farmId,
            org.name as farmName
        from
            sys_expert_mission mission
            left join sys_expert_guide guide on guide.mission_id = mission.id                                                              and guide.is_active = 'Y'
            left join sys_org org on org.id = mission.farm_id
        <where>
            and mission.is_active = 'Y'
            and org.is_active = 'Y'
            and mission.user_id = #{userId}
            and mission.farm_id = #{farmId}
            <if test="currentDate != null">
                and guide.start_date &lt;= #{currentDate}
                and guide.end_date &gt;= #{currentDate}
            </if>
        </where>
```

if you put the `guide.is_active = 'Y'` after the join condition

the records that is not active will also joined, because you used is left join 

so the not active data will returned , that is not we wanna see.

we should use inner join or put the where filter condition at the end










