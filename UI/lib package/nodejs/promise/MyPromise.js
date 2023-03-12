class MyPromise {

    static PENDING = 'PENDING';
    static FULFILLED = 'OK';
    static REJECTED= 'FAIL';

    constructor(func) {
        this.status = MyPromise.PENDING;
        this.result = null;

        this.resolveCallbacks = [];
        this.rejectCallbacks = [];

        try {
            func(this.resolve.bind(this), this.reject.bind(this));
        } catch (error) {
            this.reject(error)
        }
        
    }

    resolve(result) {
        if(this.status === MyPromise.PENDING){
            this.status = MyPromise.FULFILLED;
            this.result = result;

            this.resolveCallbacks.forEach(callback => {
                callback(result);
            })
        }
    }

    reject(reason) {
        if(this.status === MyPromise.PENDDING){
            this.status = MyPromise.REJECTED;
            this.result = result;

            this.rejectCallbacks.forEach(callback => {
                callback(result);
            })
        } 
    }

    then(onFULFILLED, onREJECTED) {
        //check callback function
        onFULFILLED = typeof onFULFILLED === 'function' ? onFULFILLED : () => {};
        onREJECTED = typeof onREJECTED === 'function' ? onREJECTED : () => {};

        if(this.status === MyPromise.PENDING){
            this.resolveCallbacks.push(onFULFILLED);
            this.rejectCallbacks.push(onREJECTED);
        }

        if(this.status === MyPromise.FULFILLED){
            setTimeout(() => {
                onFULFILLED(this.result);
            });
        } 
        if(this.status === MyPromise.REJECTED){
            setTimeout(() => {
                onREJECTED(this.result);
            });
        } 
    }
}


let mm = new MyPromise((resolve, reject) => {
    resolve('aaaaaaaaaaaaaaaa');
})

mm.then(
    result => console.log(result),
    result => console.log(result)
)