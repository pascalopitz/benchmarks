var fs = require('fs');
var http = require('http');

var matcher = /\/(\w+)/;

var server = http.createServer(function (req, res) {

	var m = req.url.match(matcher);
	var d = new Date();
	var params = {
		'.Name' : m[1],
		'.Time' : d.getHours() + ':' + d.getMinutes() + (d.getHours() > 12 ? 'pm' : 'am')
	};

	fs.readFile('./index.html', function (err, data) {
  		if (err) throw err;

  		data = String(data);

  		for(i in params) if(params.hasOwnProperty(i)) {
  			data = data.replace('{{'+i+'}}', params[i]);
  		}

  		res.end(data);
  	});

});

server.listen(8080);