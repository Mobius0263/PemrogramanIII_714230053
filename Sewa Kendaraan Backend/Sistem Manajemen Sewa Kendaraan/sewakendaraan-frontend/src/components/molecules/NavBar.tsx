import React from 'react';
import { Link } from 'react-router-dom';

const NavBar: React.FC = () => {
  return (
    <nav className="bg-gray-800 p-4">
      <div className="container mx-auto flex justify-between items-center">
        <Link to="/" className="text-white font-bold text-xl">Rent Management</Link>
        <ul className="flex space-x-4">
          <li><Link to="/" className="text-white hover:text-gray-300">Home</Link></li>
          <li><Link to="/vehicles" className="text-white hover:text-gray-300">Vehicles</Link></li>
          <li><Link to="/consumers" className="text-white hover:text-gray-300">Consumers</Link></li>
          <li><Link to="/rents" className="text-white hover:text-gray-300">Rents</Link></li>
        </ul>
      </div>
    </nav>
  );
};

export default NavBar;