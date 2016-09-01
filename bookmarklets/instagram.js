var dl_serv = "http://localhost:8080/get?url=";
var el = document.querySelector('._n3cp9._d20no%20._jjzlb%20img');
var w = null;
if (typeof el !== 'undefined' && el !== null) {
  href = el.src;
  w = window.open(dl_serv + href);
} else {
  var el = document.querySelector('._n3cp9._d20no%20._2tomm%20video');
  href=el.src;
  w = window.open(dl_serv + href);
}
if (w !== null) {
  setTimeout(function(){
    w.close();
  }, 500);
}