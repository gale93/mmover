
function Connection(address, port)
{
    console.log('Connecting to server...');

	var owner = this;
	this.address = address;
	this.port = port;
    this.connection = new WebSocket("ws://" + address + ":" + port + "/socket");
	this.connected = false;

    this.connection.onopen = function()
    {
		owner.connected = true;
        console.log('Connected!');
		createMap();
    }

    this.connection.onclose = function()
    {
		owner.connected = false;
        console.log('Retrying to connect in 3 seconds...');

		var conn = this;
		setTimeout(function(conn){ conn = new Connection(owner.address, owner.port); }, 3000);
    }
    this.connection.onerror = function()
    {
        console.log('[Connection Error]');
    }
    this.connection.onmessage = function(msg)
    {
		console.log(msg.data);
		readMessage(JSON.parse(msg.data));
    }

    var conn_ptr = this.connection;
    this.send = function(obj) {
		if (conn_ptr != null && owner.connected)
        	conn_ptr.send(JSON.stringify(obj));
    }
}


function readMessage(data)
{
	switch (data.header)
	{
		case "starting_pos":
		{
			$("#inlat").val(data.body["lat"]);
			$("#inlng").val(data.body["lng"]);
			applyCoordinates();
		}

		break;
	}
}
