/* mmover.js */

var connection;
var map, marker;




$(document).ready(function()
{
	connection = new Connection("127.0.0.1", location.port);

	map = new GMaps(
		{
			div: '#map',
			lat: 43.7679289,
			lng: 11.2509548,
			panControl : false,
			streetViewControl : false,
			mapTypeControl: false,
			overviewMapControl: false,
			center_changed: onMove
		});
	map.setZoom(21);

	marker = map.addMarker({
					lat: 43.7679289,
					lng: 11.2509548,
					title: "Your Position"
				});

});


function onMove(e)
{
	var data = {lat: map.getCenter().lat().toString(), lng: map.getCenter().lng().toString(), update_rate_ms: 1000};

	marker.setPosition(map.getCenter());

	connection.send(data);
}
