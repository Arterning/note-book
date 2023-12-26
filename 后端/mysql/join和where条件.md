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

join 后面只需要跟连接条件
where后面才需要跟过滤条件
这一点需要注意。


如果在join后面跟了过滤条件，会有什么问题？
因为是left join
那么不满足join条件的记录也会返回！但是右边的连接表的记录为空
这就导致NPE问题

如果需要早join加过滤条件 最好使用inner join


对于DB数据返回空的情况要比较小心 使用stream流处理最好防御编程