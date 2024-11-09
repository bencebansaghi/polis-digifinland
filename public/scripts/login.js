const signupButton = document.getElementById("newUser");

signupButton.addEventListener('click', function() {
    document.getElementById("log-items").style.visibility = this.hidden;

    const captcha = document.getElementById("captcha");
    captcha.style.visibility = this.visibile;
});