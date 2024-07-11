import React from 'react';
import './App.css';
import {Profile, loggedUser} from './components/Profile.jsx';
import YandexMap from './components/YandexMap.jsx';
import EventsList from "./components/EventsList";
import {Header} from "./components/Header";
import {RegScreen} from "./components/Register";

function App() {
    return (
        <div className="App">
            <Header/>
            {/*<main>*/}
            {/*    <Profile user={loggedUser}/>*/}
            {/*    <div style={{display: 'inline-block', width: '78%'}}>*/}
            {/*        <YandexMap/>*/}
            {/*        <EventsList/>*/}
            {/*    </div>*/}
            {/*</main>*/}
            <RegScreen/>
        </div>
    );
}

export default App;
