import React from 'react';
import { Menu } from 'antd';

function NavBar() { 
    return (
        <Menu className="site-menu-backgroud" mode="horizontal" style={{justifyContent: "center"}}>
            <Menu.Item key="home">Home</Menu.Item>
            <Menu.Item key="shipping"><a href = "/Shipping">Shipping</a></Menu.Item>
            <Menu.Item key="tracking">Tracking</Menu.Item>
            <Menu.Item key="about">Who We Are</Menu.Item>
            <Menu.Item key="profile">My Profile</Menu.Item>
        </Menu>
    );
}

export default NavBar;