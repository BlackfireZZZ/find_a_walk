import React from 'react';

function ContactAdd(){
    
}

const Profile = ({ user }) => (
    <div id="Profile" style={{display: 'inline-block'}}>
        <div id="username"><h3>{user.nickname}</h3></div>
        <p style={{margin: '-15px 0 10px 0'}}>Nickname</p>
        <div id="userage"><h3>{user.getAge()} y.o.</h3></div>
        <p style={{margin: '-15px 0 10px 0'}}>Age</p>
        <div id="usercontacts" style={{padding: "10px"}}>
            {user.contacts.map((contact, index) => (
                <div
                    key={index}
                    style={{ borderRadius: '20px', backgroundColor: 'grey', padding: '1px'}}
                >
                    <h5 style={{ color: 'white'}}>{contact}</h5>
                </div>
            ))}
            <input 
            id = "contact_input"
            type="placeholder" 
            placeholder="Phone, Telegram, E-mail e.t.c..."
            style = {{width: "98%"}}></input>
            <input 
            type="button" 
            value="+ Add new contact"
            style = {{width: "98%"}}
            onClick={ContactAdd()}></input>
        </div>
    </div>
);

export default Profile;
