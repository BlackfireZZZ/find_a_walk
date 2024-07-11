import React, { useState, useRef } from 'react';
import DateTimePicker from 'react-datetime-picker';
import { users, loggedUser } from './Profile.jsx';
import { EventObj } from './Event.jsx';

const NewEventAdd = async (nameRef, date, addressRef, ageMinRef, ageMaxRef, maxCountRef, setCords) => {
    const name = nameRef.current.value;
    const host = loggedUser.nickname;
    const address = addressRef.current.value;
    const agemin = ageMinRef.current.value;
    const agemax = ageMaxRef.current.value;
    const maxcount = maxCountRef.current.value;

    // Fetch coordinates from Yandex Geocoder
    const response = await fetch(`https://geocode-maps.yandex.ru/1.x/?apikey=6997c194-93fd-44c8-89ce-8639d5bcd0c1&geocode=${address}&format=json`);
    const data = await response.json();
    const cords = data.response.GeoObjectCollection.featureMember[0].GeoObject.Point.pos.split(' ');
    setCords(cords);

    // Create event object and send data to the server
    const event = new EventObj(
        name,
        null,
        '17fd3c37-cdfd-4170-b7c0-2d6f640c0b8d',
        parseFloat(cords[0]),
        parseFloat(cords[1]),
        null,
        null,
        date.toISOString(),
        maxcount,
        0,
        []
    );

    const xhr = new XMLHttpRequest();
    const url = 'http://localhost/api/events';
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.send(JSON.stringify(event));
    console.log('Event created:', event);
}

const NewEvent = () => {
    const nameRef = useRef(null);
    const addressRef = useRef(null);
    const ageMinRef = useRef(null);
    const ageMaxRef = useRef(null);
    const maxCountRef = useRef(null);
    const [cords, setCords] = useState([]);
    const [date, setDate] = useState(new Date());
    const [suggestions, setSuggestions] = useState([]);
    const [isDatePickerOpen, setIsDatePickerOpen] = useState(false);

    const handleAddressChange = async (e) => {
        const address = e.target.value;
        const response = await fetch(`https://geocode-maps.yandex.ru/1.x/?apikey=6997c194-93fd-44c8-89ce-8639d5bcd0c1&geocode=${address}&format=json`);
        const data = await response.json();
        const suggestions = data.response.GeoObjectCollection.featureMember.map(member => member.GeoObject.name);
        setSuggestions(suggestions);
    };

    const handleSuggestionClick = (suggestion) => {
        addressRef.current.value = suggestion;
        setSuggestions([]);
    };

    const handleDateChange = (date) => {
        setDate(date);
        setIsDatePickerOpen(false); // Close the date picker after selecting a date
    };

    return (
        <div id="CreateEvent" style={{ padding: '20px', border: '1px solid #000' }}>
            <div>
                <h1 style={{ display: 'inline-block' }}>Создание нового события</h1>
            </div>
            <div style={{ display: 'inline-block', verticalAlign: 'top' }}>
                <input id="name_input" type="search" placeholder="Название" ref={nameRef} />
                <br />
                <div>
                    <DateTimePicker
                        onChange={handleDateChange}
                        value={date}
                        isOpen={isDatePickerOpen}
                        onCalendarClose={() => setIsDatePickerOpen(false)}
                        onCalendarOpen={() => setIsDatePickerOpen(true)}
                    />
                </div>
                <br />
                <input id="address_input" type="search" placeholder="Точка сбора" ref={addressRef} onChange={handleAddressChange} />
                {suggestions.length > 0 && (
                    <ul className="suggestions">
                        {suggestions.map((suggestion, index) => (
                            <li key={index} onClick={() => handleSuggestionClick(suggestion)}>{suggestion}</li>
                        ))}
                    </ul>
                )}
            </div>
            <div style={{ display: 'inline-block' }}>
                <p style={{ display: "inline-block" }}>Мин. возраст</p>
                <input id="agemin_input" type="text" placeholder="Мин. возраст" ref={ageMinRef} />
                <br />
                <p style={{ display: "inline-block" }}>Макс. возраст</p>
                <input id="agemax_input" type="text" placeholder="Макс. возраст" ref={ageMaxRef} />
                <br />
                <p style={{ display: "inline-block" }}>Макс. кол-во человек</p>
                <input id="maxcount_input" type="text" placeholder="Макс. кол-во участников" ref={maxCountRef} />
                <br />
                <input type="submit" value='Опубликовать' className='ToGoButton' onClick={() => NewEventAdd(nameRef, date, addressRef, ageMinRef, ageMaxRef, maxCountRef, setCords)} />
            </div>
        </div>
    );
};

export { NewEvent };
