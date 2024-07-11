import React from 'react';
import './App.css';
import './index.css'
import {Profile, loggedUser} from './components/Profile.jsx';
import YandexMap from './components/YandexMap.jsx';
import EventsList from "./components/EventsList";
import {Header} from "./components/Header";
import {RegScreen} from "./components/Register";
import {LoginScreen} from "./components/Login";
import {NewEvent} from "./components/NewEvent";
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';



function App() {
    return (
        <Router>
            <Routes>
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
                    </div>}/>
                <Route path="/register" element={<RegScreen/>}/>
                <Route path="/login" element={<LoginScreen/>}/>
                <Route path="/newEvent" element={<NewEvent/>}/>
            </Routes>
        </Router>

    );
}

export default App;
