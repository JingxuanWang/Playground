// bind a function with a specific scope
function bind(func, scope){
	return function(){
		return func.apply(scope, arguments);
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

// sum up some prop in an array of objects
function sum(arr, prop) {
	return arr.map(function(k) {
		return k[prop];
	}).reduce(function(a, b) {
		return a + b;
	});
}

// sum up some prop in an array of objects
function sum2(arr, func) {
	return arr.map(function(k) {
		return func(k);
	}).reduce(function(a, b) {
		return a + b;
	});
}

// sort an object array by it's property
function sortByProp(arr, prop, order) {
	return arr.sort(function(a, b) {
		return order * (a[prop] - b[prop]);
	});
}

//Schwartzian transform
function sortByFunc(arr, func, order) {
	return arr.map(function (x) {
		return [x, func(x)];
	}).sort(function (a, b) {
		return order * (a[1] - b[1]);
	}).map(function (x) {
		return x[0];
	});
}

function lot(arr, func, total_prob) {
	total_prob ||= sum2(arr, func);
	var rand = Math.floor((Math.random() * total_prob) + 1);  // 1 ~ total_prob
	for (var i = 0; i < arr.length; i++) {
		rand -= func(arr);
		if (rand < 0) {
			return arr[i];
		}	
	}
}
