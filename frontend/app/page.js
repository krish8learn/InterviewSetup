'use client';

import { useEffect, useState } from 'react';

export default function Home() {
  const [message, setMessage] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/`)
      .then((res) => {
        if (!res.ok) throw new Error(`HTTP ${res.status}`);
        return res.json();
      })
      .then((data) => setMessage(data.message ?? JSON.stringify(data)))
      .catch((err) => setError(err.message))
      .finally(() => setLoading(false));
  }, []);

  return (
    <main style={{ fontFamily: 'sans-serif', padding: '2rem' }}>
      <h1>Hello World</h1>
      {loading && <p>Loading from backend…</p>}
      {error && <p style={{ color: 'red' }}>Backend error: {error}</p>}
      {message && <p>Backend says: {message}</p>}
    </main>
  );
}
