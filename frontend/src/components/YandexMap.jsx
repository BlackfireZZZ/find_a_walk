import React, { useEffect, useRef } from 'react';
import { events } from './Event.jsx';

const YandexMap = () => {
    const mapRef = useRef(null);

    useEffect(() => {
        const loadScript = (url) => {
            return new Promise((resolve, reject) => {
                const script = document.createElement('script');
                script.src = url;
                script.async = true;
                script.onload = () => resolve();
                script.onerror = () => reject(new Error(`Failed to load script ${url}`));
                document.body.appendChild(script);
            });
        };

        loadScript("https://api-maps.yandex.ru/2.1/?apikey=6997c194-93fd-44c8-89ce-8639d5bcd0c1&lang=ru_RU")
            .then(() => {
                window.ymaps.ready(() => {
                    const map = new window.ymaps.Map(mapRef.current, {
                        center: [55.7558, 37.6176],
                        zoom: 10
                    });

                    events.forEach(event => {
                        const placemark = new window.ymaps.Placemark(
                            [event.coords[0], event.coords[1]],
                            { balloonContent: event.name }
                        );
                        map.geoObjects.add(placemark);
                    });
                });
            })
            .catch((error) => console.error(error));
    }, []);

    return (
        <div ref={mapRef} style={{ position: 'absolute', top: '10vh', width: '78%', height: '480px' }}></div>
    );
};

export default YandexMap;
