class Event {
    constructor(name, host, address, agemin, agemax, date, maxcount){
        this.name = name;
        this.host = host;
        this.address = address;
        this.agemin = agemin;
        this.agemax = agemax;
        this.date = date;
        this.count = 0;
        this.maxcount = maxcount;
    }
    join(){
        this.count++;
    }
    show(){
        let eventList = document.getElementById('CurrentEvents');
        let el = document.createElement('div');
        el.className = "Event";

        let leftside = document.createElement('div');
        leftside.style = "display: inline-block";
        let rightside = document.createElement('div');
        rightside.style = "display: inline-block; position: absolute; right: 15px; margin: 10px;";

        let Name = document.createElement('h2');
        Name.innerHTML = this.name;

        let Host = document.createElement('p');
        Host.innerHTML = this.host.nickname;
        Host.style = "margin-top: -10px;";

        let Info = document.createElement('h5');
        Info.innerHTML = `${this.address}, ${this.date}, ${this.agemin} - ${this.agemax} лет`;
        Info.style = "margin-top: -10px";

        let Button = document.createElement('input');
        Button.type = "button";
        Button.className = "ToGoButton";
        Button.value = "Принять";
        Button.onclick = "this.join()";

        let Counter = document.createElement('h4');
        Counter.innerHTML = `${this.count}/${this.maxcount}`;

        leftside.appendChild(Name);
        leftside.appendChild(Host);
        leftside.appendChild(Info);
        rightside.appendChild(Button);
        rightside.appendChild(Counter);

        el.appendChild(leftside);
        el.appendChild(rightside);

        eventList.appendChild(el);
    }
}
function CurrentEventsPanelSwitch(){
    let div = document.getElementById('CurrentEvents');
    switch(div.style.display){
        case "none":
            div.style.display = "block";
            break;
        default:
            div.style.display = "none";
            break;
    }
}
let creatingnewevent = false;
function CreateNewEvent() {
    d = ['none', 'block'];
    let newevent = document.getElementById('CreateEvent');
    let currentevents = document.getElementById('CurrentEvents');
    creatingnewevent = !creatingnewevent;
    newevent.style.display = d[Number(creatingnewevent)];
    currentevents.style.display = d[Number(!creatingnewevent)];
}
events = [
    new Event('Чилл без бухла', users[0], 'Станция Новокосино', 16, 19, '27.07.2024', 5),
    new Event('Помогите с фронтендом', users[0], 'НИУ ВШЭ, Покровский бульвар 11', 14, 18, '11.07.2024', 15)
];
events.forEach(element => {
    element.show();
});