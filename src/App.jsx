import React from 'react';
import './App.css';
import User from './components/User';
import Profile from './components/Profile.jsx';
import { Event, EventComponent } from './components/Event.jsx';
import YandexMap from './components/YandexMap.jsx';

const users = [
    new User('Chinese developers Team', [1, 7, 2024], [
        '89456783542',
        'nightmarefuel'
    ])
];

const loggedUser = users[0];

const events = [
    new Event('Чилл без бухла', users[0], 'Станция Новокосино', 16, 19, '27.07.2024'),
    new Event('Чилл без бухла', users[0], 'Станция Новокосино', 16, 19, '27.07.2024')
];

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <h3>Events and User Profile</h3>
                
            </header>
            <main>
                <Profile user={loggedUser} />
                <YandexMap />
                <div id="CurrentEvents" style={{}}>
                    {events.map((event, index) => (
                        <EventComponent key={index} event={event} />
                    ))}
                </div>
                
            </main>
        </div>
    );
}

export default App;