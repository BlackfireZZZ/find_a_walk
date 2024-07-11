import React from "react";
import '../Menu.css'

const Menu = ({header, items, active, setActive}) => {
    return (
        <div className={active ? 'menu active' : 'menu '} onClick={() => setActive(false)}>
            <div className="menu-content">
                <div className="menu-header">{header}</div>
                <a>почта - dinosaur.rawr@gmail.com</a>
                <ul>
                <a>Увлечения:</a>
                    {items.map(item =>
                        <li>
                            <a class="items" href={item.href}>{item.value}</a>
                        </li>
                    )}
                </ul>
            </div>
        </div>
    )
}
export default Menu