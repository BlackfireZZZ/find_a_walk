import React from 'react';
import './App.css';
import {Profile, loggedUser} from './components/Profile.jsx';
import YandexMap from './components/YandexMap.jsx';
import EventsList from "./components/EventsList";
import {Header} from "./components/Header";
import { BrowserRouter, Route } from 'react-router-dom';

import {RegScreen} from "./components/Register";
import {NewEvent} from "./components/NewEvent";
import {LoginScreen} from "./components/Login";

function App() {
    return (
        <BrowserRouter>
            <Route path="/" element={
                <div className="App">
                    <Header/>
                    <main>
                        <Profile user={loggedUser}/>
                        <div style={{display: 'inline-block', width: '78%'}}>
                        <YandexMap/>
                        <EventsList/>
                        </div>
                    </main>
                </div>} />
            <Route path="/register" element={<RegScreen/>} />
            <Route path="/login" element={<LoginScreen/>} />
        </BrowserRouter>


    );
}

export default App;
