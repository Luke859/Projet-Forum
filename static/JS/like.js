<<<<<<< HEAD:static/JS/like.js
var like = document.getElementById('addlike');

document.onclick = changeLike;

function changeLike(e) {
    like.textContent = "1";

    if(like.textContent === "1") {
        document.onclick = change2Like;
    }

}

function change2Like(e) {
    like.textContent ="0"

    if(like.textContent === "0") {
        document.onclick = changeLike;
    }
}
=======
document.onclick = changeLike;

function changeLike() {
    like.textContent = "1";

    if(like.textContent === "1") {
        document.onclick = change2Like;
    }

}

function change2Like() {
    like.textContent ="0";

    if(like.textContent === "0") {
        document.onclick = changeLike;
    }
}


// function changeLike(){

//     var like = document.getElementById('like');
//     like.submit()
// }
>>>>>>> 85fb1ee1353def558eb2ee15fc16152ac7cb0006:Static/JS/like.js
