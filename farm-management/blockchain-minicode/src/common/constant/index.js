// 开始仓储表单字段
export const warStartFormItems = {
  ruleFormItems: [
    {
      label: '开始时间',
      type: 'picker',
      name: 'startDate',
      required: true,
      mode: 'date'
    },
    {
      label: '仓储厂',
      type: 'picker',
      name: 'storageFactory',
      required: true,
      mode: 'selector'
    },
    {
      label: '进厂重量',
      type: 'input',
      name: 'enterWeight',
      required: true,
      unit: 'kg'
    },
    {
      label: '冷藏温度',
      type: 'input',
      name: 'refTemperature',
      required: true,
      unit: '℃'
    },
    {
      label: '仓储批次',
      type: 'input',
      name: 'storage_batch',
      required: false,
      disabled: true
    }
  ],
  qrBusFormItems: [
    {
      label: '二维码编号',
      type: 'qrcode',
      name: 'qrCode',
      required: true
    },
    {
      label: '重量',
      type: 'input',
      name: 'weight',
      required: true,
      unit: 'kg'
    }
  ],
  imgFormItems: [
    { label: '照片', type: 'image', name: 'images', required: false }
  ]
}

export const bzFormItems = {
  ruleFormItems: [
    {
      label: '加工厂',
      name: 'upFlowEnterpriseName'
    },
    {
      label: '包装时间',
      name: 'packDate'
    },
    {
      label: '包装规格',
      name: 'machPackNormList'
    }
  ],

  imgFormItems: [
    { label: '照片', type: 'image', name: 'images', required: false }
  ]
}
export const jgFormItems = {
  ruleFormItems: [
    {
      label: '加工批次',
      name: 'jgBatch'
    },
    {
      label: '加工厂',
      name: 'upFlowEnterpriseName'
    },
    {
      label: '加工时间',
      name: 'jgDate'
    },
    {
      label: '出厂重量',
      name: 'ccWeight'
    }
  ],

  imgFormItems: [
    { label: '照片', type: 'image', name: 'images', required: false }
  ]
}
export const hgEndFormItems = {
  ruleFormItems: [
    {
      label: '烘干批次',
      name: 'hgBatch'
    },
    {
      label: '烘干厂',
      name: 'hgFactoryName'
    },

    {
      label: '开始时间',
      name: 'hgDate'
    },

    {
      label: '烘干后含水量',
      name: 'hgHsl',
      unit: '%'
    },
    {
      label: '烘干后含重量',
      name: 'hgWeight',
      unit: 'kg'
    }
  ],
  qrBusFormItems: [
    {
      label: '二维码编号',
      type: 'qrcode',
      name: 'qrCode',
      required: true
    },
    {
      label: '重量',
      type: 'input',
      name: 'weight',
      required: true,
      unit: 'kg'
    }
  ],
  imgFormItems: [
    { label: '照片', type: 'image', name: 'images', required: false }
  ]
}
// 结束仓储表单字段
export const warEndFormItems = {
  ruleFormItems: [
    {
      label: '仓储厂',
      name: 'ccFactoryName'
    },
    {
      label: '仓储批次',
      name: 'storageBatch'
    },

    {
      label: '开始时间',
      name: 'startDate'
    },
    {
      label: '结束时间',
      name: 'endDate'
    },
    {
      label: '进厂重量',
      name: 'enterWeight',
      unit: 'kg'
    },
    {
      label: '出厂重量',
      name: 'outWeight',
      unit: 'kg'
    }
  ],
  qrBusFormItems: [
    {
      label: '二维码编号',
      type: 'qrcode',
      name: 'qrCode',
      required: true
    },
    {
      label: '重量',
      type: 'input',
      name: 'weight',
      required: true,
      unit: 'kg'
    }
  ],
  imgFormItems: [
    { label: '照片', type: 'image', name: 'images', required: false }
  ]
}

// 仓储待处理列表字段
export const storageDisList = [
  {
    name: '',
    filds: [
      { key: '品种：', value: '' },
      { key: '时间：', value: '' },
      { key: '重量（kg）：', value: '' },
      { key: '任务：', value: '' }
    ]
  }
]

// 仓储已处理列表字段
export const storageList = [
  {
    name: '',
    filds: [
      { key: '开始：', value: '' },
      { key: '结束：', value: '' },
      { key: '品种：', value: '' },
      { key: '烘干厂：', value: '' },
      { key: '出仓重量（kg）：', value: '' },
      { key: '出仓二维码编号：', value: '' }
    ]
  }
]

// 加工待处理列表字段
export const macDisList = [
  {
    name: '',
    filds: [
      { key: '品种：', value: '' },
      { key: '时间：', value: '' },
      { key: '重量（kg）：', value: '' },
      { key: '任务：', value: '' }
    ]
  }
]

// 加工已处理列表字段
export const macList = [
  {
    name: '',
    filds: [
      { key: '时间：', value: '' },
      { key: '品种：', value: '' },
      { key: '名称：', value: '' },
      { key: '规格：', value: '' },
      { key: '上环节处理场所：', value: '' },
      { key: '溯源码编号：', value: '' }
    ]
  }
]

// 加工信息表单字段
export const macFormItems = {
  ruleFormItems: [
    {
      label: '加工日期',
      type: 'picker',
      name: 'jgDate',
      required: true,
      mode: 'date'
    },
    {
      label: '加工厂',
      type: 'picker',
      name: 'storageFactory',
      required: true,
      mode: 'selector'
    },
    {
      label: '出厂重量',
      type: 'input',
      name: 'ccWeight',
      required: true,
      unit: 'kg'
    },
    {
      label: '加工批次',
      type: 'input',
      name: 'batchId',
      required: false,
      disabled: true
    }
  ],
  qrBusFormItems: [
    {
      label: '二维码编号',
      type: 'qrcode',
      name: 'qrCode',
      required: true
    },
    {
      label: '重量',
      type: 'input',
      name: 'weight',
      required: true,
      unit: 'kg'
    }
  ],
  imgFormItems: [
    { label: '照片', type: 'image', name: 'images', required: false }
  ]
}
