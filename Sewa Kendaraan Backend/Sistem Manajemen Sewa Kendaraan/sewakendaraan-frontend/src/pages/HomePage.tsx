import React from 'react';
import VehicleList from '../components/organisms/VehicleList';

const HomePage: React.FC = () => {
  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-6">Vehicle Rental Management</h1>
      <VehicleList />
    </div>
  );
};

export default HomePage;
