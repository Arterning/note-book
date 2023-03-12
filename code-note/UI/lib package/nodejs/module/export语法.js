//简写
var tag = 'in a.js';
exports.doAdd = function(x,y) {
    console.log(tag,x,y);
    return x+y;
};

//或者
module.exports = {
    fieldx: a,
    setFieldX: function(nv) {
        this.fieldx = nv;
    }
};