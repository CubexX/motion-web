function do_motion() {
	var request = new XMLHttpRequest();
	request.open('POST', '/do_motion', true);
	request.send();
}

function check_motion() {
	var request = new XMLHttpRequest();
	request.open('GET', '/check_motion', true);

	request.onload = function() {
	  if (this.status >= 200 && this.status < 400) {
	    var resp = this.response;

	    var status = document.getElementById("status");
		if (resp) {
			status.innerHTML = "on";
			status.style.color = "green";
		} else {
			status.innerHTML = "off";
			status.style.color = "red";
		}
	  }
	};

	request.send();
}

setInterval(function(){
	check_motion();
}, 3000);
