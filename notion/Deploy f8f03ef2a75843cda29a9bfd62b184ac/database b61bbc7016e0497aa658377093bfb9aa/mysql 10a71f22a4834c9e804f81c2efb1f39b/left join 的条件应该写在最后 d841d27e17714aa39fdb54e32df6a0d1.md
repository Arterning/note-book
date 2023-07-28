# left join 的条件应该写在最后

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

对于DB数据返回空的情况要比较小心 使用stream流处理最好防御编程