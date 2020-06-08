import React from "react";
import Navbar from "react-bootstrap/Navbar";

const Header = () =>(
    <div className="Header">
        <Navbar bg="primary" variant="dark">
            <Navbar.Brand href="#home">Real Time Chat App</Navbar.Brand>

        </Navbar>
    </div>
)
export default Header