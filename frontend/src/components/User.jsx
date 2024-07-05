export default class User {
    constructor(nickname, birthdate, contacts) {
        this.nickname = nickname;
        this.birthdate = birthdate;
        this.contacts = contacts;
    }

    getAge() {
        return new Date().getFullYear() - this.birthdate[2];
    }
}
