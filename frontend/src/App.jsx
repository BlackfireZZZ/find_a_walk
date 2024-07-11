import React from 'react';
import './App.css';
import {RegScreen} from "./components/Register";
import {LoginScreen} from "./components/Login";
import MainScreen from "./components/MainPage"
import {NewEvent} from "./components/NewEvent";
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';


function App() {
    return (
        <Router>
            <Routes>
                <Route path="/main" element={<MainScreen/>}/>
                <Route path="/register" element={<RegScreen/>}/>
                <Route path="/login" element={<LoginScreen/>}/>
                <Route path="/newEvent" element={<NewEvent/>}/>
            </Routes>
        </Router>
    );
}

export default App;