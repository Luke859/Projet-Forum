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