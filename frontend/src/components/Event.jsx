import React from 'react';

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
    </div>
);

export { EventComponent };
