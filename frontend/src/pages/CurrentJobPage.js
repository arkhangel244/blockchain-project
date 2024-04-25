// CurrentJobPage.js
import React, { useState, useEffect } from 'react';
import NavigationBar from './NavigationBar';
import '../css/CurrentJobPage.css'

const CurrentJobPage = () => {
  const [currentJob, setCurrentJob] = useState(null);

  useEffect(() => {
    const fetchCurrentJob = async () => {
      try {
        const response = await fetch("/api/current-job");
        if (!response.ok) {
          throw new Error('Failed to fetch current job');
        }
        const data = await response.json();
        setCurrentJob(data);
      } catch (error) {
        console.error('Error fetching current job:', error.message);
      }
    };

    fetchCurrentJob();
  }, []);

  if (!currentJob) {
    return <div>Loading...</div>;
  }

  return (
    <div style={{marginLeft: 20, marginTop: 20}}>
        <NavigationBar/>
        <div>
            <br/>
            <h1>{currentJob.title}</h1>
            <h2>{currentJob.company}</h2>
            <h3>Salary Transactions:</h3>
            <table className="table">
                <thead>
                <tr>
                    <th>ID</th>
                    <th>Date</th>
                    <th>Amount</th>
                </tr>
                </thead>
                <tbody>
                {currentJob.salaryTransactions.map(transaction => (
                    <tr key={transaction.id}>
                        <td>{transaction.id}</td>
                        <td>{transaction.date}</td>
                        <td>{transaction.amount}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    </div>
  );
};

export default CurrentJobPage;
