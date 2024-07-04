import React from 'react';

class Event {
    constructor(name, host, address, agemin, agemax, date) {
        this.name = name;
        this.host = host;
        this.address = address;
        this.agemin = agemin;
        this.agemax = agemax;
        this.date = date;
        this.count = 0;
    }

    join() {
        this.count += 1;
    }
}

const EventComponent = ({ event }) => (
    <div className="Event">
        <div style={{display: 'inline-block'}}>
            <h2>{event.name}</h2>
            <p style={{ marginTop: '-10px' }}>{event.host.nickname}</p>
            <h5 style={{ marginTop: '-10px' }}>
            {event.address}, {event.date}, {event.agemin} - {event.agemax} лет
            </h5>
        </div>
        <div style={{display: 'inline-block'}}>
            <input type="button" value="Let's go!" style={{backgroundColor: 'green', color: 'white'}}></input>
        </div>
    </div>
);

export { Event, EventComponent };
