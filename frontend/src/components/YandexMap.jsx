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
         loadScript("https:api-maps.yandex.ru/2.1/?apikey=6997c194-93fd-44c8-89ce-8639d5bcd0c1&lang=ru_RU")
             .then(() => {
                 window.ymaps.ready(() => {

                     const map = new window.ymaps.Map('map', {
                         center: [55.7558, 37.6176],
                         zoom: 10
                     });

                     var dots = [];
                     let events = [[55.76, 37.64]];
                     events.forEach(event => {
                         dots.push(
                             new window.ymaps.Placemark(
                             [event[0], event[1]], 
                             {
                                balloonContentHeader: 'Чилл без бухла',
                                balloonContentBody: '<p>Курский вокзал, 13.07.2024</p><input type="button" class="ToGoButton" value="Я приду!"></input>',
                                balloonContentFooter: '16-19 лет, до 8 человек',
                                hintContent: 'Чилл без бухла', })
                         )
                     })
                     dots.forEach(dot => {
                         map.geoObjects.add(dot);
                     })

                 });

             })
             .catch((error) => console.error(error));
     }, []);
     return (
         <div id="map"></div>
     );
 };

 export default YandexMap;
