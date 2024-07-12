import {EventObj, Event} from "./Event";
import config from '../config';
const EventsList = () => {


    function NewEventFromJson(json) {
        return new EventObj(
            json['name'],
            json['author_id'],
            json['start_longitude'],
            json['start_latitude'],
            json['end_longitude'],
            json['end_latitude'],
            Date.parse(json['date']),
            json['capacity'],
            json['members_count'],
            json['tags']
        )
    }

    const getEvents = () => {
        try {
            let xhr = new XMLHttpRequest();
            let url = config.Host_url + 'events';
            xhr.open("GET", url, true);
            xhr.send();
            return JSON.parse(xhr.responseText).map((event) => NewEventFromJson(event));
        } catch (error) {
            console.error('Error fetching events:', error);
            return [];
        }
    }
    let events = getEvents();
    return (
        <div id="CurrentEvents" style={{}}>
            {events.length > 0 ? 
                events.map((event, index) => (
                    <Event event={event}/>
                )) :
                <h2 style={{textAlign: 'center'}}>
                    Здесь ничего нет, будь первым
                </h2>
            }
        </div>
    );
};

export default EventsList;