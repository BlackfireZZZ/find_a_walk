import React, {useRef} from "react";

export class EventObj {
    constructor(name, id, author_id, start_lon, start_lat, end_lon, end_lat, date, capacity, members_count, tags) {
        this.name = name;
        this.id = id;
        this.author_id = author_id;
        this.start_lon = start_lon;
        this.start_lat = start_lat;
        this.end_lon = end_lon;
        this.end_lat = end_lat;
        this.date = date;
        this.capacity = capacity;
        this.members_count = members_count;
        this.tags = tags;
    }

    join() {
        // Implementation of join method
    }
}

export const Event = ({event}) => {
    return (
        <div className="Event" ref={useRef('CurrentEvents')} style={{padding: '0 20px 0 20px'}}>
            <div style={{display: 'inline-block'}}>
                <h2>{event.name}</h2>
                <p style={{marginTop: '-10px'}}>{event.host}</p>
                <h5 style={{marginTop: '-10px'}}>
                    {event.date}
                </h5>
            </div>
            <div style={{display: 'inline-block', position: 'absolute', right: '0', marginRight: '20px'}}>
                <input type="button" value="Я приду" className="ToGoButton" onClick={event.join}></input>
                {console.log(event.capacity > 0)}
                {event.capacity > 0 ? <h5>{event.members_count}/{event.capacity}</h5> : <h5>---</h5>}
            </div>
        </div>
    )
}

const components = { Event, EventObj };

export default components;
