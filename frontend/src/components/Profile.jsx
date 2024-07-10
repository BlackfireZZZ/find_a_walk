import React from 'react';

class User {
    constructor(nickname, birthdate, contacts, coords){
        this.nickname = nickname;
        this.birthdate = birthdate;
        this.contacts = contacts;
        this.coords = coords;
    }
    getAge(){
        return new Date().getFullYear() - this.birthdate[2];
    }
}
let users = [
    new User('Chinese developers Team',
        [1, 7, 2007],
        ['89456783542', 'nightmarefuel'],
        [55.7558, 37.6176],
    ),
    new User('IvanGutche',
        [1, 7, 2006],
        ['@stupidcabage']
    )
];
const loggedUser = users[0];

const ContactShow = ({index, contact}) => (
    <div key={index} className='ContactElement'>
        <h5 style={{ color: 'white'}}>{contact}</h5>
    </div>
);
const ContactAdd = () => {
    let div = document.getElementById('usercontacts');
    let newcontact = document.createElement('div');
    newcontact.className = "ContactElement";
    newcontact.innerHTML = `<h5 style="color: white">${document.getElementById('contact_input').value}</h5>`;
    //Отправить данные бэкендерам
    div.prepend(newcontact);
}
const Profile = ({ user }) => (
    <div id="Profile" style={{display: 'inline-block'}}>
        <div id="username"><h3>{user.nickname}</h3></div>
        <p style={{margin: '-15px 0 10px 0'}}>Псевдоним</p>
        <div id="userage"><h3>{user.getAge()} лет</h3></div>
        <p style={{margin: '-15px 0 10px 0'}}>Возраст</p>
        <div id="usercontacts" style={{padding: "10px"}}>
            {user.contacts.map((contact, index) => (
                <ContactShow index={index} contact={contact} />
            ))}
            <input 
            id = "contact_input"
            type="placeholder" 
            placeholder="Phone, Telegram, E-mail e.t.c..."
            style = {{width: "95%"}}></input>
            <input 
            type="button" 
            value="+ Добавить контакт"
            style = {{width: "98%"}}
            onClick={ContactAdd}></input>
        </div>
    </div>
);

export {Profile, users, loggedUser};
