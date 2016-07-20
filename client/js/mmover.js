/* mmover.js */

var connection;
var map, marker;




$(document).ready(function()
{
	connection = new Connection("127.0.0.1", location.port);

});

function createMap()
{
	map = new GMaps(
		{
			div: '#map',
			lat: 0,
			lng: 0,
			panControl : false,
			streetViewControl : false,
			mapTypeControl: false,
			overviewMapControl: false,
			center_changed: onMove
		});
	map.setZoom(21);

	marker = map.addMarker({
					lat: 0,
					lng: 0,
					title: "Your Position"
				});
}


function onMove(e)
{
	var data = {lat: map.getCenter().lat().toString(), lng: map.getCenter().lng().toString(), update_rate_ms: 1000};

	marker.setPosition(map.getCenter());
	$("#inlat").val(data.lat);
	$("#inlng").val(data.lng);


	connection.send(data);
}


function applyCoordinates()
{
	map.setCenter($("#inlat").val(),$("#inlng").val());
	onMove(null);
}
