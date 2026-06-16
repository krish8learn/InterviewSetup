'use client';

import { useEffect, useState } from 'react';

export default function Home() {
  const [message, setMessage] = useState(null);
  const [mongoHealth, setMongoHealth] = useState(null);
  const [mysqlHealth, setMysqlHealth] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    Promise.all([
      fetch(`${process.env.NEXT_PUBLIC_API_URL}/`)
        .then((res) => {
          if (!res.ok) throw new Error(`HTTP ${res.status}`);
          return res.json();
        })
        .then((data) => setMessage(data.message ?? JSON.stringify(data)))
        .catch((err) => setError(err.message)),
      fetch(`${process.env.NEXT_PUBLIC_API_URL}/mongo-health`)
        .then((res) => {
          if (!res.ok) throw new Error(`HTTP ${res.status}`);
          return res.json();
        })
        .then((data) => setMongoHealth(data.status ?? JSON.stringify(data)))
        .catch((err) => setMongoHealth(`Error: ${err.message}`)),
      fetch(`${process.env.NEXT_PUBLIC_API_URL}/mysql-health`)
        .then((res) => {
          if (!res.ok) throw new Error(`HTTP ${res.status}`);
          return res.json();
        })
        .then((data) => setMysqlHealth(data.status ?? JSON.stringify(data)))
        .catch((err) => setMysqlHealth(`Error: ${err.message}`)),
    ]).finally(() => setLoading(false));
  }, []);

  return (
    <main style={{ fontFamily: 'sans-serif', padding: '2rem' }}>
      <h1>Hello World</h1>
      {loading && <p>Loading from backend…</p>}
      {error && <p style={{ color: 'red' }}>Backend error: {error}</p>}
      {message && <p>Backend says: {message}</p>}
      <h2>Database Health</h2>
      {mongoHealth && (
        <p style={{ color: mongoHealth.includes('connected') ? 'green' : 'red' }}>
          MongoDB: {mongoHealth}
        </p>
      )}
      {mysqlHealth && (
        <p style={{ color: mysqlHealth.includes('connected') ? 'green' : 'red' }}>
          MySQL: {mysqlHealth}
        </p>
      )}
    </main>
  );
}
