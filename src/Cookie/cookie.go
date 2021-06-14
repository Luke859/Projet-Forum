document.cookie = "test1=Hello; SameSite=None; Secure; HttpOnly";
document.cookie = "test2=World; SameSite=None; Secure; HttpOnly";

const cookieValue1 = document.cookie
  .split('; ')
  .find(row => row.startsWith('test1='))
  .split('=')[1];

const cookieValue2 = document.cookie
.split('; ')
.find(row => row.startsWith('test2='))
.split('=')[1];


function showCookieValue() {
  const output = document.getElementById('cookie-value')
  output.textContent = '> ' + cookieValue1 + " " + cookieValue2
}
