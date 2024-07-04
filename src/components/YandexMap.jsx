import React, { useEffect } from 'react';

const YandexMap = () => {
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
                    const map = new window.ymaps.Map('map', {
                        center: [55.7558, 37.6176],
                        zoom: 10
                    });
                });
            })
            .catch((error) => console.error(error));
    }, []);

    return (
        <div id="map" style={{ display: 'inline-block', width: '78%', height: '100vh' }}></div>
    );
};

export default YandexMap;
