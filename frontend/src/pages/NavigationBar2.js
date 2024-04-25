import React from 'react';
import { Link } from 'react-router-dom';
import '../css/NavigationBar.css';

const NavigationBar2 = () => {
  return (
    <div className="navbar">
      <Link to="/hire-new" className="nav-link">Hire</Link>
      <Link to="/employees" className="nav-link">Employees-list</Link>
    </div>
  );
};

export default NavigationBar2;