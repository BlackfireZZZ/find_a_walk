class User {
    constructor(nickname, birthdate, contacts){
        this.nickname = nickname;
        this.birthdate = birthdate;
        this.contacts = contacts;
    }
}
class Event {
    constructor(name, host, address, agemin, agemax, date){
        this.name = name;
        this.host = host;
        this.address = address;
        this.agemin = agemin;
        this.agemax = agemax;
        this.date = date;
        this.count = 0;
    }
    join(){
        count = count + 1;
    }
    show(){
        let eventList = document.getElementById('CurrentEvents');
        let el = document.createElement('div');
        el.className = "Event";

        let Name = document.createElement('h2');
        Name.innerHTML = this.name;
        let Host = document.createElement('p');
        Host.innerHTML = this.host.nickname;
        Host.style = "margin-top: -10px;";
        let Info = document.createElement('h5');
        Info.innerHTML = `${this.address}, ${this.date}, ${this.agemin} - ${this.agemax} лет`;
        Info.style = "margin-top: -10px;";
        el.appendChild(Name);
        el.appendChild(Host);
        el.appendChild(Info);

        eventList.appendChild(el);
    }
}
users = [
    new User('Chinese developers Team', [1, 7, 2024], ['89456783542', 'nightmarefuel']),
];
loggeduser = users[0];
profilescreen = document.getElementById('Profile');
let UserName = document.getElementById('username');
let UserAge = document.getElementById('userage');
let UserContacts = document.getElementById('usercontacts');
loggeduser.contacts.forEach(element => {
    let newcontact = document.createElement('div');
    newcontact.style = "border-radius: 20px; background-color: grey";
    let text = document.createElement('h5');
    text.style = "color: white; margin-top: -10px";
    text.innerHTML = element;
    newcontact.appendChild(text);
    UserContacts.appendChild(newcontact);
})
UserName.innerHTML = loggeduser.nickname;
UserAge.innerHTML = new Date().getFullYear() - loggeduser.birthdate[2];


events = [
    new Event('Чилл без бухла', users[0], 'Станция Новокосино', 16, 19, '27.07.2024'),
    new Event('Чилл без бухла', users[0], 'Станция Новокосино', 16, 19, '27.07.2024')
];
events.forEach(element => {
    element.show();
});
