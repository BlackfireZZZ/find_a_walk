import React from 'react';

const LoginCheck = () => {
    //проверка на правильность имени и пароля
};
const Redirect = () => {
    //перенаправление я уже не сделаю
}
const LoginScreen = () => (
    <div style={{position: 'absolute', left: '40%', top: '30%', width: '20%', border: '1px black solid', borderRadius: '20px', padding: '10px'}}>
        <h1>Findy. Log in</h1>
        <input style={{border: '1px black solid', width: '95%'}} placeholder="Nickname"/>
        <br></br>
        <input type="password" style={{border: '1px black solid', width: '95%'}} placeholder="Password"/>
        <br></br>
        <input class='ToGoButton' style={{width: '97%'}} type="button" value="Log in"/>
        <input style={{width: '97%', borderRadius: '10px'}} type="button" value="Register"/>
    </div>
);

export {LoginScreen, LoginCheck};
