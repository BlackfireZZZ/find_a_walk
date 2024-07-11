import React, { useState } from 'react';
import '../RegScreen.css';

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
            tags: interests
        };

        try {
            const response = await fetch('http://localhost/api/users', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(user)
            });

            if (response.ok) {
                console.log('User registered successfully');
                // Optionally, you can handle a successful registration here (e.g., redirect to another page)
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
