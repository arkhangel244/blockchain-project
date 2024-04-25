import React, { useState, useEffect } from 'react';
import NavigationBar from './NavigationBar';

const OffersPage = () => {

    const [offers, setOffers] = useState([]);

    useEffect(() => {
        const fetchOffers = async () => {
        try {
            const response = await fetch("/api/offers");
            if (!response.ok) {
            throw new Error('Failed to fetch offers');
            }
            const data = await response.json();
            setOffers(data);
        } catch (error) {
            console.error('Error fetching offers:', error.message);
        }
        };

        fetchOffers();
    }, []);

    return (
        <div style={{marginLeft: 20, marginTop: 20}}>
        <NavigationBar/>
        <div>
            <br/>
            <h1>Job offers</h1>
            <br/>
            <div className="offer-list">
            {offers.map((offer, index) => (
                <div className="offer-card" key={index} style = {{marginBottom: 60}}>
                <h2>{offer.title}</h2>
                <p><strong>Company:</strong> {offer.company}</p>
                <p><strong>Salary:</strong> {offer.salary}</p>
                <p><strong>Location:</strong> {offer.place}</p>
                </div>
            ))}
            </div>
        </div>
        </div>
    );
};

export default OffersPage;
