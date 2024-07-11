import React, { useState } from 'react'
// import {Profile, loggedUser} from './Profile.jsx';
import YandexMap from './YandexMap.jsx';
import EventsList from "./EventsList.jsx";
import Menu from "./Menu.jsx"
import {Header} from "./Header.jsx";

const MainScreen = () => {
    const [menuActive, setMenuActive] = useState(false)
    const items = [{value: "люблю кататься"}, {value: "обожаю начос"}, {value: "да не умер я в конце драйва"},]
    return (
    <div className="App">


        <Header/>
        <main>
        <input type="button" value="Profile" class="menu-button" onClick={() => setMenuActive(!menuActive)}></input>  

            {/* <Profile user={loggedUser}/> */}
            <div style={{display: 'inline-block', width: '100%'}}>
                
                <YandexMap/>

                <EventsList/>
            </div>
        </main>
        <Menu  header={"Райан Гослинг"} items={items} active={menuActive} setActive={setMenuActive}/>
    </div>
    )
}
export default MainScreen