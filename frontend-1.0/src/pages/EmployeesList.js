import React, { useState, useEffect } from 'react';
import NavigationBar2 from './NavigationBar2';

const EmployeesList = () => {
  const [hiredEmployees, setHiredEmployees] = useState([]);

  useEffect(() => {
    const fetchHiredEmployees = async () => {
      try {
        const response = await fetch("/api/hired-employees");
        if (!response.ok) {
          throw new Error('Failed to fetch hired employees');
        }
        const data = await response.json();
        setHiredEmployees(data);
      } catch (error) {
        console.error('Error fetching hired employees:', error.message);
      }
    };

    fetchHiredEmployees();
  }, []);

  return (
    <div style={{marginLeft: 20, marginTop: 20}}>
        <NavigationBar2/>
        <h1>Hired Employees</h1>
        <ul>
            {hiredEmployees.map(employee => (
            <li key={employee.id}>
                <p>Name: {employee.name}</p>
                <p>Role: {employee.role}</p>
            </li>
            ))}
        </ul>
    </div>
  );
};

export default EmployeesList;
