import React, { useEffect, useState } from 'react';
import axios from 'axios';

interface Consumer {
  id: string;
  name: string;
  phoneNumber: string;
}

const ConsumerListTemplate: React.FC = () => {
  const [consumers, setConsumers] = useState<Consumer[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchConsumers = async () => {
      try {
        const response = await axios.get('http://localhost:3000/api/consumers');
        setConsumers(response.data);
      } catch (error) {
        console.error('Error fetching consumers:', error);
        setError('Failed to fetch consumers');
      }
    };

    fetchConsumers();
  }, []);

  if (error) return <div>Error: {error}</div>;
  if (consumers.length === 0) return <div>Loading...</div>;

  return (
    <div>
      <h2 className="text-2xl font-bold mb-4">Consumer List</h2>
      <table className="min-w-full bg-white">
        <thead>
          <tr>
            <th className="py-2 px-4 border-b">ID</th>
            <th className="py-2 px-4 border-b">Name</th>
            <th className="py-2 px-4 border-b">Phone</th>
          </tr>
        </thead>
        <tbody>
          {consumers.map((consumer) => (
            <tr key={consumer.id}>
              <td className="py-2 px-4 border-b">{consumer.id}</td>
              <td className="py-2 px-4 border-b">{consumer.name}</td>
              <td className="py-2 px-4 border-b">{consumer.phoneNumber}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default ConsumerListTemplate;