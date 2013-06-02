var w = window,
  d = w.document,
	xmlhttp = w.XMLHttpRequest ? XMLHttpRequest() : ActiveXObject("Microsoft.XMLHTTP"),
	data = '';

xmlhttp.onreadystatechange=function(){
	if (xmlhttp.readyState==4 && xmlhttp.status==200){
		data = xmlhttp.responseText;
	}
}

setInterval(function(){
	xmlhttp.open('get',"http://"+document.location.host+"/data",true);
	xmlhttp.send();
},2000);
