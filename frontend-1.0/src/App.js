import React from 'react';
import { useEffect, useState } from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LoginForm from './pages/LoginForm';
import Dashboard from './pages/Dashboard';
import JobsPage from './pages/JobsPage';
import OffersPage from './pages/OffersPage';
import CurrentJobPage from './pages/CurrentJobPage';
import HireNewPage from './pages/HireNewPage';
import EmployeesList from './pages/EmployeesList';


const App = () => {

  return (
    <Router>
      <Routes>
        <Route exact path="/" element={<LoginForm />} />
        <Route path="/job-search" element={<JobsPage />} />
        <Route path="/offers" element={<OffersPage />} />
        <Route path="/current-job" element={<CurrentJobPage />} />
        <Route path="/hire-new" element={<HireNewPage />} />
        <Route path="/employees" element={<EmployeesList />} />
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </Router>
  );
};

export default App;
