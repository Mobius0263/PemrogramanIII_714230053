import React, { useEffect, useState } from 'react';
import axios from 'axios';

interface Vehicle {
  id: string;
  brand: string;
  model: string;
  year: number;
  isAvailable: boolean;
}

const VehicleListTemplate: React.FC = () => {
  const [vehicles, setVehicles] = useState<Vehicle[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchVehicles = async () => {
      try {
        const response = await axios.get('http://localhost:3000/api/vehicles');
        setVehicles(response.data);
      } catch (error) {
        console.error('Error fetching vehicles:', error);
        setError('Failed to fetch vehicles');
      }
    };

    fetchVehicles();
  }, []);

  if (error) return <div>Error: {error}</div>;
  if (vehicles.length === 0) return <div>Loading...</div>;

  return (
    <div>
      <h2 className="text-2xl font-bold mb-4">Vehicle List</h2>
      <table className="min-w-full bg-white">
        <thead>
          <tr>
            <th className="py-2 px-4 border-b">ID</th>
            <th className="py-2 px-4 border-b">Brand</th>
            <th className="py-2 px-4 border-b">Model</th>
            <th className="py-2 px-4 border-b">Year</th>
            <th className="py-2 px-4 border-b">Available</th>
          </tr>
        </thead>
        <tbody>
          {vehicles.map((vehicle) => (
            <tr key={vehicle.id}>
              <td className="py-2 px-4 border-b">{vehicle.id}</td>
              <td className="py-2 px-4 border-b">{vehicle.brand}</td>
              <td className="py-2 px-4 border-b">{vehicle.model}</td>
              <td className="py-2 px-4 border-b">{vehicle.year}</td>
              <td className="py-2 px-4 border-b">{vehicle.isAvailable ? 'Yes' : 'No'}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default VehicleListTemplate;