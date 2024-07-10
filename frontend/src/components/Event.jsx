import React, { useRef } from 'react';
import {users, loggedUser} from './Profile.jsx'

class Event {
    constructor(name, host, address, coords, agemin, agemax, date, maxcount) {
        this.name = name;
        this.host = host;
        this.coords = coords;
        this.radius = 1000;
        this.address = address;
        this.agemin = agemin;
        this.agemax = agemax;
        this.date = date;
        this.count = 0;
        this.maxcount = maxcount;
        this.memberslist = [];
    }

    join() {
        
    }
}

let events = [
    new Event(
        'Чилл без бухла', 
        users[0].nickname, 
        'Станция Новокосино', [55, 37],
        19, 27, '27.07.2024', 5),
    new Event(
        'ААА помогите с докером', 
        users[1].nickname, 
        'НИУ ВШЭ, Покровский бульвар 11', [56, 38],
        16, 19, '11.07.2024', 0)
];

const NewEventPanelShow = () => {
    let div = document.getElementById('CreateEvent');
    div.style.display = "block";
    console.log(div.style.display);
}
const NewEventPanelHide = () => {
    let div = document.getElementById('CreateEvent');
    div.style.display = "none";
}

const EventComponent = ({ event }) => (
    <div className="Event" ref={useRef('CurrentEvents')} style={{padding: '0 20px 0 20px'}}>
        <div style={{display: 'inline-block'}}>
            <h2>{event.name}</h2>
            <p style={{ marginTop: '-10px' }}>{event.host}</p>
            <h5 style={{ marginTop: '-10px' }}>
                {event.address}, {event.date}, {event.agemin} - {event.agemax} лет
            </h5>
        </div>
        <div style={{display: 'inline-block', position: 'absolute', right: '0', marginRight: '20px'}}>
            <input type="button" value="Я приду" class="ToGoButton" onClick={event.join()}></input>
            {console.log(event.maxcount > 0)}
            {event.maxcount > 0 ? <h5>{event.count}/{event.maxcount}</h5> : <h5>---</h5>}
            
        </div>
    </div>
);
const NewEventAdd = () => {
    let name = document.getElementById('name_input').value;
    let host = loggedUser.nickname;
    let address = 0;
    let coords = [57, 62];
    let agemin = document.getElementById('agemin_input').value;
    let agemax = document.getElementById('agemax_input').value;
    let maxcount = document.getElementById('maxcount_input').value;
    let date = document.getElementById('date_input').value;
    
    let xhr = new XMLHttpRequest();
    let url = 'http://localhost/api/events';
    xhr.open("POST", url, true);
    let data = JSON.stringify({
        'name': name,
        'author_id': "17fd3c37-cdfd-4170-b7c0-2d6f640c0b8d",
        'agemin': agemin,
        'agemax': agemax,
        'maxcount': maxcount,

    });
    xhr.send(data);
}
const NewEvent = () => (
    <div id="CreateEvent" style={{ display: 'none' }}>
        <div>
            <h1 style={{display: 'inline-block'}}>Создание нового события</h1>
            <input type="button" value="X" className='NegativeButton' onClick={NewEventPanelHide}></input>
        </div>
        <div style={{ display: 'inline-block', verticalAlign: 'top' }}>
            <input id="name_input" type="search" placeholder="Название" />
            <br />
            <input id="date_input" type="search" placeholder="Время сбора" />
            <br />
            <p>Точка сбора</p>
        </div>
        <div style={{display: 'inline-block'}}>
            <p style={{display: "inline-block"}}>Мин. возраст</p>
            <input id="agemin_input" type="placeholder" placeholder="Мин. возраст"></input>
            <br></br>
            <p style={{display: "inline-block"}}>Макс. возраст</p>
            <input id="agemax_input" type="placeholder" placeholder="Макс. возраст"></input>
            <br></br>
            <p style={{display: "inline-block"}}>Макс. кол-во человек</p>
            <input id="maxcount_input" type="placeholder" placeholder="Макс. кол-во участников"></input>
            <br></br>
            <input type="submit" placeholder='Опубликовать' className='ToGoButton' onClick={NewEventAdd}></input>
        </div>
    </div>
);

export { Event, EventComponent, NewEvent, NewEventPanelShow, events };
