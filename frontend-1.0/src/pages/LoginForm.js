import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../css/LoginForm.css'

const LoginForm = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [choice, setChoice] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    console.log(JSON.stringify({ username, password }));
    if(choice === 'employee'){
        try {
            const response = await fetch('/employeeLogin', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
              },
              body: JSON.stringify({ username, password }),
            });
            const data = await response.json();
            console.log(data);
            if (data.success) {
              navigate('/job-search');
            }
          } catch (error) {
            console.error('Login failed:', error.message);
          }
    }
    else{
        try {
            const response = await fetch('/employerLogin', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
              },
              body: JSON.stringify({ username, password }),
            });
            const data = await response.json();
            console.log(data);
            if (data.success) {
              navigate('/hire-new');
            }
          } catch (error) {
            console.error('Login failed:', error.message);
          }
    }
    
  };

  return (
    <div className="login-container">
      <form onSubmit={handleSubmit} className="login-form">
        <input type="text" placeholder="Username" value={username} onChange={e => setUsername(e.target.value)} />
        <input type="password" placeholder="Password" value={password} onChange={e => setPassword(e.target.value)} />
        <div className="radio-group">
            <label>
            <input
                type="radio"
                name="choice"
                value="employee"
                checked={choice === 'employee'}
                onChange={() => setChoice('employee')}
            />
            <span>Employee</span>
            </label>
            <label>
            <input
                type="radio"
                name="choice"
                value="employer"
                checked={choice === 'employer'}
                onChange={() => setChoice('employer')}
            />
            <span>Employer</span>
            </label>
        </div>
        <button type="submit">Login</button>
      </form>
    </div>
  );
};

export default LoginForm;
