const BaseURL = 'http://localhost:8082'
export const myRequest = (opthion) => {
	return new Promise((resolve,reject) => {
		
		uni.request({
			url: BaseURL+opthion.url,
			data: opthion.data || {},
			method: opthion.method || 'GET',
			success: res => {
				if (res.data.status !== 0) {
					return uni.showToast({
						title:'数据获取失败！'
					})
				}
				resolve(res)
			},
			fail: err => {
					return uni.showToast({
						title:'数据获取失败！'
					})
				reject(err)
			}
		})
	})
}