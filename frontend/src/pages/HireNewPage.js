import React, { useState, useEffect } from 'react';
import NavigationBar2 from './NavigationBar2';

const HireNewPage = () => {
  const [employees, setEmployees] = useState([]);

  useEffect(() => {
    const fetchEmployees = async () => {
      try {
        const response = await fetch("/api/ready-to-hire");
        if (!response.ok) {
          throw new Error('Failed to fetch employees');
        }
        const data = await response.json();
        setEmployees(data);
      } catch (error) {
        console.error('Error fetching employees:', error.message);
      }
    };

    fetchEmployees();
  }, []);

  const hireEmployee = async (employeeId) => {
    try {
        const response = await fetch(`/api/hire/${employeeId}`, {
            method: 'POST',
            headers: {
            'Content-Type': 'application/json'
            },
        });
        if (!response.ok) {
            throw new Error('Failed to hire employee');
        }
    } 
    catch (error) {
        console.error('Error hiring employee:', error.message);
    }
    console.log(`Hiring employee with ID ${employeeId}`);
  };

  return (
    <div style={{marginLeft: 20, marginTop: 20}}>
        <NavigationBar2/>
        <h1>Employees Ready to be Hired</h1>
        <ul>
            {employees.map(employee => (
            <li key={employee.id}>
                <div>
                <p>Name: {employee.name}</p>
                <p>Experience: {employee.experience}</p>
                <p>Role: {employee.role}</p>
                </div>
                <button onClick={() => hireEmployee(employee.id)}>Hire</button>
            </li>
            ))}
        </ul>
    </div>
  );
};

export default HireNewPage;
