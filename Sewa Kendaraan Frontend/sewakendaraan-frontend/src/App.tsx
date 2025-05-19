import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Link } from 'react-router-dom';
import Dashboard from './components/templates/Dashboard';
import RentListTemplate from './components/templates/RentListTemplate';
import VehicleListTemplate from './components/templates/VehicleListTemplate.tsx';
import ConsumerListTemplate from './components/templates/ConsumerListTemplate.tsx';

function App() {
  return (
    <BrowserRouter>
      <div>
        <nav className="bg-gray-800 p-4">
          <ul className="flex space-x-4">
            <li>
              <Link to="/" className="text-white hover:text-gray-300">Dashboard</Link>
            </li>
            <li>
              <Link to="/rents" className="text-white hover:text-gray-300">Rents</Link>
            </li>
            <li>
              <Link to="/vehicles" className="text-white hover:text-gray-300">Vehicles</Link>
            </li>
            <li>
              <Link to="/consumers" className="text-white hover:text-gray-300">Consumers</Link>
            </li>
          </ul>
        </nav>

        <main className="container mx-auto mt-4">
          <Routes>
            <Route path="/" element={<Dashboard />} />
            <Route path="/rents" element={<RentListTemplate />} />
            <Route path="/vehicles" element={<VehicleListTemplate />} />
            <Route path="/consumers" element={<ConsumerListTemplate />} />
          </Routes>
        </main>
      </div>
    </BrowserRouter>
  );
}

export default App;