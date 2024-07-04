import React from 'react';

const Profile = ({ user }) => (
    <div id="Profile" style={{display: 'inline-block', top: '0'}}>
        <div id="username"><h3>{user.nickname}</h3></div>
        <p style={{margin: '-15px 0 10px 0'}}>Nickname</p>
        <div id="userage"><h3>{user.getAge()}</h3></div>
        <p style={{margin: '-15px 0 10px 0'}}>Age</p>
        <div id="usercontacts">
            {user.contacts.map((contact, index) => (
                <div
                    key={index}
                    style={{ borderRadius: '20px', backgroundColor: 'grey', padding: '1px'}}
                >
                    <h5 style={{ color: 'white'}}>{contact}</h5>
                </div>
            ))}
        </div>
    </div>
);

export default Profile;