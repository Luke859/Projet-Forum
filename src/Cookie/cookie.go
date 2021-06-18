import(
	uuid "github.com/satori/go.uuid"
)
myuuid, _ := uuid.NewV4()

document.cookie = "test1="+ myuuid.String() +"; SameSite=None; Secure; HttpOnly";

const cookieValue1 = document.cookie
	.split('; ')
	.find(row => row.startsWith('test1='))
	.split('=')[1];

func showCookieValue(){
	const output = document.getElementById('cookie-value')
	output.textContent = '> ' + cookieValue1
	return output
}
