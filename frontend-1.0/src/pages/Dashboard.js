import React from 'react';
import { useEffect, useState } from 'react'

function Dashboard() {

    const [backendData, setBackendData] = useState([{}])

    useEffect(() => {
        fetch("/api").then(
        response => response.json()
        ).then(
        data => {
            setBackendData(data)
        }
        )
    }, [])


    return (
        <div>hi</div>
    )
}

export default Dashboard;