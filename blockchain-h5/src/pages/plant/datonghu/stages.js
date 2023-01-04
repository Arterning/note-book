const stages = [
  {
    name: '苗期',
    steps: [
      {
        date: '2021-03-18',
        title: '选种浸种',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [
            '（1）用清水选种，去除空秕粒。',
            '（2）咪鲜胺浸种消毒8-12小时后，再按昼浸夜露清水浸种2昼夜，沥干。',
            '（3）高温破胸露白，根长不超过2mm。晾干。'
          ],
          imgs: []
        }
      },
      {
        date: '2021-03-21',
        title: '催芽',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [
            '（1）采用水稻有序机抛播种流水线播种；在播种浇底水时，用1000倍-2000倍液浓度的敌璜钠或者多菌灵药液对秧盘中钵土进行消毒。',
            '（2）置控温（28-35℃）保湿（湿度＞90%）的暗室中1-2天促芽至90%以上出芽露尖。'
          ],
          imgs: []
        }
      },
      {
        date: '2021-03-23',
        title: '叠盘出苗',
        extra: `信息采集员：莫磊`,
        body: {
          desc: ['叠盘高度不超过20层或1m；出芽不能过长。'],
          imgs: []
        }
      },
      {
        date: '2021-03-24',
        title: '摆盘盖膜',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [
            '（1）平整好厢面。',
            '（2）铺上隔泥层（编织布或无纺布）。',
            '（3）按厢面中间对开，安放准绳，秧盘沿准绳整齐摆设整齐。摆盘时两盘并列对放，之间紧密无空隙，盘底紧贴厢面。',
            '（4）对厢面喷施杀菌剂，用敌克松1包兑水15公斤喷施厢面作土壤消毒。',
            '（5）插好拱棚竹片，盖拱膜，四周封严封实。',
            '（6）厢面太干时，灌一次跑马水。'
          ],
          imgs: []
        }
      },
      {
        date: '2021-04-17',
        title: '施有机肥',
        extra: `信息采集员：莫磊`,
        body: {
          desc: ['有机质≥70% 有机肥：100公斤/亩'],
          imgs: []
        }
      },
      {
        date: '2021-04-18',
        title: '旋耕',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [],
          imgs: []
        }
      },
      {
        date: '2021-04-27',
        title: '移栽',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [],
          imgs: []
        }
      },
      {
        date: '2021-04-29',
        title: '封闭除草',
        extra: `信息采集员：莫磊`,
        body: {
          desc: ['丁草胺100ML/亩 + 苄嘧磺隆10克/亩'],
          imgs: []
        }
      },
      {
        date: '2021-05-01',
        title: '人工补苗',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [],
          imgs: []
        }
      }
    ]
  },
  {
    name: '分蘖期',
    steps: [
      {
        date: '2021-05-07',
        title: '施分蘖肥',
        extra: `信息采集员：莫磊`,
        body: {
          desc: ['尿素7.5公斤/亩'],
          imgs: []
        }
      },
      {
        date: '2021-06-12',
        title: '排水晒田',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [],
          imgs: []
        }
      }
    ]
  },
  {
    name: '孕穗期',
    steps: [
      {
        date: '2021-06-18',
        title: '浅水间歇灌溉',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [],
          imgs: []
        }
      },
      {
        date: '2021-06-22',
        title: '施穗肥',
        extra: `信息采集员：莫磊`,
        body: {
          desc: ['配方肥：（25-10-16 ）10公斤/亩'],
          imgs: []
        }
      }
    ]
  },
  {
    name: '抽穗扬花期',
    steps: [
      {
        date: '2021-07-15',
        title: '破口期施药',
        extra: `信息采集员：莫磊`,
        body: {
          desc: ['阿维菌素（100克） + 吡蚜酮（24克） + 己唑醇（16克）/亩'],
          imgs: []
        }
      },
      {
        date: '2021-07-17',
        title: '干湿间歇灌溉',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [],
          imgs: []
        }
      }
    ]
  },
  {
    name: '灌浆期',
    steps: [
      {
        date: '2021-07-24',
        title: '病虫害防治',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [
            '己唑醇16g/亩+吡蚜酮12g/亩 + 果肥48g/亩 + 啶氧丙环唑70ml/亩 + 浓缩氨基酸液1kg/亩'
          ],
          imgs: []
        }
      },
      {
        date: '2021-07-27',
        title: '排水晒田',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [],
          imgs: []
        }
      }
    ]
  },
  {
    name: '成熟期',
    steps: [
      {
        date: '2021-08-08',
        title: '收割',
        extra: `信息采集员：莫磊`,
        body: {
          desc: ['测产503.45公斤/亩'],
          imgs: []
        }
      }
    ]
  },
  {
    name: '再生季孕穗',
    steps: [
      {
        date: '2021-08-09',
        title: '追施肥',
        extra: `信息采集员：莫磊`,
        body: {
          desc: ['尿素：15kg/亩'],
          imgs: []
        }
      },
      {
        date: '2021-08-30',
        title: '喷药',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [
            '甲氨基阿维菌素苯甲酸盐13ml/亩 + 吡虫啉30克/亩 + 丙环唑15瓶/亩 + 赤霉素20g/亩'
          ],
          imgs: []
        }
      }
    ]
  },
  {
    name: '再生季灌浆',
    steps: [
      {
        date: '2021-09-11',
        title: '叶面营养',
        extra: `信息采集员：莫磊`,
        body: {
          desc: ['叶肥24g/亩 + 浓缩氨基酸30ml/亩'],
          imgs: []
        }
      }
    ]
  },
  {
    name: '再生季蜡熟',
    steps: [
      {
        date: '2021-09-20',
        title: '排水晒田',
        extra: `信息采集员：莫磊`,
        body: {
          desc: [],
          imgs: []
        }
      }
    ]
  },
  {
    name: '再生季黄熟',
    steps: [
      {
        date: '2021-10-07',
        title: '收割',
        extra: `信息采集员：莫磊`,
        body: {
          desc: ['测产260.68公斤/亩'],
          imgs: []
        }
      }
    ]
  }
]

const arrs = [
  [1, 2, 3, 4, 5, 6, 7, 8, 9],
  [1, 2],
  [1, 2],
  [1, 2],
  [1, 2],
  [1],
  [1, 2],
  [1],
  [1],
  [1]
]

arrs.forEach((arr, i) => {
  arr.forEach((item, j) => {
    stages[i].steps[j].body.imgs.push(
      `https://iland.zoomlion.com/staticfile/blockchain-h5/img/stage${i +
        1}/${item}.jpg`
    )
  })
})

export default stages
