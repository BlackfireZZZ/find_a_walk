class User {
    constructor(nickname, birthdate, contacts){
        this.nickname = nickname;
        this.birthdate = birthdate;
        this.contacts = contacts;
    }
}


function ProfilePanelSwitch(){
    let div = document.getElementById('Profile');
    switch(div.style.display){
        case "none":
            div.style.display = "inline-block";
            break;
        default:
            div.style.display = "none";
            break;
    }
}

function ContactAppend(contacttext){
    if(contacttext){
        let newcontact = document.createElement('div');
        newcontact.className = "ContactElement";
        let text = document.createElement('h5');
        text.style = "color: white; margin-left: 5px";
        text.innerHTML = contacttext;
        newcontact.appendChild(text);
        UserContacts.prepend(newcontact);
        document.getElementById('contactinput').value = "";
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

UserName.innerHTML = loggeduser.nickname;
UserAge.innerHTML = new Date().getFullYear() - loggeduser.birthdate[2] + ' years old';
loggeduser.contacts.forEach(element => {
    ContactAppend(element);
})