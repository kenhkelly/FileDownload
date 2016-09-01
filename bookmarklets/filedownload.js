var dl_serv = "http://localhost:8080/get?url=";
var href = null;
var w = null;

if (el = document.querySelector('._n3cp9._d20no%20._jjzlb%20img')) { //Instagram image
  href = el.src;
}
else if (el = document.querySelector('._n3cp9._d20no%20._2tomm%20video')) { // Instagram video
  href = el.src;
}
else if (el = document.querySelector('.spotlight')) { // Facebook image
  href = el.src;
}

if (href != null) {
  w = window.open(dl_serv + encodeURIComponent(href));
}

if (w !== null) {
  setTimeout(function(){
    w.close();
  }, 500);
}
