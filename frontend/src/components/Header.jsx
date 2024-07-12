import React from "react";

export const Header = () => {
    return (
    <header className="App-header">
        <h3 style={{display: 'inline-block'}}>Findy.</h3>
        <p style={{display: 'inline-block'}}>Powered by Chinese Developers</p>
        <input type="button" value="+ New event" onClick={() => window.location.href = 'newEvent'}></input>
    </header>
    )
}