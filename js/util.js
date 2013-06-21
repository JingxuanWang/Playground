// bind a function with a specific scope
function bind(func, scope){
    return function(){
        return func.apply(scope, arguments);
    };
}       

// sort an object array by it's property
function sortBy(prop, order) {
    return function(a, b) {
        return order * (a[prop] - b[prop]);
    };
}   

// deep clone a object
function clone(obj){
    if(obj == null || typeof(obj) != 'object')
        return obj;

    var temp = {};
    //var temp = obj.constructor(); // changed

    for(var key in obj)
        temp[key] = clone(obj[key]);
    return temp;
}

