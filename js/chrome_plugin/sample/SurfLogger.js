// Author: Wang Jingxuan
// Created: 2013.3.13


var SurfLogger = {

	storage: chrome.storage.local,

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

		/*
		var local = {
			logger : {
				'www.sohu.com' : 1,
				'www.sina.com.cn' : 2,
				//...
			}
		};
		*/

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
	list: function() {
		chrome.storage.local.get(null, function(item) {
			var list_table = "<table>";
		
			for (var host in item) {
				console.log("===> host: "+ host +" count: "+ item[host]);
				list_table += "<tr colspan=2><td width=100>";
				list_table += host;
				list_table += "</td><td width=50>";
				list_table += item[host];
				list_table += "</td></tr>";
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
	//SurfLogger.test();
	//SurfLogger.clear();
	//SurfLogger.log();
	//SurfLogger.list();
	// bind buttons
	var btn_clear = document.getElementById('btn_clear');
	var btn_log = document.getElementById('btn_log');
	var btn_list = document.getElementById('btn_list');

	btn_clear.addEventListener('click', function() {
		SurfLogger.clear();
	});


	/*
	btn_log.addEventListener('click', function() {
		SurfLogger.log();
	});
	btn_list.addEventListener('click', function() {
		SurfLogger.list();
	});
	*/

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

