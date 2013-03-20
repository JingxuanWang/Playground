// Author: Wang Jingxuan
// Created: 2013.3.13


var SurfLogger = {

	storage: chrome.storage.local,
	DEFAULT_LIMIT: 10,

	/**
	 * Logging visited site & Count 
	 * each time when chrome jump to a URL
	 *
	 * @public
	 */
	log: function(details) {
		// parse host from current connecting site
		var url = details.url;
		var url_slices = url.split("/");
		var host = url_slices[2];

		if (!host) {
			alert("host is empty: " + host);
		}

		var count = 0;

		//chrome.storage.local.get(host, function(items) {
		chrome.storage.local.get(null, function(item) {
			if (item) {
				console.log("===>get: " + url);
				console.log(item);
				if (!item[host]) {
					item[host] = 0;
				}
				++item[host];

				chrome.storage.local.set(item, function() {
					// Notify that we saved.
					console.log('Settings saved '+ url + " : " + item[host]);
					SurfLogger.list();
				});
			} else {
				// if we do not have that before
				// create a new one
				chrome.storage.local.set({'logger': item}, function() {
					// Notify that we saved.
					console.log('Logger Created!');
				});
			}
		});
	},
	/**
	 * List most frequent visited sites today
	 * each time when chrome jump to a URL
	 *
	 * @public
	 */
	clear: function() {
		chrome.storage.local.clear(function() {
			SurfLogger.list();
		});
	},

	/**
	 * List most frequent visited sites today
	 * each time when chrome jump to a URL
	 *
	 * @public
	 */
	list: function(limit) {
		if (!limit) {
			limit = this.DEFAULT_LIMIT;
		} 
		chrome.storage.local.get(null, function(item) {
			var sorted_item = [];
			item = sortByValue(item);

			//var list_table = "<table border=1>";
			var list_table = "<table>";
				list_table += "<tr colspan=3><td width=30>";
				list_table += "Rank";
				list_table += "</td><td width=250 align=center>";
				list_table += "Host";
				list_table += "</td><td width=50 align=right>";
				list_table += "Count";
				list_table += "</td></tr>";

			for (var i = 0; i < item.length; i++) {
				console.log("===> host: "+ item[i][0] +" count: "+ item[i][1]);
				list_table += "<tr colspan=3><td width=30>";
				list_table += i + 1;
				list_table += "</td><td width=250>";
				list_table += "<a href=\"http://"+item[i][0]+"\">"+item[i][0]+"</a>";
				list_table += "</td><td width=50 align=right>";
				list_table += item[i][1];
				list_table += "</td></tr>";
				
				if (i == limit - 1) {
					break;
				}
			}
			list_table += "</table>";
		
			document.getElementById('list').innerHTML = list_table;
		});
	
	},

	/**
	 * List most frequent visited sites today
	 * each time when chrome jump to a URL
	 *
	 * @public
	 */
	test: function() {
		alert("hello world");
	}
};

// Run our kitten generation script as soon as the document's DOM is ready.
document.addEventListener('DOMContentLoaded', function () {
	SurfLogger.storage = chrome.storage.local;
	
	// bind buttons
	var btn_clear = document.getElementById('btn_clear');

	btn_clear.addEventListener('click', function() {
		SurfLogger.clear();
	});

	SurfLogger.list();
});


chrome.storage.onChanged.addListener(function(changes, namespace) {
	for (key in changes) {
		var storageChange = changes[key];
		console.log(
			'Storage key "%s" in namespace "%s" changed. ' +
			'Old value was "%s", new value is "%s".',
			key,
			namespace,
			storageChange.oldValue,
			storageChange.newValue
		);
	}
});

/*
chrome.webRequest.onCompleted.addListener(
	function(details){
		console.log(details);
		//SurfLogger.log();
	},
	{urls: ["<all_urls>"]},
	[]
);
//*/

/*
chrome.webNavigation.onCompleted.addListener(function(details){
	SurfLogger.log(details.url);
});
//*/
/*
chrome.webNavigation.onCommitted.addListener(function(details){
	SurfLogger.log(details.url);
});
//*/

///*
chrome.history.onVisited.addListener(function(details) {
	SurfLogger.log(details);
});
//*/


function sortByValue(obj, callback,  context) {
	var tuples = [];
	for (var key in obj) {
		tuples.push([key, obj[key]]);
	}
	
	/*
	 * tuples[0] = [key, obj[key]]
	 * tuples[1] = [key, obj[key]]
	 * ...
	 */

	tuples.sort(function(a, b) {
		return a[1] < b[1] ? 1 : a[1] > b[1] ? -1 : 0 
	});
	
	var length = tuples.length;
	return tuples;
}

