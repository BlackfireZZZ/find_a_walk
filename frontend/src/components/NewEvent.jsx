import React, { useState, useRef } from 'react';
import config from "../config";

const NewEventAdd = async (nameRef, date, addressRef, maxCountRef, setCords, interests) => {
    const name = nameRef.current.value;
    const address = addressRef.current.value;
    const maxcount = maxCountRef.current.value;
    const tags = interests;

    // Fetch coordinates from Yandex Geocoder
    const geocodeResponse = await fetch(`https://geocode-maps.yandex.ru/1.x/?apikey=6997c194-93fd-44c8-89ce-8639d5bcd0c1&geocode=${address}&format=json`);
    const geocodeData = await geocodeResponse.json();
    const cords = geocodeData.response.GeoObjectCollection.featureMember[0].GeoObject.Point.pos.split(' ');
    setCords(cords);

    // Create event object and send data to the server
    const event = {
        name: name,
        start_longitude: parseFloat(cords[0]),
        start_latitude: parseFloat(cords[1]),
        end_longitude: null,
        end_latitude: null,
        date: new Date(date).toISOString(),
        capacity: parseInt(maxcount),
        tags: tags,
    };

    const serverResponse = await fetch(config.Host_url + '/events', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(event),
    });

    if (serverResponse.ok) {
        const createdEvent = await serverResponse.json();
        console.log('Event created:', createdEvent);
        // Handle the created event, e.g., navigate to event details or show a success message
    } else {
        console.error('Failed to create event');
        // Handle the error, e.g., show an error message
    }
}

const NewEvent = () => {
    const nameRef = useRef(null);
    const addressRef = useRef(null);
    const maxCountRef = useRef(null);
    const [cords, setCords] = useState([]);
    const [date, setDate] = useState(new Date());
    const [suggestions, setSuggestions] = useState([]);
    const [isDatePickerOpen, setIsDatePickerOpen] = useState(false);
    const [interests, setInterests] = useState([]);
    const [currentInterest, setCurrentInterest] = useState('');

    const addInterest = (e) => {
        if (e.key === 'Enter' && currentInterest) {
            setInterests([...interests, currentInterest]);
            setCurrentInterest('');
        }
    };

    const removeInterest = (index) => {
        setInterests(interests.filter((_, i) => i !== index));
    };

    const handleAddressChange = async (e) => {
        if (document.getElementById('address_input').value.length > 0) {
            const address = e.target.value;
            const geocodeResponse = await fetch(`https://geocode-maps.yandex.ru/1.x/?apikey=6997c194-93fd-44c8-89ce-8639d5bcd0c1&geocode=${address}&format=json`);
            const geocodeData = await geocodeResponse.json();
            const suggestions = geocodeData.response.GeoObjectCollection.featureMember.map(member => member.GeoObject.name);
            setSuggestions(suggestions);
        }
    };

    const handleDateChange = (e) => {
        setDate(e.target.value);
    };

    const handleSuggestionClick = (suggestion) => {
        addressRef.current.value = suggestion;
        setSuggestions([]);
    };

    const handleSubmit = async () => {
        await NewEventAdd(nameRef, date, addressRef, maxCountRef, setCords, interests);
    };

    return (
        <div id="CreateEvent">
            <div>
                <h1 style={{display: 'inline-block'}}>Создание нового события</h1>
                <h1 style={{display: 'inline-block'}} className='NegativeButton'
                    onClick={() => window.location.href = '/'}>X</h1>
            </div>
            <div style={{display: 'inline-block', verticalAlign: 'top'}}>
                <input id="name_input" type="text" placeholder="Название" ref={nameRef}/>
                <br/>
                <div>
                    <input id="date_input" style={{width: "50%"}} type="date" onChange={handleDateChange}></input>
                    <input id="time_input" style={{width: "33%"}} type="time"></input>
                </div>
                <input id="address_input" type="search" placeholder="Точка сбора" ref={addressRef}
                       onChange={handleAddressChange}/>
                {suggestions.length > 0 && (
                    <div className="suggestions">
                        {suggestions.map((suggestion, index) => (
                            <div className='GeoSuggest' key={index}
                                 onClick={() => handleSuggestionClick(suggestion)}>{suggestion}</div>
                        ))}
                    </div>
                )}
            </div>
            <div style={{display: 'inline-block'}}>
                <input id="maxcount_input" type="number" placeholder="Макс. кол-во участников" ref={maxCountRef}/>
            </div>
            <br/>
            <div className="interests-section">
                <input
                    className="input-field"
                    placeholder="Ваши интересы"
                    value={currentInterest}
                    onChange={(e) => setCurrentInterest(e.target.value)}
                    onKeyDown={addInterest}
                />
                <ul className="interests-list">
                    {interests.map((interest, index) => (
                        <li key={index} className="interest-item">
                            {interest}
                            <button
                                className="remove-button"
                                onClick={() => removeInterest(index)}
                            >
                                &#x2716;
                            </button>
                        </li>
                    ))}
                </ul>
            </div>
            <input type="submit" value='Опубликовать'
                   style={{width: '100%'}} className='ToGoButton'
                   onClick={handleSubmit}/>
        </div>
    );
};

export { NewEvent };
