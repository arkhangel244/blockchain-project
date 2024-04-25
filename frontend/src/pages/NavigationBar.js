import React from 'react';
import { Link } from 'react-router-dom';
import '../css/NavigationBar.css';

const NavigationBar = () => {
  return (
    <div className="navbar">
      <Link to="/job-search" className="nav-link">Job Search</Link>
      <Link to="/offers" className="nav-link">Offers</Link>
      <Link to="/current-job" className="nav-link">Current Job</Link>
    </div>
  );
};

export default NavigationBar;