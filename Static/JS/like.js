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
