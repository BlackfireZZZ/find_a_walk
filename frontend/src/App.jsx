import React, { useState, useEffect } from 'react';
import './App.css';
import User from './components/User';
import Profile from './components/Profile.jsx';
import { Event, EventComponent, NewEvent } from './components/Event.jsx';
import YandexMap from './components/YandexMap.jsx';

const users = [
    new User('Chinese developers Team', [1, 7, 2024], [
        '89456783542',
        'nightmarefuel'
    ])
];

const loggedUser = users[0];
function App() {
    const [events, setEvents] = useState([]);
    const [bounds, setBounds] = useState({ lon1: 0, lat1: 0, lon2: 0, lat2: 0 });

    useEffect(() => {
        if (bounds.lon1 !== 0 && bounds.lon2 !== 0 && bounds.lat1 !== 0 && bounds.lat2 !== 0) {
            const { lon1, lat1, lon2, lat2 } = bounds;
            fetch(`http://localhost/api/v1/events/?lon1=${lon1}&lat1=${lat1}&lon2=${lon2}&lat2=${lat2}`)
                .then(response => response.json())
                .then(data => {
                    const eventsData = data.map(event => ({
                        id: event.id,
                        name: event.name,
                        host: loggedUser,  // Assuming loggedUser is the host for all events
                        address: `${event.start_latitude}, ${event.start_longitude}`,
                        agemin: event.agemin,
                        agemax: event.agemax,
                        date: new Date(event.date).toLocaleDateString(),
                        tags: event.tags.map(tag => tag.name)
                    }));
                    setEvents(eventsData);
                })
                .catch(error => console.error('Error fetching events:', error));
        }
    }, [bounds]);

    return (
        <div className="App">
            <header className="App-header">
            </header>
            <main>
                <Profile user={loggedUser} />
                <div className="App-header">
                    <YandexMap onBoundsChange={setBounds} />
                    <div id="CurrentEvents">
                        {events.map((event, index) => (
                            <EventComponent key={index} event={event} />
                        ))}
                    </div>
                    <NewEvent />
                </div>
            </main>
        </div>
    );
}

export default App;
