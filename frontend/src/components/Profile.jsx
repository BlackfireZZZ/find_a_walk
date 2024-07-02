import React from 'react';

const Profile = ({ user }) => (
    <div id="Profile">
        <div id="username">{user.nickname}</div>
        <div id="userage">{user.getAge()}</div>
        <div id="usercontacts">
            {user.contacts.map((contact, index) => (
                <div
                    key={index}
                    style={{ borderRadius: '20px', backgroundColor: 'grey' }}
                >
                    <h5 style={{ color: 'white', marginTop: '-10px' }}>{contact}</h5>
                </div>
            ))}
        </div>
    </div>
);

export default Profile;