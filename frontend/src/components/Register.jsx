import React, { useState } from 'react';
import Cookies from 'js-cookie';
import '../RegScreen.css';
import config from "../config";
import { redirect } from "react-router-dom";

const RegScreen = () => {
    const [nickname, setNickname] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [interests, setInterests] = useState([]);
    const [currentInterest, setCurrentInterest] = useState('');
    const [emailError, setEmailError] = useState('');

    const addInterest = (e) => {
        if (e.key === 'Enter' && currentInterest) {
            setInterests([...interests, currentInterest]);
            setCurrentInterest('');
        }
    };

    const removeInterest = (index) => {
        setInterests(interests.filter((_, i) => i !== index));
    };

    const validateEmail = (email) => {
        const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return re.test(String(email).toLowerCase());
    };

    const handleRegister = async () => {
        if (!validateEmail(email)) {
            setEmailError('Invalid email format');
            return;
        }
        setEmailError('');

        const user = {
            name: nickname,
            email: email,
            password: password,
        };

        console.log(JSON.stringify(user))
        try {
            const response = await fetch(config.Host_url + 'users', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(user)
            });

            if (response.ok) {
                const data = await response.json();
                const userId = data.id;
                Cookies.set('Authorization', `Bearer ${userId}`);

                // Send interests to the server
                const interestsResponse = await fetch(config.Host_url + 'users/interests', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${userId}`
                    },
                    body: JSON.stringify(interests)
                });

                if (interestsResponse.ok) {
                    console.log('Interests added successfully');
                    return redirect('/');
                } else {
                    console.error('Error adding interests');
                }
            } else {
                console.error('Error registering user');
            }
        } catch (error) {
            console.error('Error:', error);
        }
    };

    return (
        <div className="reg-screen">
            <h1>Findy. Join us!</h1>
            <input
                className="input-field"
                placeholder="Nickname"
                value={nickname}
                onChange={(e) => setNickname(e.target.value)}
            />
            <br />
            <input
                className="input-field"
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
            />
            {emailError && <p className="error-message">{emailError}</p>}
            <br />
            <input
                type="password"
                className="input-field"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />
            <br />
            <div className="interests-section">
                <input
                    className="input-field"
                    placeholder="Ваши интересы"
                    value={currentInterest}
                    onChange={(e) => setCurrentInterest(e.target.value)}
                    onKeyDown={addInterest}
                />
                <ul className="interests-list">
                    {interests.map((interest, index) => (
                        <li key={index} className="interest-item">
                            {interest}
                            <button
                                className="remove-button"
                                onClick={() => removeInterest(index)}
                            >
                                &#x2716;
                            </button>
                        </li>
                    ))}
                </ul>
            </div>
            <br />
            <input
                className="ToGoButton"
                type="button"
                value="Register"
                onClick={handleRegister}
            />
        </div>
    );
};

export { RegScreen };
