$('#signup').on('submit', signup)

function signup(event){
    event.preventDefault()

    password = $('#password').val()
    passwordConfirm = $('#password-confirm').val()

    if (!isPasswordsMatches(password, passwordConfirm)) {
        alert("Password does not match")
        return
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $('#name').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            password,
        }
    })

}

function isPasswordsMatches(password, passwordConfirm) {
    return password === passwordConfirm
}