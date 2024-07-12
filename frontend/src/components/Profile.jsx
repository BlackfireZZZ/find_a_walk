import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Cookies from 'js-cookie';
import config from '../config';

class User {
    constructor(id, name, email, interests) {
        this.id = id;
        this.name = name;
        this.email = email;
        this.interests = interests || [];
    }

    getAge() {
        return new Date().getFullYear() - this.birthdate[2];
    }
}

const ContactShow = ({ index, contact }) => (
    <div key={index} className='ContactElement'>
        <h5 style={{ color: 'white' }}>{contact}</h5>
    </div>
);

const ContactAdd = () => {
    let div = document.getElementById('usercontacts');
    let newcontact = document.createElement('div');
    newcontact.className = "ContactElement";
    newcontact.innerHTML = `<h5 style="color: white">${document.getElementById('contact_input').value}</h5>`;
    div.prepend(newcontact);
}

const Profile = ({ user }) => (
    <div id="Profile" style={{ display: 'inline-block' }}>
        <div id="username"><h3>{user.name}</h3></div>
        <p style={{ margin: '-15px 0 10px 0' }}>Псевдоним</p>
        <div id="userage"><h3>{user.getAge()} лет</h3></div>
        <p style={{ margin: '-15px 0 10px 0' }}>Возраст</p>
        <div id="usercontacts" style={{ padding: "10px" }}>
            {user.interests.map((contact, index) => (
                <ContactShow index={index} contact={contact.name} />
            ))}
            <input
                id="contact_input"
                type="text"
                placeholder="Phone, Telegram, E-mail e.t.c..."
                style={{ width: "95%" }}></input>
            <input
                type="button"
                value="+ Добавить контакт"
                style={{ width: "98%" }}
                onClick={ContactAdd}></input>
        </div>
    </div>
);

const fetchLoggedUser = async () => {
    const token = Cookies.get('jwt');
    if (!token) {
        window.location.href = '/login';
        return null;
    }

    try {
        const response = await axios.get(`${config.Host_url}/users/me`, {
            headers: { Authorization: `Bearer ${token}` }
        });
        const data = response.data;
        return new User(data.id, data.name, data.email, data.interests);
    } catch (error) {
        console.error('Error fetching user data:', error);
        window.location.href = '/login';
        return null;
    }
};

let users = [
    new User('Chinese developers Team',
        [1, 7, 2006],
        ['89456783542', 'nightmarefuel'],
        [55.7558, 37.6176],
    ),
    new User('IvanGutche',
        [1, 7, 2006],
        ['@stupidcabage']
    )
];

const App = () => {
    const [user, setUser] = useState(null);

    useEffect(() => {
        const getUser = async () => {
            const fetchedUser = await fetchLoggedUser();
            if (fetchedUser) {
                setUser(fetchedUser);
            }
        };
        getUser();
    }, []);

    return user ? <Profile user={user} /> : <div>Loading...</div>;
};

export { App as Profile, users };
