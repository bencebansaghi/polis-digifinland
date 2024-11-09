const popUp = (a, b, ) => {
    const captchaBox = document.querySelector('.captcha-box');
    const popUp = document.createElement('div');
    popUp.className = 'popup';
    popUp.innerHTML = `
        <p>Solve this challenge to proceed:</p>
        <p id="captcha-prompt"><strong>{{ .NumberA }} + {{ .NumberB }}</strong></p>
        <input type="text" id="captcha-answer">
        <button onclick="verifyCaptcha()">Submit</button>
    `;
}