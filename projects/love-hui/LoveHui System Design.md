
## Score Manage



the db data is 
```js
{
	"date":"2015-01-01",
	"score": 90,
	"parterId": 1
}
```

the highchart should accept data like this, for example if parterId is 1 or 2

```js
{
   "1": [
	   [{date}, {score}],[..]]
   ],
   "2": [
	   [..],[..],[..]
   ]
   
}
```


```js
type GroupedData = {
  [parterId: string]: ScoreHistory[];
};


async findChartData() {
    const data = await this.findAll();
    const group: GroupedData = this.groupData(data);
    const chartData = {};
    
    Object.entries(group).forEach(([parterId, items]) => {
      const list: DataPoint[] = items.map(({ date, score }) => [
        new Date(date).getTime(),
        score,
      ]);
      console.log(list);
      chartData[parterId] = list;
    });
    console.log(chartData);
    return chartData;
  }

  /***
  group all data by parterId
  **/
  groupData(items: ScoreHistory[]) {
    return items.reduce((result, item) => {
      const { partnerId } = item;
      if (!result[partnerId]) {
        result[partnerId] = [];
      }
      result[partnerId].push(item);
      return result;
    }, {});
  }
```


```js
import {
  Column,
  CreateDateColumn,
  Entity,
  PrimaryGeneratedColumn,
  UpdateDateColumn,
} from 'typeorm';

@Entity('score_history')
export class ScoreHistory {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  partnerId: string;

  @Column()
  score: number;

  @Column({
    type: 'date',
    nullable: false,
  })
  date: Date;

  @CreateDateColumn({
    type: 'datetime',
    comment: '创建时间',
    name: 'create_at',
  })
  createAt: Date;

  @UpdateDateColumn({
    type: 'datetime',
    comment: '更新时间',
    name: 'update_at',
  })
  updateAt: Date;
}

```




update score logic
```js
async updateTodayScore(addScoreDto: AddScoreDto) {
    console.log(addScoreDto);
    const entity = await this.scoreHistoryRepository.findOne({
      where: {
        partnerId: addScoreDto.partnerId,
        date: addScoreDto.date,
      },
    });
    if (!entity) {
      const latest = await this.findLatest();
      console.log(latest);
      const found = latest.find((e) => e.partnerId == addScoreDto.partnerId);
      console.log('found', found, found.partnerId, addScoreDto.partnerId);
      if (found) {
        const createScoreDto: CreateScoreDto = {
          date: addScoreDto.date,
          partnerId: addScoreDto.partnerId,
          score: found.score + parseInt(String(addScoreDto.add)),
        };
        await this.create(createScoreDto);
      }

      const another = latest.find((e) => e.partnerId != addScoreDto.partnerId);
      console.log('another', another);
      const createDto: CreateScoreDto = {
        date: addScoreDto.date,
        partnerId: another.partnerId,
        score: another.score,
      };
      return await this.create(createDto);
    } else {
      //由于没有使用Pipe 导致这里接受到的是字符串 而不是number
      const score = parseInt(String(addScoreDto.add)) + entity.score;
      console.log('update score', score);
      entity.score = score;
      return this.scoreHistoryRepository.save(entity);
    }
  }
```