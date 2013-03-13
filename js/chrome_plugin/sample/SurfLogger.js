// Copyright (c) 2012 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.


var SurfLogger = {

	storage: chrome.storage.local,

	/**
	 * Logging visited site & Count 
	 * each time when chrome jump to a URL
	 *
	 * @public
	 */
	log: function() {
		// parse host from current connecting site
		var url = "www.sohu.com/news";
		var host = "www.sohu.com";

		var local = {
			logger : {
				'www.sohu.com' : 1,
				'www.sina.com.cn' : 2,
				//...
			}
		};


		var count = 0;

		//chrome.storage.local.get(host, function(items) {
		chrome.storage.local.get(null, function(item) {
			if (item) {
				console.log("===>get: " + url);
				console.log(item);
				if (!item[host]) {
					item[host] = 0;
				}
			}
			
			++item[host];

			chrome.storage.local.set(item, function() {
				// Notify that we saved.
				console.log('Settings saved '+ url + " : " + item[host]);
			});
		});
	},
	/**
	 * List most frequent visited sites today
	 * each time when chrome jump to a URL
	 *
	 * @public
	 */
	clear: function() {
		chrome.storage.local.clear();
	},

	/**
	 * List most frequent visited sites today
	 * each time when chrome jump to a URL
	 *
	 * @public
	 */
	list: function() {
		chrome.storage.local.get(null, function(item) {
			for (var host in item) {
				console.log("===> host: "+ host +" count: "+item[host]);
			}
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
	SurfLogger.log();
	SurfLogger.list();

});

chrome.storage.onChanged.addListener(function(changes, namespace) {
	for (key in changes) {
		var storageChange = changes[key];
		console.log('Storage key "%s" in namespace "%s" changed. ' +
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
*/

chrome.webNavigation.onCompleted.addListener(function(details){
	console.log(details);
	alert("hey");
});

