import React from 'react';

function NewEventPanelShow() {
    let div = document.getElementById('CreateEvent');
    if (div) {
        div.style.display = 'block';
        console.log(div.style.display);
    }
}

function NewEventPanelHide() {
    let div = document.getElementById('CreateEvent');
    if (div) {
        div.style.display = 'none';
    }
}

function NewEventAdd() {
    NewEventPanelHide();
    let name = document.getElementById('name_input').value;
    let host = 0; // loggeduser.nickname;
    let address = 0;
    let agemin = document.getElementById('agemin_input').value;
    let agemax = document.getElementById('agemax_input').value;
    let maxcount = document.getElementById('maxcount_input').value;
    let date = document.getElementById('date_input').value;
    console.log(name, host, address, agemin, agemax, maxcount, date);
    // events.append(new Event(name, host, address, agemin, agemax, maxcount, date));
}

const EventComponent = ({ event }) => (
    <div className="Event">
        <h2>{event.name}</h2>
        <p className="event-host">{event.host.nickname}</p>
        <h5 className="event-details">
            {event.address}, {event.date}, {event.agemin} - {event.agemax} лет
        </h5>
        <div className="event-tags">
            {event.tags.map((tag, index) => (
                <span key={index} className="event-tag">{tag}</span>
            ))}
        </div>
        <div style={{ display: 'inline-block' }}>
            <input type="button" value="Я приду" className="ToGoButton" />
        </div>
    </div>
);

const NewEvent = () => (
    <div id="CreateEvent" style={{ display: 'none' }}>
        <div>
            <h1 style={{ display: 'inline-block' }}>Создание нового события</h1>
            <input type="button" value="X" className="NegativeButton" onClick={NewEventPanelHide} />
        </div>
        <div style={{ display: 'inline-block', verticalAlign: 'top' }}>
            <input id="name_input" type="search" placeholder="Название" />
            <br />
            <input id="date_input" type="search" placeholder="Время сбора" />
            <br />
            <p>Точка сбора</p>
        </div>
        <div style={{ display: 'inline-block' }}>
            <p style={{ display: 'inline-block' }}>Мин. возраст</p>
            <input id="agemin_input" type="range" min="0" max="100" step="1" defaultValue="14" />
            <br />
            <p style={{ display: 'inline-block' }}>Макс. возраст</p>
            <input id="agemax_input" type="range" min="0" max="100" step="1" defaultValue="18" />
            <br />
            <p style={{ display: 'inline-block' }}>Макс. кол-во человек</p>
            <input id="maxcount_input" type="range" min="0" max="100" step="1" defaultValue="50" />
            <br />
            <input type="submit" value="Опубликовать" className="ToGoButton" onClick={NewEventAdd} />
        </div>
    </div>
);

export { EventComponent, NewEvent };