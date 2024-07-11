import React from 'react'
import {Profile, loggedUser} from './Profile.jsx';
import YandexMap from './YandexMap.jsx';
import EventsList from "./EventsList.jsx";
import {Header} from "./Header.jsx";

const MainScreen = () => (
    <div className="App">
        <Header/>
        <main>
            <Profile user={loggedUser}/>
            <div style={{display: 'inline-block', width: '78%'}}>
                <YandexMap/>
                <EventsList/>
            </div>
        </main>
    </div>
)
export default MainScreen