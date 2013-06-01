var data;
var sec=0.1;
setInterval(function(){
  d3.json("http://"+document.location.host+"/data",function(json){
  data=json;
  document.getElementsByTagName("body")[0].innerHTML=JSON.stringify(json);
  })
},1000*sec);
