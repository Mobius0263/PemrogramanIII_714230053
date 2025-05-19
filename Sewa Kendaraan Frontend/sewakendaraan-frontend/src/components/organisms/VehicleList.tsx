import React, { useEffect, useState } from 'react';
import { getVehicles } from '../../services/VehicleService.ts';

type Vehicle = {
  id: string;
  brand: string;
  model: string;
  year: number;
  isAvailable: boolean;
};

const VehicleList: React.FC = () => {
  const [vehicles, setVehicles] = useState<Vehicle[]>([]);

  useEffect(() => {
    getVehicles().then(data => setVehicles(data));
  }, []);

  return (
    <div>
      <h2 className="text-xl font-bold mb-4">Vehicle List</h2>
      <ul>
        {vehicles.map(vehicle => (
          <li key={vehicle.id} className="border p-4 mb-2 rounded">
            <strong>{vehicle.brand} {vehicle.model}</strong> - {vehicle.year} - 
            {vehicle.isAvailable ? ' Available' : ' Not Available'}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default VehicleList;
