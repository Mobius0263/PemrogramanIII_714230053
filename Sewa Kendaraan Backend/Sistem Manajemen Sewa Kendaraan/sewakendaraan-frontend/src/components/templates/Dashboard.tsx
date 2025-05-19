import React, { useEffect, useState } from 'react';
import axios from 'axios';

interface DashboardData {
  rents: number;
  vehicles: number;
  consumers: number;
}

const Dashboard: React.FC = () => {
  const [data, setData] = useState<DashboardData>({ rents: 0, vehicles: 0, consumers: 0 });

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [rentsRes, vehiclesRes, consumersRes] = await Promise.all([
          axios.get('http://localhost:3000/api/rents'),
          axios.get('http://localhost:3000/api/vehicles'),
          axios.get('http://localhost:3000/api/consumers')
        ]);

        setData({
          rents: rentsRes.data.length,
          vehicles: vehiclesRes.data.length,
          consumers: consumersRes.data.length
        });
      } catch (error) {
        console.error('Error fetching dashboard data:', error);
      }
    };

    fetchData();
  }, []);

  return (
    <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
      <div className="bg-blue-100 p-4 rounded-lg shadow">
        <h2 className="text-xl font-bold mb-2">Rents</h2>
        <p className="text-3xl">{data.rents}</p>
      </div>
      <div className="bg-green-100 p-4 rounded-lg shadow">
        <h2 className="text-xl font-bold mb-2">Vehicles</h2>
        <p className="text-3xl">{data.vehicles}</p>
      </div>
      <div className="bg-yellow-100 p-4 rounded-lg shadow">
        <h2 className="text-xl font-bold mb-2">Consumers</h2>
        <p className="text-3xl">{data.consumers}</p>
      </div>
    </div>
  );
};

export default Dashboard;