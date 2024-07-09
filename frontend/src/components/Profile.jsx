import React from 'react';
import './App.css';
import User from './components/User';
import {Profile, users, loggedUser} from './components/Profile.jsx';
import { Event, EventComponent, NewEvent, NewEventPanelShow, events } from './components/Event.jsx';
import YandexMap from './components/YandexMap.jsx';

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <h3 style={{display: 'inline-block'}}>Find the walk.</h3>
                <p style={{display: 'inline-block'}}>Powered by Chinese Developers</p>
                <input type="button" value="+ New event" onClick={NewEventPanelShow}></input>
            </header>
            <main>
                <Profile user={loggedUser} />
                <div style={{display: 'inline-block', width: '78%'}}>
                    <YandexMap />
                    <div id="CurrentEvents" style={{}}>
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
