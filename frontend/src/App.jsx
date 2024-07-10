import React from 'react';
import './App.css';
import {Profile, loggedUser} from './components/Profile.jsx';
import YandexMap from './components/YandexMap.jsx';
import EventsList from "./components/EventsList";

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <h3 style={{display: 'inline-block'}}>Find the walk.</h3>
                <p style={{display: 'inline-block'}}>Powered by Chinese Developers</p>
                <input type="button" value="+ New event" onClick={console.log}></input>
            </header>
            <main>
                <Profile user={loggedUser}/>
                <div style={{display: 'inline-block', width: '78%'}}>
                    <YandexMap/>
                    <EventsList/>
                </div>
            </main>
        </div>
    );
}

export default App;
