import React, { useState, useEffect } from 'react';
import NavigationBar from './NavigationBar';

const JobsPage = () => {
  const [jobs, setJobs] = useState([]);

  useEffect(() => {
    const fetchJobs = async () => {
      try {
        const response = await fetch("/api/jobs");
        if (!response.ok) {
          throw new Error('Failed to fetch jobs');
        }
        const data = await response.json();
        setJobs(data);
      } catch (error) {
        console.error('Error fetching jobs:', error.message);
      }
    };

    fetchJobs();
  }, []);

  return (
    <div style={{marginLeft: 20, marginTop: 20}}>
      <NavigationBar/>
      <div>
        <br/>
        <h1>Available Jobs</h1>
        <br/>
        <div className="job-list">
          {jobs.map((job, index) => (
            <div className="job-card" key={index} style = {{marginBottom: 60}}>
              <h2>{job.title}</h2>
              <p><strong>Company:</strong> {job.company}</p>
              <p><strong>Salary:</strong> {job.salary}</p>
              <p><strong>Location:</strong> {job.place}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default JobsPage;
